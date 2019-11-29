package bitmex_test

import (
	"encoding/json"

	. "github.com/adampointer/go-bitmex"
	"github.com/adampointer/go-bitmex/fakes"
	"github.com/adampointer/go-bitmex/swagger/models"
	"github.com/adampointer/go-bitmex/types"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
	"github.com/wacul/ptr"
	"go.uber.org/multierr"
)

var _ = Describe("WebsocketClient", func() {
	var (
		wsClient *WebsocketClient
		fake     *fakes.Websocket
	)

	BeforeEach(func() {
		fake = fakes.NewWebsocket()
		cfg := NewWebsocketClientConfig().WithURL(fake.URL())
		wsClient = NewWebsocketClient(cfg)
	})

	AfterEach(func() {
		wsClient.Shutdown()
		fake.Shutdown()
	})

	Context("when Start() is called", func() {
		It("connects to a remote websocket server", func(done Done) {
			go func() {
				<-fake.Connected
				close(done)
			}()
			wsClient.Start()
		})

		It("sends an auth message if credentials are present", func(done Done) {
			cfg := NewWebsocketClientConfig().WithURL(fake.URL()).WithAuth("foo", "bar")
			wsClient.SetConfig(cfg)
			go func() {
				defer GinkgoRecover()
				msg := <-fake.Messages
				cmd := types.Command{}
				Expect(json.Unmarshal(msg, &cmd)).NotTo(HaveOccurred())
				Expect(cmd.Op).To(Equal(types.CommandOpAuth))
				Expect(cmd.Args[0]).To(Equal("foo"))
				close(done)
			}()
			wsClient.Start()
		})

		It("fires a websocket connected event on connect which you can subscribe to", func(done Done) {
			err := wsClient.SubscribeToEvents(types.EventWebsocketConnected, func(_ interface{}) {
				close(done)
			})
			Expect(err).NotTo(HaveOccurred())
			wsClient.Start()
		})

		It("send a subscribe message if the user has configured topic subscriptions", func(done Done) {
			err := multierr.Combine(
				wsClient.SubscribeToApiTopic(types.SubscriptionTopicOrder, func(response *types.SubscriptionResponse) {}),
				wsClient.SubscribeToApiTopic(types.SubscriptionTopicLiquidation, func(response *types.SubscriptionResponse) {}),
			)
			Expect(err).NotTo(HaveOccurred())
			go func() {
				defer GinkgoRecover()
				msg := <-fake.Messages
				cmd := types.Command{}
				Expect(json.Unmarshal(msg, &cmd)).NotTo(HaveOccurred())
				Expect(cmd.Op).To(Equal(types.CommandOpSubscribe))
				Expect(cmd.Args).To(Equal(types.CommandArgs{"order", "liquidation"}))
				close(done)
			}()
			wsClient.Start()
		})
	})

	Context("when a message is received for a subscribed topic", func() {
		It("triggers an appropriate event", func(done Done) {
			err := wsClient.SubscribeToApiTopic(types.SubscriptionTopicChat, func(response *types.SubscriptionResponse) {
				defer GinkgoRecover()
				Expect(response.Table).To(Equal("chat"))
				last := func(_ interface{}) string { return "0" }
				Expect(response.ChatData()).Should(MatchElements(last, IgnoreExtras, Elements{
					"0": PointTo(MatchFields(IgnoreExtras, Fields{
						"ChannelID": Equal(200.0),
						"User":      PointTo(Equal("spammer")),
					})),
				}))
				close(done)
			})

			Expect(err).NotTo(HaveOccurred())
			wsClient.Start()

			t := strfmt.NewDateTime()
			fake.SendWebsocketMessage(GinkgoT(), fakeSuccessResponse("chat", true))
			fake.SendWebsocketMessage(GinkgoT(), fakeSubscriptionResponse(GinkgoT(), "chat", &models.Chat{
				ChannelID: 200,
				Date:      &t,
				FromBot:   ptr.Bool(true),
				ID:        100,
				Message:   ptr.String("hello"),
				User:      ptr.String("spammer"),
			}))
		})
	})
})

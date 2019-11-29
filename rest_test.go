package bitmex_test

import (
	"os"

	. "github.com/adampointer/go-bitmex"
	"github.com/adampointer/go-bitmex/swagger/client"
	"github.com/adampointer/go-bitmex/swagger/client/order"
	"github.com/adampointer/go-bitmex/swagger/client/trade"
	"github.com/adampointer/go-bitmex/swagger/client/user"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
	"github.com/wacul/ptr"
)

const url = "https://testnet.bitmex.com/"

var _ = Describe("RestClient", func() {
	var (
		restClient  *client.BitMEX
		key, secret string
	)

	JustBeforeEach(func() {
		cfg := NewClientConfig().WithURL(url).WithAuth(key, secret)
		restClient = NewBitmexClient(cfg)
	})

	Context("when a secured endpoint is called", func() {
		BeforeEach(func() {
			key = os.Getenv("BITMEX_API_KEY")
			secret = os.Getenv("BITMEX_API_SECRET")
			if len(key) == 0 || len(secret) == 0 {
				GinkgoT().Fatal("BITMEX_API_KEY and BITMEX_API_SECRET must be set to run acceptance tests")
			}
		})

		It("should respond with a valid wallet object when user/wallet is queried", func() {
			res, err := restClient.User.UserGetWallet(user.NewUserGetWalletParams())
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Payload).Should(PointTo(MatchFields(IgnoreExtras, Fields{
				"Currency": PointTo(Equal("XBt")),
			})))
		})

		It("should allow a limit orders to be posted and cancelled", func() {

			By("Posting an order")

			createParams := order.NewOrderNewParams()
			createParams.SetSymbol("XBTUSD")
			createParams.SetOrdType(ptr.String("Limit"))
			createParams.SetOrderQty(ptr.Int32(10))
			createParams.SetPrice(ptr.Float64(100))

			res1, err := restClient.Order.OrderNew(createParams)

			Expect(err).NotTo(HaveOccurred())
			Expect(res1.Payload).Should(PointTo(MatchFields(IgnoreExtras, Fields{
				"OrdStatus": Equal("New"),
				"OrdType":   Equal("Limit"),
				"OrderID":   Not(BeNil()),
			})))

			By("Then cancelling it")

			cancelParams := order.NewOrderCancelParams()
			cancelParams.SetOrderID(res1.Payload.OrderID)

			res2, err := restClient.Order.OrderCancel(cancelParams)

			Expect(err).NotTo(HaveOccurred())
			Expect(res2.Payload).Should(MatchElements(last, IgnoreExtras, Elements{
				"0": PointTo(MatchFields(IgnoreExtras, Fields{
					"OrdStatus": Equal("Canceled"),
				})),
			}))
		})
	})

	Context("when a public endpoint is called", func() {
		It("should return trade data when trade/bucketed is queried", func() {
			params := trade.NewTradeGetBucketedParams()
			params.SetSymbol(ptr.String("XBTUSD"))
			params.SetBinSize(ptr.String("1d"))
			params.SetReverse(ptr.Bool(true))
			params.SetStartTime(DateTimePtr("30 Sep 19 00:00 UTC"))
			params.SetEndTime(DateTimePtr("30 Sep 19 23:59 UTC"))

			res, err := restClient.Trade.TradeGetBucketed(params)

			Expect(err).NotTo(HaveOccurred())
			Expect(res.Payload).Should(MatchElements(last, IgnoreExtras, Elements{
				"0": PointTo(MatchFields(IgnoreExtras, Fields{
					"Open":  Equal(8439.5),
					"High":  Equal(8555.0),
					"Low":   Equal(8062.0),
					"Close": Equal(8130.0),
				})),
			}))
		})
	})
})

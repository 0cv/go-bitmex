package bitmex_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"
	"time"

	"github.com/adampointer/go-bitmex/types"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
	"github.com/wacul/ptr"
)

func TestBitmex(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Bitmex Suite")
}

var _ = BeforeSuite(func() {
	logrus.SetOutput(ioutil.Discard)
	logrus.SetLevel(logrus.DebugLevel)
})

func fakeSubscriptionResponse(t GinkgoTInterface, table string, data interface{}) *types.CompositeResponse {
	d := []interface{}{data}
	j, err := json.Marshal(d)
	if err != nil {
		t.Fatal(err)
	}
	return &types.CompositeResponse{
		SubscriptionResponse: types.SubscriptionResponse{
			Table:  table,
			Action: "update",
			Data:   j,
		},
		Request: &types.Command{Op: types.CommandOpSubscribe, Args: types.CommandArgs{table}},
	}
}

func fakeSuccessResponse(table string, success bool) *types.CompositeResponse {
	return &types.CompositeResponse{
		SuccessResponse: types.SuccessResponse{
			Subscribe: table,
			Success:   ptr.Bool(success),
			Request:   &types.Command{Op: types.CommandOpSubscribe, Args: types.CommandArgs{table}},
		},
		Request: &types.Command{Op: types.CommandOpSubscribe, Args: types.CommandArgs{table}},
	}
}

func last(_ interface{}) string {
	return "0"
}

func DateTimePtr(timeString string) *strfmt.DateTime {
	t, err := time.Parse(time.RFC822, timeString)
	if err != nil {
		panic(err)
	}
	d := strfmt.DateTime(t)
	return &d
}

package types

type CommandOp string

const (
	CommandOpSubscribe      CommandOp = "subscribe"
	CommandOpUnsubscribe    CommandOp = "unsubscribe"
	CommandOpAuth           CommandOp = "authKeyExpires"
	CommandOpPing           CommandOp = "ping"
	CommandOpCancelAllAfter CommandOp = "cancelAllAfter"
)

type CommandArgs []interface{}

type Command struct {
	Op   CommandOp   `json:"op"`
	Args CommandArgs `json:"args,omitempty"`
}

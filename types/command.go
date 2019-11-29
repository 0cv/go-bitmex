package types

// CommandOp is a websocket API command operation
type CommandOp string

const (
	// CommandOpSubscribe subscribes to a topic
	CommandOpSubscribe CommandOp = "subscribe"
	// CommandOpUnsubscribe unsubscribes from a topic
	CommandOpUnsubscribe CommandOp = "unsubscribe"
	// CommandOpAuth authenticates the connection
	CommandOpAuth CommandOp = "authKeyExpires"
	// CommandOpPing is a ping
	CommandOpPing CommandOp = "ping"
	// CommandOpCancelAllAfter sets up a dead-man's switch
	CommandOpCancelAllAfter CommandOp = "cancelAllAfter"
)

// CommandArgs are arguments for a command
type CommandArgs []interface{}

// Command can be marshaled to JSON and sent on the websocket
type Command struct {
	Op   CommandOp   `json:"op"`
	Args CommandArgs `json:"args,omitempty"`
}

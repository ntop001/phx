package phx

// Event represents a phoenix channel event for a message, and can be almost anything the user wants, such as "ping",
// "shout", "talk", etc. There are several reserved for control messages for protocol overhead.
type Event string

// Special/reserved Events
const (
	// JoinEvent is sent when we join a channel.
	JoinEvent Event = "phx_join"

	// CloseEvent is sent by the server when a channel is closed, such as before shutting down the socket
	// This event is also generated by the client whenever `channel.Leave()` is called. Triggers channel.OnClose().
	CloseEvent Event = "phx_close"

	// ErrorEvent is sent by the server whenever something catastrophic happens on the server side, such as the channel
	// process crashing, or attempting to join a channel already joined.
	ErrorEvent Event = "phx_error"

	// ReplyEvent is sent by the server in reply to any event sent by the client.
	ReplyEvent Event = "phx_reply"

	// LeaveEvent is sent by the client when we leave a channel and unsubscribe to a topic.
	LeaveEvent Event = "phx_leave"

	// HeartBeatEvent is a special message for heartbeats on the special topic "phoenix"
	HeartBeatEvent Event = "heartbeat"
)
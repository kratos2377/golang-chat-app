package websocket

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	BroadCast  chan *Message
}

func NewPool() {

}

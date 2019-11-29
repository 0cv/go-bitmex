package fakes

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/gorilla/websocket"
	"github.com/onsi/ginkgo"
	log "github.com/sirupsen/logrus"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Websocket wraps httptest.Server to provide a fake websocket server for testing
type Websocket struct {
	Connected chan bool
	Messages  chan []byte
	Server    *httptest.Server
	conn      *websocket.Conn
}

// NewWebsocket returns a new initialised fake websocket
func NewWebsocket() *Websocket {
	ws := &Websocket{
		Connected: make(chan bool, 10),
		Messages:  make(chan []byte, 10),
	}
	ws.Server = httptest.NewServer(http.HandlerFunc(ws.handler))
	return ws
}

// URL returns the url of the fake
func (w *Websocket) URL() string {
	u, _ := url.Parse(w.Server.URL)
	u.Scheme = "ws"
	return u.String()
}

// SendWebsocketMessage sends a message to the client
func (w *Websocket) SendWebsocketMessage(t ginkgo.GinkgoTInterface, msg interface{}) {
	ginkgo.GinkgoRecover()
	if err := w.conn.WriteJSON(msg); err != nil {
		t.Fatal(err)
	}
}

// Shutdown cleans up channels
func (w *Websocket) Shutdown() {
	close(w.Connected)
	close(w.Messages)
}

func (w *Websocket) handler(res http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(res, req, nil)
	if err != nil {
		http.Error(res, fmt.Sprintf("cannot upgrade: %v", err), http.StatusInternalServerError)
	}
	w.conn = conn
	w.Connected <- true
	_, msg, err := conn.ReadMessage()
	if err != nil {
		log.Errorf("cannot read message: %v", err)
		return
	}
	w.Messages <- msg
}

package hub

import (
	"time"

	"github.com/fiffu/vapormail/dto"
	"github.com/fiffu/vapormail/utils"
	"github.com/gorilla/websocket"
)

type Client struct {
	hub       *HubService
	id        dto.ClientID
	inboxes   []dto.Inbox
	sessionID string
	conn      *websocket.Conn
	sendQueue chan dto.IMessage
	key       string
}

type ClientEvent struct {
	client    *Client
	kind      string
	timestamp time.Time
}

func makeID() dto.ClientID {
	id := utils.RandomUUID()
	return dto.ClientID(id)
}

func (h *HubService) NewClient(sessionID, key string, conn *websocket.Conn, inboxes []string) *Client {
	var ibx []dto.Inbox
	for _, i := range inboxes {
		ibx = append(ibx, dto.Inbox(i))
	}
	return &Client{
		id:        makeID(),
		sessionID: sessionID,
		key:       key,
		inboxes:   ibx,
		hub:       h,
		conn:      conn,
		sendQueue: make(chan dto.IMessage),
	}
}

func (c *Client) Send(msg dto.IMessage) {
	c.sendQueue <- msg
}

func (c *Client) SendSystemMessage(opcode Opcode, isError bool) {
	errorLevel := dto.LevelDefault
	if isError {
		errorLevel = dto.LevelError
	}
	if len(c.inboxes) == 0 {
		// No inboxes can receive this message
		return
	}
	// Pick any inbox to send to.
	inbox := string(c.inboxes[0])
	msg := &SystemMessage{
		[]string{inbox},
		opcode,
		errorLevel,
		time.Now(),
	}
	c.Send(msg)
}

func (c *Client) onJoin() {
	c.hub.EventBus <- ClientEvent{
		kind:      EventClientJoined,
		timestamp: time.Now(),
		client:    c,
	}
	c.SendSystemMessage(OpcodeWelcome, false)
}

func (c *Client) onLeave() {
	c.hub.EventBus <- ClientEvent{
		kind:      EventClientLeft,
		timestamp: time.Now(),
		client:    c,
	}
}

func (c *Client) Evict() {
	c.hub.EventBus <- ClientEvent{
		kind:      EventClientEvicted,
		timestamp: time.Now(),
		client:    c,
	}
	close(c.sendQueue)
}

// go join
func (c *Client) Join() {
	c.onJoin()

	interval := c.hub.config.ClientHeartbeat
	ticker := time.NewTicker(time.Duration(interval))
	defer func() {
		ticker.Stop()
		c.conn.Close()
		c.onLeave()
	}()
	for {
		select {
		case message, ok := <-c.sendQueue:
			c.conn.SetWriteDeadline(time.Now().Add(time.Duration(interval)))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			// Get writer.
			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			// Sink the queue.
			w.Write(message.Contents())
			n := len(c.sendQueue)
			for i := 0; i < n; i++ {
				w.Write((<-c.sendQueue).Contents())
			}

			// Exit if socket closed.
			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(time.Duration(interval)))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

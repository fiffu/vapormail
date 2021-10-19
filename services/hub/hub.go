package hub

import (
	"time"

	"github.com/fiffu/vprmail/config"
	"github.com/fiffu/vprmail/dto"
)

type HubService struct {
	config     config.Config
	clients    map[dto.Inbox]Client
	nonces     map[string]time.Time
	MessageBus chan dto.IMessage
	EventBus   chan ClientEvent
}

func NewHubService(cfg config.Config) *HubService {
	return &HubService{
		config:     cfg,
		clients:    make(map[dto.Inbox]Client),
		MessageBus: make(chan dto.IMessage),
		EventBus:   make(chan ClientEvent),
	}
}

func (h *HubService) Start() error {
	for {
		var err error

		select {
		case m := <-h.MessageBus:
			err = h.handleMessage(m)

		case evt := <-h.EventBus:
			err = h.handleEvent(evt)
		}

		if err != nil {
			return err
		}
	}
}

func (h *HubService) handleMessage(msg dto.IMessage) error {
	for _, r := range msg.Recipients() {
		err := msg.Error()
		if err != nil {
			return err
		}
		inbox := dto.Inbox(r)
		if client, ok := h.clients[inbox]; ok {
			client.Send(msg)
			client.SendSystemMessage(OpcodeEOF, false)
		}
	}
	return nil
}

func (h *HubService) handleEvent(evt ClientEvent) error {
	switch evt.kind {

	case EventClientJoined:
		for _, inbox := range evt.client.inboxes {
			h.clients[inbox] = *evt.client
		}

	case EventClientLeft, EventClientEvicted:
		for _, inbox := range evt.client.inboxes {
			delete(h.clients, inbox)
		}
	}

	return nil
}

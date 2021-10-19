package hub

import (
	"time"
)

type Opcode string

const (
	// Opcodes are events delivered to clients. They don't have a payload.
	OpcodeWelcome = Opcode("WELCOME")
	OpcodeEOF     = Opcode("EOF")

	// Sentinel value when delivering notifs to client
	SystemSender = "SYSTEM"

	// Internal system events
	EventClientJoined  = "CLIENT_JOINED"
	EventClientLeft    = "CLIENT_LEFT"
	EventClientEvicted = "CLIENT_EVICTED"
)

type SystemMessage struct {
	recipients []string
	opcode     Opcode
	errorLevel int
	timestamp  time.Time
}

func (m *SystemMessage) Error() error {
	return nil
}

func (m *SystemMessage) Timestamp() time.Time {
	return m.timestamp
}

func (m *SystemMessage) Origin() string {
	return SystemSender
}

func (m *SystemMessage) Sender() string {
	return SystemSender
}

func (m *SystemMessage) Recipients() []string {
	return m.recipients
}

func (m *SystemMessage) Contents() []byte {
	return []byte(m.opcode)
}

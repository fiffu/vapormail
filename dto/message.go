package dto

import (
	"time"
)

const (
	LevelInvisible = -10
	LevelDebug     = -1
	LevelDefault   = 0
	LevelInfo      = 1
	LevelWarning   = 2
	LevelError     = 3
)

type IMessage interface {
	Error() error
	Contents() []byte
	Timestamp() time.Time
	Origin() string
	Sender() string
	Recipients() []string
}

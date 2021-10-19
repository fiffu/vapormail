package utils

import (
	"strconv"
	"strings"
	"time"

	petName "github.com/dustinkirkland/golang-petname"
	"github.com/google/uuid"
)

func RandomUUID() string {
	return uuid.New().String()
}

func RandomName(words int, separator string) string {
	return petName.Generate(words, separator)
}

func TimestampString() string {
	t := time.Now().UnixNano()
	s := strconv.FormatInt(t, 36)
	return strings.ToUpper(s)
}

package hub

import (
	"github.com/fiffu/vprmail/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (h *HubService) NewSocket(
	ctx *gin.Context,
	nonce string,
	username, password string,
	inboxes []string,
) (string, string, []string, utils.ICustomError) {

	sessionID, allowedInboxes, err := h.NewSession(
		nonce, username, password, inboxes,
	)
	key := h.Hash(username, password)
	if err != nil {
		return "", "", nil, utils.Error(err)
	}

	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		return "", "", nil, utils.Error(err)
	}

	client := h.NewClient(sessionID, key, conn, allowedInboxes)
	go client.Join()
	return sessionID, key, allowedInboxes, nil
}

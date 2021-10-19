package services

import (
	"github.com/fiffu/vprmail/dto"
	"github.com/fiffu/vprmail/utils"
	"github.com/gin-gonic/gin"
)

func (s *HTTPService) setupAPI(router *gin.Engine, groupRoute string) {
	v1 := router.Group(groupRoute + "/v1")
	v1.POST("/new", s.handleNew)
	v1.POST("/connect", s.handleConnect)
}

func (s *HTTPService) handleNew(ctx *gin.Context) {
	res := dto.StartSessionResponse{Nonce: s.Hub.NewNonce()}
	s.resolveOK(ctx, res)
}

func (s *HTTPService) handleConnect(ctx *gin.Context) {
	var req dto.ConnectRequest
	if err := ctx.BindJSON(&req); err != nil {
		s.reject(ctx, utils.ClientError(err))
	}
	sessionID, key, allowedInboxes, err := s.Hub.NewSocket(
		ctx,
		req.Nonce,
		req.Username,
		req.Password,
		req.Inboxes,
	)
	if err != nil {
		s.reject(ctx, utils.ClientError(err))
	}

	res := dto.ConnectResponse{
		Success:        true,
		Key:            key,
		SessionID:      sessionID,
		AllowedInboxes: allowedInboxes,
	}
	s.resolveOK(ctx, res)
}

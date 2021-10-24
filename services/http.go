package services

import (
	"fmt"
	"net/http"

	"github.com/fiffu/vapormail/services/hub"
	"github.com/fiffu/vapormail/utils"
	"github.com/gin-gonic/gin"
	requestid "github.com/sumit-tembe/gin-requestid"
)

type HTTPService struct {
	Port        int
	Hub         *hub.HubService
	InboxDomain string
}

func NewHTTPService(port int, domain string, hubSvc *hub.HubService) HTTPService {
	return HTTPService{
		Port:        port,
		InboxDomain: domain,
		Hub:         hubSvc,
	}
}

func (s *HTTPService) Start() {
	router := gin.Default()

	// Middleware to inject requestID and add it to logging
	router.Use(requestid.RequestID(nil))
	router.Use(gin.LoggerWithConfig(requestid.GetLoggerConfig(nil, nil, nil)))

	s.setupMeta(router, "/")
	s.setupAPI(router, "/api")
	s.setupGUI(router, "/")

	router.Run(fmt.Sprintf(":%d", s.Port))
}

func (s *HTTPService) setupMeta(router *gin.Engine, groupRoute string) {
	meta := router.Group(groupRoute)
	meta.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

func (s *HTTPService) resolve(ctx *gin.Context, statusCode int, result interface{}) {
	ctx.JSON(statusCode, result)
}

func (s *HTTPService) resolveOK(ctx *gin.Context, result interface{}) {
	ctx.JSON(http.StatusOK, result)
}

func (s *HTTPService) reject(ctx *gin.Context, err utils.ICustomError) {
	s.resolve(
		ctx,
		err.Code(),
		gin.H{"error": err.Error()},
	)
}

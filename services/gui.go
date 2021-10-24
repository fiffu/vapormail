package services

import (
	"net/http"
	"path/filepath"

	"github.com/fiffu/vapormail/ui"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (s *HTTPService) setupGUI(router *gin.Engine, groupRoute string) {
	templatesPath := ui.GetTemplatesGlob()
	log.Info("Loading templates at path=%s", templatesPath)
	router.LoadHTMLGlob(filepath.Join(templatesPath, "*"))

	controller := router.Group(groupRoute)
	controller.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, ui.Index(), gin.H{
			"title": "Main website",
		})
	})
}

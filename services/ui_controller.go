package services

import (
	"net/http"
	"path/filepath"

	"github.com/fiffu/vapormail/ui"
	"github.com/fiffu/vapormail/utils"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

const (
	urlStubStatics = "/src"
)

func (s *HTTPService) setRoutesHTML(controller *gin.RouterGroup) {
	controller.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, ui.Index(), gin.H{
			"inboxDomain":      s.InboxDomain,
			"inboxPlaceholder": utils.RandomName(3, "."),
			"static":           urlStubStatics,
		})
	})
}

func (s *HTTPService) setRoutesJavascript(controller *gin.RouterGroup) {
	staticsDir := ui.GetStaticsDir()
	controller.Static(urlStubStatics, staticsDir)
}

func (s *HTTPService) setupGUI(engine *gin.Engine, uiRoute string) {
	templatesDir := ui.GetTemplatesDir()
	engine.LoadHTMLGlob(filepath.Join(templatesDir, "*"))
	log.Infof("Loaded templates at path=%s", templatesDir)

	controller := engine.Group(uiRoute)
	s.setRoutesHTML(controller)
	s.setRoutesJavascript(controller)
}

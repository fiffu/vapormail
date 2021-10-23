package ui

import (
	"path/filepath"

	"github.com/fiffu/vprmail/utils"
)

const (
	staticFolder  = "templates"
	fileHtmlIndex = "index.html"
)

func Index() string {
	return fileHtmlIndex
}

func GetTemplatesGlob() string {
	here, _ := utils.GetRuntimeDir()
	return filepath.Join(here, "ui", staticFolder)
}

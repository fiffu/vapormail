package ui

import (
	"path/filepath"

	"github.com/fiffu/vapormail/utils"
)

const (
	here            = "ui"
	templatesFolder = "templates"
	staticsFolder   = "static"
	fileHtmlIndex   = "index.html"
)

func Index() string {
	return fileHtmlIndex
}

func fromHere(args ...string) string {
	root, _ := utils.GetRuntimeDir()
	abspath := []string{root, here}
	abspath = append(abspath, args...)
	return filepath.Join(abspath...)
}

func GetStaticsDir() string {
	return fromHere(staticsFolder)
}

func GetTemplatesDir() string {
	return fromHere(templatesFolder)
}

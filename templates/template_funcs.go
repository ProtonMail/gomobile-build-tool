package templates

import (
	"text/template"

	"github.com/ProtonMail/gomobile-build-tool/utils"
)

var commonFuncs = template.FuncMap{
	"formatGoVersion":    formatGoVersion,
	"formatReplacements": formatReplacements,
	"formatRequirements": formatRequirements,
	"formatModules":      formatModules,
	"formatFlags":        formatFlags,
	"formatPackages":     formatPackages,
	"buildApple":         buildApple,
	"buildAndroid":       buildAndroid,
	"go14orAbove":        utils.Go14orBelow,
	"go16orAbove":        utils.Go16orAbove,
}

func newTemplate(name string) *template.Template {
	return template.New(name).Funcs(commonFuncs)
}

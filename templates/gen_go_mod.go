package templates

import (
	"fmt"
	"os"
	"path"
	"strings"
	"text/template"

	config "github.com/ProtonMail/gomobile-build-tool/configloader"
)

func formatGoVersion(version string) string {
	versionSplitted := strings.Split(version, ".")

	return strings.Join(versionSplitted[:2], ".")
}

func formatModule(module config.Module) string {
	return fmt.Sprintf("%s %s", module.Path, module.Version)
}

func formatRequirements(requirements []config.Requirement) string {
	requirementStrings := make([]string, len(requirements))
	for i, requirement := range requirements {
		requirementStrings[i] = formatRequirement(requirement)
	}

	return strings.Join(requirementStrings, "\n")
}

func formatRequirement(requirement config.Requirement) string {
	return "\t" + formatModule(requirement.Module)
}

func formatReplacements(replacements []config.Replacement) string {
	replacementStrings := make([]string, len(replacements))
	for i, replacement := range replacements {
		replacementStrings[i] = formatReplacement(replacement)
	}

	return strings.Join(replacementStrings, "\n")
}

func formatReplacement(replacement config.Replacement) string {
	if replacement.New == (config.Module{}) {
		return fmt.Sprintf(
			"replace %s => %s",
			formatModule(replacement.Old),
			replacement.LocalPath,
		)
	}

	return fmt.Sprintf(
		"replace %s => %s",
		formatModule(replacement.Old),
		formatModule(replacement.New),
	)
}

type GoModTemplateData struct {
	Config *config.Config
}

func GenerateGoMod(config *config.Config) (err error) {
	var template *template.Template = newTemplate("go.mod.template")
	template, err = template.ParseFiles("templates/files/go.mod.template")
	if err != nil {
		return err
	}
	filename := path.Join(config.BuildDir, "go.mod")
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	err = template.Execute(f, &GoModTemplateData{Config: config})
	if err != nil {
		return err
	}

	return f.Close()
}

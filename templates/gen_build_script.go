package templates

import (
	"fmt"
	"os"
	"path"
	"strings"
	"text/template"

	config "github.com/ProtonMail/gomobile-build-tool/configloader"
)

func formatFlags(flags []string) string {
	return strings.Join(flags, " ")
}

func formatModules(requirements []config.Requirement) string {
	moduleStrings := make([]string, len(requirements))
	for i, requirement := range requirements {
		moduleStrings[i] = fmt.Sprintf(
			"%s@%s",
			requirement.Module.Path,
			requirement.Module.Version,
		)
	}

	return strings.Join(moduleStrings, " ")
}

func formatPackages(requirements []config.Requirement) string {
	requirementStrings := make([]string, len(requirements))
	for i, requirement := range requirements {
		requirementStrings[i] = formatPackage(requirement)
	}

	return strings.Join(requirementStrings, " ")
}

func formatPackage(requirement config.Requirement) string {
	if len(requirement.Packages) == 0 {
		return requirement.Module.Path
	}
	packages := make([]string, len(requirement.Packages))
	for i, packageName := range requirement.Packages {
		packages[i] = fmt.Sprintf("%s/%s",
			requirement.Module.Path,
			packageName,
		)
	}

	return strings.Join(packages, " ")
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}

	return false
}

func buildApple(targets []string) bool {
	return len(targets) == 0 || contains(targets, "apple")
}

func buildAndroid(targets []string) bool {
	return len(targets) == 0 || contains(targets, "android")
}

type BuildScriptTemplateData struct {
	Config      *config.Config
	GomobileDir string
}

func GenerateBuildScript(config *config.Config) (err error) {
	var template *template.Template = newTemplate("build.sh.template")
	template, err = template.ParseFiles("templates/files/build.sh.template")
	if err != nil {
		return err
	}
	filename := path.Join(config.BuildDir, "build.sh")
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	err = template.Execute(f, &BuildScriptTemplateData{Config: config})
	if err != nil {
		return err
	}
	err = f.Close()
	if err != nil {
		return err
	}

	return nil
}

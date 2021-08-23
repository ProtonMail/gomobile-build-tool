package main

import (
	"os"
	"os/exec"
	"path"
	"path/filepath"

	config "github.com/ProtonMail/gomobile-build-tool/configloader"

	"github.com/ProtonMail/gomobile-build-tool/templates"

	"github.com/ProtonMail/gomobile-build-tool/utils"
)

func loadConfig() (configuration *config.Config, err error) {
	configFileName := "config.json"
	if len(os.Args) > 1 {
		configFileName = os.Args[1]
	}

	return config.LoadConfig(configFileName)
}

func generateTemplates(configuration *config.Config) (err error) {
	err = templates.GenerateGoMod(configuration)
	if err != nil {
		return err
	}

	return templates.GenerateBuildScript(configuration)
}

func createDir(path string) (err error) {
	exists, err := exists(path)
	if err != nil {
		return err
	}
	if !exists {
		return os.Mkdir(path, 0744)
	}

	return nil
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

func setDirectories(configuration *config.Config) (err error) {
	configuration.BuildDir, err = filepath.Abs(configuration.BuildDir)
	if err != nil {
		return err
	}
	configuration.OutDir, err = filepath.Abs(configuration.OutDir)
	if err != nil {
		return err
	}
	configuration.GoMobileDir, err = filepath.Abs(configuration.GoMobileDir)
	if err != nil {
		return err
	}
	err = createDir(configuration.BuildDir)
	if err != nil {
		return err
	}

	return createDir(configuration.OutDir)
}

func runScript(configuration *config.Config) (err error) {
	buildScript := path.Join(configuration.BuildDir, "build.sh")
	cmd := exec.Command("sh", buildScript) // #nosec
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func exit(err error) {
	println(err.Error())
	os.Exit(1)
}

func main() {
	configuration, err := loadConfig()
	if err != nil {
		exit(err)
	}
	version, err := utils.GetGoVersion()
	if err != nil {
		exit(err)
	}
	if version != configuration.GoVersion {
		println("Err: The output of `go version` does not match the one in your configuration")
		return //nolint:nlreturn
	}
	err = setDirectories(configuration)
	if err != nil {
		exit(err)
	}
	err = generateTemplates(configuration)
	if err != nil {
		exit(err)
	}
	err = runScript(configuration)
	if err != nil {
		exit(err)
	}
}

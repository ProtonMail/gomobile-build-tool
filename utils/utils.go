package utils

import (
	"os/exec"
	"strings"

	"github.com/hashicorp/go-version"
)

func GetGoVersion() (string, error) {
	cmd := exec.Command("go", "version")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	versionOutput := strings.Split(string(output), " ")
	versionNumber := versionOutput[2][2:]

	return versionNumber, nil
}

func Go14orBelow(versionNumber string) bool {
	v, err := version.NewVersion(versionNumber)
	if err != nil {
		panic(err)
	}
	constraints, err := version.NewConstraint("< 1.15.0")
	if err != nil {
		panic(err)
	}

	return constraints.Check(v)
}

func Go16orAbove(versionNumber string) bool {
	v, err := version.NewVersion(versionNumber)
	if err != nil {
		panic(err)
	}
	constraints, err := version.NewConstraint(">= 1.16.0")
	if err != nil {
		panic(err)
	}

	return constraints.Check(v)
}

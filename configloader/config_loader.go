package configloader

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

type Config struct {
	BuildDir      string        `json:"build_dir"`
	OutDir        string        `json:"out_dir"`
	GoMobileDir   string        `json:"go_mobile_dir"`
	Requirements  []Requirement `json:"requirements"`
	Replacements  []Replacement `json:"replacements"`
	GoMobileFlags []string      `json:"go_mobile_flags"`
	BuildName     string        `json:"build_name"`
	Targets       []string      `json:"targets"`
	JavaPkg       string        `json:"java_pkg"`
	GoVersion     string        `json:"go_version"`
	BuildTag      string        `json:"build_tag"`
}

type Requirement struct {
	Module   Module   `json:"module"`
	Packages []string `json:"packages"`
}

type Module struct {
	Path    string `json:"path"`
	Version string `json:"version"`
}

type Replacement struct {
	Old       Module `json:"old"`
	New       Module `json:"new"`
	LocalPath string `json:"local_path"`
}

func LoadConfig(filename string) (config *Config, err error) {
	byteValue, err := ioutil.ReadFile(filepath.Clean(filename))
	if err != nil {
		return nil, err
	}
	config = &Config{
		BuildDir:    "build",
		OutDir:      "out",
		GoMobileDir: "mobile",
	}
	err = json.Unmarshal(byteValue, config)
	if err != nil {
		return nil, err
	}

	return
}

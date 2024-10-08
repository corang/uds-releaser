package utils

import (
	"os"

	"github.com/corang/uds-releaser/src/types"
	goyaml "github.com/goccy/go-yaml"
)

func LoadReleaserConfig() (types.ReleaserConfig, error) {

	var config types.ReleaserConfig
	err := LoadYaml("releaser.yaml", &config)
	if err != nil {
		return types.ReleaserConfig{}, err
	}

	return config, nil
}

func LoadYaml(path string, destVar interface{}) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return goyaml.Unmarshal(data, destVar)
}

func UpdateYaml(path string, srcVar interface{}) error {
	data, err := goyaml.Marshal(srcVar)
	if err != nil {
		return err
	}
	
	yamlInfo, err := os.Stat(path)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, yamlInfo.Mode())
}

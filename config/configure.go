package config

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

func LoadYamlConf(filepath string, s interface{}) error {
	if yamlFile, err := os.ReadFile(filepath); err != nil {
		return fmt.Errorf(filepath + " get error: " + err.Error())
	} else if err = yaml.Unmarshal(yamlFile, s); err != nil {
		return fmt.Errorf(filepath + " unmarshal error: " + err.Error())
	}
	return nil
}

func LoadJsonConf(filepath string, s interface{}) error {
	if jsonFile, err := os.ReadFile(filepath); err != nil {
		return fmt.Errorf("filepaht" + "get error: " + err.Error())
	} else if err = json.Unmarshal(jsonFile, s); err != nil {
		return fmt.Errorf(filepath + " unmarshal error: " + err.Error())
	}
	return nil
}

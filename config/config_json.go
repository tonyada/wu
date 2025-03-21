package config

import (
	gojson "encoding/json"
	"wu/os/file"
)

func InitJSON(filepath string, isOverwrite bool, config any) error {
	return NewJSON(filepath, true, config)
}

// create a config.json file
// config.NewJSON("config.json", false, settings)
func NewJSON(filepath string, isOverwrite bool, config any) error {
	jsonBody, err := gojson.MarshalIndent(config, "", "\t")
	if err != nil {
		return err
	}
	if isOverwrite {
		_, err = file.Save(filepath, string(jsonBody))
		return err
	}
	if file.IsExist(filepath) {
		return nil
	}
	return nil
}

// load json from config.json file
// config.LoadJSON("config.json", &settings)
func LoadJSON(filepath string, config any) error {
	configJson, err := file.Read(filepath)
	if err != nil {
		return err
	}
	return gojson.Unmarshal(configJson, &config)
}

package main

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"
)

type Config struct {
	Openai struct {
		ApiKey       string `yaml:"token"`
		MainPrompt   string `yaml:"main_prompt"`
		CustomSyntax string `yaml:"custom_syntax"`
	} `yaml:"openai"`

	Gitlab struct {
		ApiKey     string `yaml:"token"`
		ProjectIds []int  `yaml:"project_ids"`
	} `yaml:"gitlab"`

	Telegram struct {
		Token     string `yaml:"token"`
		ChatId    int64  `yaml:"chat_id"`
		ParseMode string `yaml:"parse_mode"`
	} `yaml:"telegram"`

	Days int `yaml:"days"`
}

func readConfig(path string) Config {
	log.Println("Reading config", path)

	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Error:", err.Error())
	}

	config := Config{}
	if err := yaml.Unmarshal(content, &config); err != nil {
		log.Fatal("Error:", err.Error())
	}

	configValue := reflect.ValueOf(&config).Elem()
	var errorFields []string
	for i := 0; i < configValue.NumField(); i++ {
		field := configValue.Field(i)

		fieldKey := strings.ToUpper(configValue.Type().Field(i).Tag.Get("yaml"))
		envKey := strings.ToUpper(fieldKey)
		envValue := os.Getenv(envKey)
		if envValue != "" {
			field.SetString(envValue)
		}

		if field.String() == "" {
			errorFields = append(errorFields, fieldKey)
		}
	}
	if len(errorFields) > 0 {
		log.Fatal("Error:", "empty fields:", errorFields)
	}

	return config
}

var config = readConfig("config.yaml")

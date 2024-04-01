package config

import (
	"log"
	"os"

	"github_spider/pkg/constants"
	"github_spider/pkg/model"

	"sigs.k8s.io/yaml"
)

func InitConfig() {

	configPath := constants.DefaultConfigFile

	yamlFile, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Failed to read YAML file: %v", err)
	}

	var config model.Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Failed to unmarshal YAML: %v", err)
	}

	log.Printf("config info: %v\n", config)

	model.SetConfig(config)

}

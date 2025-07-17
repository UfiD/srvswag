package config

import (
	"flag"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type AppFlags struct {
	ConfigPath string
}

func ParseFlag() AppFlags {
	configPath := flag.String("config", "", "Path to config")
	return AppFlags{
		ConfigPath: *configPath,
	}
}

type AppConfig struct {
	HTTPConfig         `yaml:"http"`
	RabbitMQPublisher  `yaml:"rabbitmq"`
	RabbitMQSubscriber `yaml:"rabbitmq"`
}

func ParseConfig(path string, cfg *AppConfig) {
	if path == "" {
		log.Fatal()
	}
	buf, err := os.ReadFile(path)
	if err != nil {
		log.Fatal()
	}
	err = yaml.Unmarshal(buf, cfg)
	if err != nil {
		log.Fatal("yaml unmasrshal error")
	}
}

type RabbitMQPublisher struct {
	Host      string `yaml:"host"`
	Port      uint16 `yaml:"port"`
	QueueName string `yaml:"queue_publisher"`
}

type RabbitMQSubscriber struct {
	Host      string `yaml:"host"`
	Port      uint16 `yaml:"port"`
	QueueName string `yaml:"queue_subscriber"`
}

type HTTPConfig struct {
	Addr string `yaml:"address"`
}

package main

import (
	"encoding/json"
	"io/ioutil"
)

/* Configuration for Rabbit Watcher */
type Config struct {
	RabbitMQ RabbitMQConfig `json:"rabbitmq"`
	InfluxDB InfluxDBConfig `json:"influxdb"`
}

/* RabbitMQ Configuration */
type RabbitMQConfig struct {
	Url       string `json:"url"`
	User      string `json:"user"`
	Password  string `json:"password"`
	VerifySsl bool   `json:"verifySSL"`
}

/* InfluxDB configuration */
type InfluxDBConfig struct {
	Url string `json:"url"`
}

func GetConfig(data []byte) (Config, error) {
	var v = Config{}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return v, err
	}
	return v, nil
}

func LoadConfig(filename string) (Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return Config{}, err
	}
	return GetConfig(data)
}

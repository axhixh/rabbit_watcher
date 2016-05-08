package main

import (
	"log"
	"net/url"
	"os"
)

func main() {
	var configFile string
	if len(os.Args) > 1 {
		configFile = os.Args[1]
	} else {
		configFile = "config.json"
	}
	config, err := LoadConfig(configFile)
	if err != nil {
		log.Println("Unable to load configuration.")
		panic(err)
	}
	log.Println(config)
	buf, err := get(config.RabbitMQ, "/api/connections")
	if err != nil {
		panic(err)
	}
	rabbitmq, err := url.Parse(config.RabbitMQ.Url)
	if err != nil {
		panic(err)
	}
	connections := Connection{
		Name:  rabbitmq.Host,
		Count: connections(buf),
	}
	log.Println(connections)
	sendUdp(config.InfluxDB, connections)
	buf, err = get(config.RabbitMQ, "/api/nodes")
	if err != nil {
		panic(err)
	}

	nodes := nodes(buf)
	for _, node := range nodes {
		log.Println(node)
		sendUdp(config.InfluxDB, node)
	}
}

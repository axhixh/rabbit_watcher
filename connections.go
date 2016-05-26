package main

import "fmt"

/* Represents Connections of a RabbitMQ server or cluster */
type Connections struct {
	Name  string
	Count int
}

func (c Connections) String() string {
	return fmt.Sprintf("Connection: %d", c.Count)
}

func (c Connections) SeriesName() string {
	return "rabbitmq_connections"
}

func (c Connections) Tags() map[string]string {
	return map[string]string{"node": c.Name}
}

func (c Connections) Fields() map[string]interface{} {
	return map[string]interface{}{
		"connections": c.Count,
	}
}

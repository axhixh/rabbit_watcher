package main

import "fmt"

/* Represents Connections of a RabbitMQ server or cluster */
type Connection struct {
	Name  string
	Count int
}

func (c Connection) String() string {
	return fmt.Sprintf("Connection: %d", c.Count)
}

func (c Connection) SeriesName() string {
	return "rabbitmq_connections"
}

func (c Connection) Tags() map[string]string {
	return map[string]string{"node": c.Name}
}

func (c Connection) Fields() map[string]interface{} {
	return map[string]interface{}{
		"connections": c.Count,
	}
}

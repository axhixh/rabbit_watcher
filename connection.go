package main

import (
    "fmt"
    "time"
)

type Connection struct {
    User string `json:"user"`
    VHost string `json:"vhost"`
    ConnectedAt int64 `json:"connected_at"`
}

func (c Connection) String() string {
	return fmt.Sprintf("User: %s, VHost: %s, Connected At: %d, Duration: %v",
		c.User, c.VHost, c.ConnectedAt, getDuration(c.ConnectedAt))
}

func (c Connection) SeriesName() string {
	return "rabbitmq_connection_duration"
}

func (c Connection) Tags() map[string]string {
	return map[string]string{"user": c.User, "vhost": c.VHost}
}

// Use the JSON fields as the fields for InfluxDB
func (c Connection) Fields() map[string]interface{} {
	return map[string]interface{}{
		"connectioned_at": c.ConnectedAt,
        "duration": getDuration(c.ConnectedAt).Seconds(),
	}
}

func getDuration(timestamp int64) time.Duration {
    return time.Since(time.Unix(timestamp/1000, 0))
}

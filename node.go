package main

import "fmt"

/* Represents a RabbitMQ node and its metrics */
type Node struct {
	Name        string `json:"name"`
	MemUsed     uint   `json:"mem_used"`
	MemLimit    uint   `json:"mem_limit"`
	ProcUsed    uint   `json:"proc_used"`
	ProcTotal   uint   `json:"proc_total"`
	SocketUsed  uint   `json:"sockets_used"`
	SocketTotal uint   `json:"sockets_total"`
	FdUsed      uint   `json:"fd_used"`
	FdTotal     uint   `json:"fd_total"`
}

func (n Node) String() string {
	return fmt.Sprintf("Node: %s\nMem: %d/%d,\nProc: %d/%d\nSockets: %d/%d\nFD: %d/%d",
		n.Name, n.MemUsed, n.MemLimit, n.ProcUsed, n.ProcTotal,
		n.SocketUsed, n.SocketTotal, n.FdUsed, n.FdTotal)
}

func (n Node) SeriesName() string {
	return "rabbitmq_usage"
}

func (n Node) Tags() map[string]string {
	return map[string]string{"node": n.Name}
}

func (n Node) Fields() map[string]interface{} {
	return map[string]interface{}{
		"mem_used":      n.MemUsed,
		"mem_list":      n.MemLimit,
		"proc_used":     n.ProcUsed,
		"proc_total":    n.ProcTotal,
		"sockets_used":  n.SocketUsed,
		"sockets_total": n.SocketTotal,
		"fd_used":       n.FdUsed,
		"fd_total":      n.FdTotal,
	}
}

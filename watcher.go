package main

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func connections(data []byte) int {
	var v []interface{}
	err := json.Unmarshal(data, &v)
	if err != nil {
		log.Printf("Unable to parse connections %v", data)
		return -1
	}
	return len(v)
}

func nodes(data []byte) []Node {
	var v []Node
	err := json.Unmarshal(data, &v)
	if err != nil {
		log.Printf("Unable to parse node data %v", data)
		return make([]Node, 0)
	}
	return v
}

func get(config RabbitMQConfig, apiEndpoint string) (data []byte, err error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: !config.VerifySsl},
	}
	req, err := http.NewRequest("GET", config.Url+apiEndpoint, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(config.User, config.Password)
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

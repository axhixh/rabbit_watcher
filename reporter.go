package main

import (
	"time"

	"github.com/influxdata/influxdb/client/v2"
)

type Reportable interface {
	Tags() map[string]string
	Fields() map[string]interface{}
	SeriesName() string
}

func sendUdp(config InfluxDBConfig, report Reportable) {
	pt, err := client.NewPoint(report.SeriesName(), report.Tags(),
		report.Fields(), time.Now())
	if err != nil {
		panic(err.Error())
	}
	bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Precision: "s",
	})
	bp.AddPoint(pt)

	c, _ := client.NewUDPClient(client.UDPConfig{Addr: config.Url})
	c.Write(bp)
}

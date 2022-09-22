package main

import (
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	influxdb_manage "influxdb-go/components/influxdb"
	influxdb_conf "influxdb-go/conf/influxdb"
	"influxdb-go/service_centers"
	"time"
)

func main() {
	service_centers.Init()
	for {
		influxdbSvc := service_centers.Discover(influxdb_conf.Name).(*influxdb_manage.InfluxdbManage)
		client := influxdbSvc.Client
		writeAPI := client.WriteAPI(influxdbSvc.Conf.Org, influxdbSvc.Conf.Bucket)
		// write line protocol
		p := influxdb2.NewPoint("stat",
			map[string]string{"unit": "temperature"},
			map[string]interface{}{"avg": 24.5, "max": 45},
			time.Now())
		// write point asynchronously
		writeAPI.WritePoint(p)
		// create point using fluent style
		p = influxdb2.NewPointWithMeasurement("stat").
			AddTag("unit", "temperature").
			AddField("avg", 23.2).
			AddField("max", 45).
			SetTime(time.Now())
		// write point asynchronously
		writeAPI.WritePoint(p)
		// Flush writes
		writeAPI.Flush()
		time.Sleep(5 * time.Second)
		fmt.Sprintf("insert data to influxdb")
	}
}

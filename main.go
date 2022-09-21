package main

import (
	"fmt"
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
		writeAPI.WriteRecord(fmt.Sprintf("stat,unit=temperature avg=%f,max=%f", 23.5, 45.0))
		writeAPI.WriteRecord(fmt.Sprintf("stat,unit=temperature avg=%f,max=%f", 22.5, 45.0))
		writeAPI.WriteRecord(fmt.Sprintf("CPU load%f, Memory%f", 23.5, 45.0))
		writeAPI.WriteRecord(fmt.Sprintf("CPU load%f, Memory%f", 24.6, 46.0))
		writeAPI.WriteRecord(fmt.Sprintf("CPU load%f, Memory%f", 25.7, 46.0))
		// Flush writes
		writeAPI.Flush()
		time.Sleep(5 * time.Second)
		fmt.Sprintf("insert data to influxdb")
	}
}

package influxdb_manage

import (
	"context"
	"fmt"
	influxdb_conf "influxdb-go/conf/influxdb"
	"influxdb-go/service_centers"
	"sync"
	"time"
)

func HealthCheck(wg sync.WaitGroup) {
	// validate client connection health
	for {
		client := service_centers.Discover(influxdb_conf.Name).(*InfluxdbManage).Client
		if client != nil {
			_, err := client.Health(context.Background())
			if err != nil {
				fmt.Errorf("influxdb health is: %s", err.Error())
			}
			time.Sleep(time.Second * 5)
		}
	}
	wg.Done()
}

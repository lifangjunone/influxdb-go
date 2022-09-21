package influxdb_manage

import (
	"fmt"
	"github.com/influxdata/influxdb-client-go/v2"
	influxdb_conf "influxdb-go/conf/influxdb"
	"influxdb-go/service_centers"
	"sync"
)

var (
	confFile string = "/home/lifangjun/Desktop/go_study/influxdb-go/etc/demo.toml"
)

type InfluxdbManage struct {
	Client influxdb2.Client
	Conf   *influxdb_conf.Config
}

func (i *InfluxdbManage) Name() string {
	return influxdb_conf.Name
}

func (i *InfluxdbManage) Config() error {
	config, err := influxdb_conf.LoadConfigFromToml(confFile)
	if err != nil {
		fmt.Errorf("load influxdb config is error %s", err)
		return err
	}
	i.Conf = config
	return nil
}

func (i *InfluxdbManage) InitService() error {
	i.Client = influxdb2.NewClient(influxdb_conf.GetAddr(), influxdb_conf.GetConfig().Token)
	return nil
}

func (i *InfluxdbManage) HealthCheck(group sync.WaitGroup) {
	go HealthCheck(group)
}

func init() {
	// registry service to service center
	svc := &InfluxdbManage{}
	service_centers.Registry(svc)
}

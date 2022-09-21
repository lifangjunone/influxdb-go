package service_centers

import "sync"

var (
	servicesCenter map[string]Services = make(map[string]Services, 10)
)

type Services interface {
	Name() string
	Config() error
	InitService() error
	HealthCheck(group sync.WaitGroup)
}

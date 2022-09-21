package service_centers

import (
	"fmt"
)

func Registry(svc Services) {
	_, ok := servicesCenter[svc.Name()]
	if ok {
		fmt.Print("service registry yet, not need registry again!!!")
	}
	servicesCenter[svc.Name()] = svc
}

func Init() {
	for _, svc := range servicesCenter {
		svc.Config()
		svc.InitService()
		// svc.HealthCheck(wg)
	}
}

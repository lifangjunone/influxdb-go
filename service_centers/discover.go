package service_centers

func Discover(name string) Services {
	svc, ok := servicesCenter[name]
	if ok {
		return svc
	}
	return nil
}

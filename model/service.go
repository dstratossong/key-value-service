package model

import ()

type Service struct {
	Name       string
	Url        string
	Version    string
	Properties []string
}

var Services = make(map[string][]*Service)

func GetServices() map[string][]*Service {
	return Services
}

// Warning: No locking!
func RegisterService(service *Service) {
	list, found := Services[service.Name]
	if !found {
		Services[service.Name] = make([]*Service, 0)
	} else {
		for _, v := range list {
			if v.Url == service.Url {
				// Ignores request when duplicate found
				// TODO: Fail the request instead
				return
			}
		}
	}
	Services[service.Name] = append(Services[service.Name], service)
}

package ioc

import (
	"fmt"
)

type Controller interface {
	Init() error
	Name() string
}

type CONTROLLERSCONTAINER map[string]Controller

var controllers = make(CONTROLLERSCONTAINER, 64)

func Controllers() CONTROLLERSCONTAINER {
	return controllers
}

func (cc CONTROLLERSCONTAINER) Registry(sc Controller) error {
	if _, ok := cc[sc.Name()]; ok {
		return fmt.Errorf("the service controller for the %s unit already exists", sc.Name())
	}

	cc[sc.Name()] = sc
	return nil
}

func (cc CONTROLLERSCONTAINER) GetServiceController(scn string) any {
	if sc, ok := cc[scn]; ok {
		return sc
	}
	return fmt.Errorf("the service controller of unit %s does not exist", scn)
}

func (cc CONTROLLERSCONTAINER) Init() error {
	for scName, sc := range cc {
		err := sc.Init()
		if err != nil {
			return fmt.Errorf("initializing the service controller of unit %s failed", scName)
		}
	}

	return nil
}

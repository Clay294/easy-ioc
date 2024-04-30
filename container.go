package easyioc

import "fmt"

var container = NewDefaultContainer()

type DefaultContainter struct {
	Containter map[string]*ObjectsContainer
}

func NewDefaultContainer() *DefaultContainter {
	return &DefaultContainter{
		make(map[string]*ObjectsContainer),
	}
}

func (dc *DefaultContainter) registryObjectContainer(name string) *ObjectsContainer {
	if oc, ok := dc.Containter[name]; ok {
		return oc
	}
	dc.Containter[name] = newObjectsContainer()
	return dc.Containter[name]
}

func GetObjectsContainer(ocName string) *ObjectsContainer {
	if oc, ok := container.Containter[ocName]; ok {
		return oc
	}
	return nil
}

func RegistryObject(ocName string, o Object) error {
	name := container.registryObjectContainer(ocName).registryObject(o)
	if name == "" {
		return fmt.Errorf("the %s %s has registered", ocName, o.Name())
	}
	return nil
}

func GetObject(ocName, name string) (Object, error) {
	oc := GetObjectsContainer(ocName)
	if oc == nil {
		return nil, fmt.Errorf("the %s %s has not registered", ocName, name)
	}

	o := oc.getObject(name)
	if o == nil {
		return nil, fmt.Errorf("the %s %s has not registered", ocName, name)
	}

	return o, nil
}

func InitObjects() error {
	for ocName, oc := range container.Containter {
		if oc == nil {
			return fmt.Errorf("the %s container is empty", ocName)
		}

		err := oc.initObjects()
		if err != nil {
			return err
		}
	}

	return nil
}

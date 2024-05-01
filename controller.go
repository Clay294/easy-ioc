package easyioc

import "fmt"

const (
	ControllersName = "controller"
)

type Controller interface {
	Object
}

type ControllerImpl struct {
	ObjectImpl
}

func RegistryController(c Controller) error {
	return RegistryObject(ControllersName, c)
}

func GetController(name string) (Controller, error) {
	o, err := GetObject(ControllersName, name)
	if err != nil {
		return nil, err
	}
	return o.(Controller), nil
}

func InitControllers() error {
	oc := GetObjectsContainer(ControllersName)
	if oc == nil {
		return fmt.Errorf("the %s container is empty", ControllersName)
	}
	for _, cc := range oc.Containers {
		err := cc.Init()
		if err != nil {
			return err
		}
	}
	return nil
}

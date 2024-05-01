package easyioc

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

const (
	HandlersName = "handler"
)

type Handler interface {
	Object
	Registry(gin.IRouter)
}

type HandlerImpl struct {
	ObjectImpl
}

func (HandlerImpl) Registry(gin.IRouter) {

}

func RegistryHandler(h Handler) error {
	return RegistryObject(HandlersName, h)
}

func GetHandler(name string) (Handler, error) {
	o, err := GetObject(HandlersName, name)
	if err != nil {
		return nil, err
	}
	return o.(Handler), nil
}

func InitHandler(urlPrefix string, router gin.IRouter) error {
	oc := GetObjectsContainer(HandlersName)
	if oc == nil {
		return fmt.Errorf("the %s container is empty", ControllersName)
	}

	err := oc.initObjects()
	if err != nil {
		return err
	}

	for _, o := range oc.Containers {
		if h, ok := o.(Handler); ok {
			h.Registry(router.Group(urlPrefix))
		}
	}

	return nil
}

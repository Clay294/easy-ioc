package easyioc

import (
	"fmt"
)

const ControllerObjectsTableName = "controller"

type ControllerObjct interface {
	IocObject
}

type ControllerObjectImpl struct {
	IocObject
}

// InitGrpcIocObjects 初始化所有grpc对象
func InitControllerObjects() error {
	err := InitIocObjects()
	if err != nil {
		return err
	}
	return nil
}

// 按需创建GrpcIocObjects的map
func NewControllerObjectTable(controllerObject ControllerObjct) *IocObjectsContainers {
	if controllerObjectTable, ok := container.Containter[ControllerObjectsTableName]; ok {
		return controllerObjectTable
	}

	iocObjectsContainers := new(IocObjectsContainers)
	iocObjectsContainers.Containers = make(map[string]IocObject)

	return iocObjectsContainers
}

func RegistryControllerObject(controllerObject ControllerObjct) error {
	iocObjectsContainers := NewControllerObjectTable(controllerObject)
	if iocObjectsContainers.IsExists(controllerObject.Name()) {
		return fmt.Errorf("the controller ioc object already exists")
	}
	iocObjectsContainers.Containers[controllerObject.Name()] = controllerObject
	container.Containter[ControllerObjectsTableName] = iocObjectsContainers
	return nil
}

func GetControllerObject(name string) any {
	controllerObject := container.Containter[ControllerObjectsTableName].GetIocObject(name)
	if controllerObject == nil {
		return fmt.Errorf("the grpc ioc object does not exist")
	}

	return controllerObject
}

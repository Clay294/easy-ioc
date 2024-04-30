package easyioc

import (
	"fmt"

	"google.golang.org/grpc"
)

const (
	GrpcIocOjectsTableName = "grpc"
)

// GrpcIocObject grpc对象的方法约束
type GrpcIocObject interface {
	IocObject
	Registry(*grpc.Server)
}

// GrpcIocObjImpl grpc模板对象
type GrpcIocObjImpl struct {
	IocObject
}

// Registry 向grpc服务器注册grpc服务接口
func (*GrpcIocObjImpl) Registry(*grpc.Server) {
}

// InitGrpcIocObjects 初始化所有grpc对象
func InitGrpcIocObjects(server *grpc.Server) error {
	err := InitIocObjects()
	if err != nil {
		return err
	}
	for _, grpcIocObject := range container.Containter[GrpcIocOjectsTableName].Containers {
		grpcIocObject.(GrpcIocObject).Registry(server)
	}
	return nil
}

// 按需创建GrpcIocObjects的map
func NewGrpcIocObjectsTable(grpcIocObject GrpcIocObject) *IocObjectsContainers {
	if grpcObjetsTable, ok := container.Containter[GrpcIocOjectsTableName]; ok {
		return grpcObjetsTable
	}
	return container.Containter[GrpcIocOjectsTableName]
}

func RegistryGrpcIocObject(grpcIocObject GrpcIocObject) error {
	grpcIocObjectsTable := NewGrpcIocObjectsTable(grpcIocObject)
	if grpcIocObjectsTable.IsExists(grpcIocObject.Name()) {
		return fmt.Errorf("the grpc ioc object already exists")
	}
	grpcIocObjectsTable.Containers[grpcIocObject.Name()] = grpcIocObject
	return nil
}

func GetGrpcIocObject(name string) any {
	grpcIocObject := container.Containter[GrpcIocOjectsTableName].GetIocObject(name)
	if grpcIocObject == nil {
		return fmt.Errorf("the grpc ioc object does not exist")
	}

	return grpcIocObject
}

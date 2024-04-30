package easyioc

import (
	"fmt"

	"google.golang.org/grpc"
)

const (
	GrpcServersName = "grpc_server"
)

// GrpcServer grpc对象的方法约束
type GrpcServer interface {
	Object
	Registry(*grpc.Server)
}

// GrpcServerImpl grpc模板对象
type GrpcServerImpl struct {
	Object
}

// Registry 向grpc服务器注册grpc服务接口
func (*GrpcServerImpl) Registry(*grpc.Server) {
}

func RegistryGrpcServer(gs GrpcServer) error {
	return RegistryObject(GrpcServersName, gs)
}

func GetGrpcServer(name string) (GrpcServer, error) {
	o, err := GetObject(GrpcServersName, name)
	if err != nil {
		return nil, err
	}

	return o.(GrpcServer), nil
}

func InitGrpcServers(server *grpc.Server) error {
	oc := GetObjectsContainer(GrpcServersName)
	if oc == nil {
		return fmt.Errorf("the %s container is empty", GrpcServersName)
	}

	err := oc.initObjects()
	if err != nil {
		return err
	}

	for _, o := range oc.Containers {
		if gs, ok := o.(GrpcServer); ok {
			gs.Registry(server)
		}
		return fmt.Errorf("the ioc object %s is not grpcserver", o.Name())
	}

	return nil
}

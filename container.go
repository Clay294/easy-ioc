package easyioc

var container = NewDefaultContainer()

type DefaultContainter struct {
	Containter map[string]*IocObjectsContainers
}

func NewDefaultContainer() *DefaultContainter {
	return &DefaultContainter{
		make(map[string]*IocObjectsContainers),
	}
}

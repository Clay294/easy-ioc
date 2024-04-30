package easyioc

type IocObject interface {
	Init() error
	Name() string
}

type IocObjectsContainers struct {
	Containers map[string]IocObject
}

type IocObjectsImpl struct {
}

func (*IocObjectsImpl) Init() error {
	return nil
}

func (*IocObjectsImpl) Name() string {
	return ""
}

func (iocs *IocObjectsContainers) IsExists(name string) bool {
	if _, ok := iocs.Containers[name]; ok {
		return true
	}
	return false
}

func (iocs *IocObjectsContainers) GetIocObject(name string) IocObject {
	if !iocs.IsExists(name) {
		return nil
	}

	return iocs.Containers[name]
}

func InitIocObjects() error {
	for _, iocObjects := range container.Containter {
		for _, iocObject := range iocObjects.Containers {
			err := iocObject.Init()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

package easyioc

type Object interface {
	Init() error
	Name() string
}

type ObjectsContainer struct {
	Containers map[string]Object
}

type ObjectImpl struct {
}

func (ObjectImpl) Init() error {
	return nil
}

func (ObjectImpl) Name() string {
	return ""
}

func (oc *ObjectsContainer) isExists(name string) bool {
	if _, ok := oc.Containers[name]; ok {
		return true
	}

	return false
}

func newObjectsContainer() *ObjectsContainer {
	return &ObjectsContainer{
		Containers: make(map[string]Object),
	}
}

func (oc *ObjectsContainer) getObject(name string) Object {
	if !oc.isExists(name) {
		return nil
	}

	return oc.Containers[name]
}

func (oc *ObjectsContainer) registryObject(o Object) string {
	if oc.getObject(o.Name()) != nil {
		return o.Name()
	}

	oc.Containers[o.Name()] = o
	return ""
}

func (oc *ObjectsContainer) initObjects() error {
	for _, o := range oc.Containers {
		err := o.Init()
		if err != nil {
			return err
		}
	}

	return nil
}

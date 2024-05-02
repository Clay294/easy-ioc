package controller_test

import (
	"testing"

	"github.com/Clay294/easyioc"
)

var impl = new(Impl)

type Impl struct {
	easyioc.ControllerImpl
}

func (i *Impl) Name() string {
	return "easyioc_controller"
}

func TestEasyiocController(t *testing.T) {
	err := easyioc.RegistryController(impl)
	if err != nil {
		t.Fatal(err)
	}

	err = easyioc.InitControllers()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", impl)
}

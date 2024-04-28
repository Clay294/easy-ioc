package ioc

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/url"
)

type Handler interface {
	Controller
	Registry(gin.IRouter)
}

type HANDLERSCONTAINER map[string]Handler

var handlers = make(HANDLERSCONTAINER, 64)

func Handlers() HANDLERSCONTAINER {
	return handlers
}

func (hc HANDLERSCONTAINER) Registry(ah Handler) error {
	if _, ok := hc[ah.Name()]; ok {
		return fmt.Errorf("the api handler of unit %s already exists", ah.Name())
	}

	hc[ah.Name()] = ah
	return nil
}

func (hc HANDLERSCONTAINER) Init(router gin.IRouter, apiUrl string) error {
	for ahName, ah := range hc {
		err := ah.Init()
		if err != nil {
			return fmt.Errorf("initializing the api handler of unit %s failed: %s", ahName, err)
		}

		unitUrl, err := url.JoinPath(apiUrl, ahName)
		if err != nil {
			return fmt.Errorf("initializing the api handler of unit %s failed: %s", ahName, err)
		}

		ah.Registry(router.Group(unitUrl))
	}
	return nil
}
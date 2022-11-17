package framework

import "net/http"

type Core struct {
}

func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	//TODO implement me
	panic("implement me")
}

func NewCore() *Core {
	return &Core{}
}

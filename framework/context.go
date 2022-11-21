package framework

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"
)

type Context struct {
	request        *http.Request
	responseWriter http.ResponseWriter
	ctx            context.Context
	handlers       []ControllerHandler
	index          int //当前请求调用到调用链的哪个节点
	// 是否超时标记位
	hasTimeout bool
	// 写保护机制
	writerMux *sync.Mutex
}

func (ctx *Context) BaseContext() context.Context {
	return ctx.request.Context()
}

func (ctx *Context) Done() <-chan struct{} {
	return ctx.BaseContext().Done()
}

func NewContext(r *http.Request, w http.ResponseWriter) *Context {
	return &Context{
		request:        r,
		responseWriter: w,
		ctx:            r.Context(),
		index:          -1,
		writerMux:      &sync.Mutex{},
	}
}

func (ctx *Context) WriterMux() *sync.Mutex {
	return ctx.WriterMux()
}

func (ctx *Context) GetRequest() *http.Request {
	return ctx.request
}
func (ctx *Context) GetResponse() http.ResponseWriter {
	return ctx.responseWriter
}
func (ctx *Context) SetHandlers(handlers []ControllerHandler) {
	ctx.handlers = handlers
}
func (ctx *Context) SetHasTimeout() {
	ctx.hasTimeout = true
}

func (ctx *Context) HasTimeout() bool {
	return ctx.hasTimeout
}
func (ctx *Context) Json(status int, obj interface{}) error {
	if ctx.HasTimeout() {
		return nil
	}
	ctx.responseWriter.Header().Set("Content-Type", "application/json")
	ctx.responseWriter.WriteHeader(status)
	byt, err := json.Marshal(obj)
	if err != nil {
		ctx.responseWriter.WriteHeader(500)
		return err
	}
	_, err = ctx.responseWriter.Write(byt)
	if err != nil {
		return err
	}
	return nil
}

func (ctx *Context) Next() error {
	ctx.index++
	if ctx.index < len(ctx.handlers) {
		if err := ctx.handlers[ctx.index](ctx); err != nil {
			return err
		}
	}
	return nil
}

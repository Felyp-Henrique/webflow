package webflow

import "Felyp-Henrique/webflow/engine"

type Context struct {
	_engine engine.Engine
}

func (c *Context) Parent() engine.ElementHandler {
	return c._engine.Handler().Parent()
}

func (c *Context) Window() engine.WindowHandler {
	return c._engine.Window()
}

func (c *Context) Engine() engine.Engine {
	return c._engine
}

func NewContext(engine engine.Engine) *Context {
	return &Context{
		_engine: engine,
	}
}

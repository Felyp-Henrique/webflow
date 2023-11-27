package webflow

import "Felyp-Henrique/webflow/engine"

type WebflowEntryPoint func(*Context)

type Webflow struct {
	_engine engine.Engine
}

func (w *Webflow) Start(entryPoint WebflowEntryPoint) {
	var (
		context *Context = NewContext(w._engine)
	)
	entryPoint(context)
}

func NewWebflow(engine engine.Engine) *Webflow {
	return &Webflow{
		_engine: engine,
	}
}

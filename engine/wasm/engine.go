package wasm

import (
	"Felyp-Henrique/webflow/engine"
	"syscall/js"
)

type Engine struct {
	_factory engine.ElementFactory
	_handler *engine.Stack[engine.ElementHandler]
}

func (e *Engine) GetFactory() engine.ElementFactory {
	return e._factory
}

func (e *Engine) GetHandler() engine.ElementHandler {
	return e._handler.Peek()
}

func (e *Engine) Append(handler engine.ElementHandler) {
	e._handler.Peek().Append(handler)
}

func New() *Engine {
	return &Engine{
		_factory: &ElementFactory{
			_document: js.Global().Get("document"),
		},
		_handler: engine.NewStack[engine.ElementHandler](),
	}
}

type ElementHandler struct {
	_value js.Value
}

func (e *ElementHandler) SetId(id string) {
	e._value.Set("id", id)
}

func (e *ElementHandler) SetClass(class string) {
	e._value.Set("class", class)
}

func (e *ElementHandler) SetContent(content string) {
	e._value.Set("innerText", content)
}

type ElementFactory struct {
	_document js.Value
}

func (e *ElementFactory) Create(el string) engine.ElementHandler {
	return &ElementHandler{
		_value: e._document.Call("createElement", el),
	}
}

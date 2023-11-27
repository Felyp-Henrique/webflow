package engine

type ActionScope = func(ElementHandler)

type Engine interface {
	Factory() ElementFactory
	Handler() ElementHandler
	Window() WindowHandler
	ActionInlineCreate(string) ElementHandler
	ActionBlockCreate(string, ActionScope) ElementHandler
}

type ElementHandler interface {
	SetId(string)
	SetClass(string)
	SetContent(string)
	Parent() ElementHandler
}

type ElementFactory interface {
	Create(string) ElementHandler
	CreateChild(ElementHandler, string) ElementHandler
}

type WindowHandler interface {
	Alert(string)
}

package engine

type Engine interface {
	GetFactory() ElementFactory
	GetHandler() ElementHandler
	Append(ElementHandler)
}

type ElementHandler interface {
	SetId(string)
	SetClass(string)
	SetContent(string)
	Append(ElementHandler)
}

type ElementFactory interface {
	Create(string) ElementHandler
}

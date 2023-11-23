package dom

import "syscall/js"

func Win() *Window {
	return NewWindow()
}

func Doc() *Document {
	return NewDocument()
}

type ElOption func(*Element)

func WithId(id string) ElOption {
	return func(e *Element) {
		e.SetId(id)
	}
}

func WithClass(class string) ElOption {
	return func(e *Element) {
		e.SetClass(class)
	}
}

func WithAttr(name string, value any) ElOption {
	return func(e *Element) {
		e.SetAttr(name, value)
	}
}

func WithContent(content any) ElOption {
	return func(e *Element) {
		e.SetContent(content)
	}
}

func WithEvent(event string, handler js.Func) ElOption {
	return func(e *Element) {
		e.AddEvent(event, handler)
	}
}

func WithChildren(children ...*Element) ElOption {
	return func(e *Element) {
		for _, child := range children {
			e.Add(child.Object)
		}
	}
}

func El(el string, options ...ElOption) *Element {
	var (
		document *Document
		element  *Element
	)
	document = NewDocument()
	element = NewElement(document.Create(el))
	for _, option := range options {
		option(element)
	}
	return element
}

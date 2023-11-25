package dom

import "container/list"

var (
	_document      *Document
	_elementsRoot  *Element
	_elementsStack *list.List
)

type Attr map[string]any

type Scope func(*Element)

func El(name string, attrs Attr, scope Scope) {
	element := _document.Create(name)
	parent := _elementsStack.Back().Value.(*Element)
	for attr, value := range attrs {
		switch attr {
		case "id":
			element.SetId(value.(string))
			continue
		case "class":
			element.SetClass(value.(string))
			continue
		default:
			element.SetAttr(name, value)
		}
	}
	parent.Add(*element)
	_elementsStack.PushBack(element)
	scope(element)
	_elementsStack.Remove(_elementsStack.Back())
}

func init() {
	_document = NewDocument()
	_elementsRoot = _document.Get("div#app")
	_elementsStack = list.New()
	_elementsStack.PushBack(_elementsRoot)
}

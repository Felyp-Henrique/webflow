package components

import (
	"Felyp-Henrique/webflow/dom"
	"syscall/js"
)

func Render(selector string, root Component) js.Value {
	object := dom.Doc().Get(selector)
	object.Add(root.Render().Object)
	return object.GetValue()
}

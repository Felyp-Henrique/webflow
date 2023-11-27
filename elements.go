package webflow

import "Felyp-Henrique/webflow/engine"

func Text(ctx *Context, text string) {
	p := ctx.Engine().ActionInlineCreate("p")
	p.SetContent(text)
}

func Container(ctx *Context, scope engine.ActionScope) {
	ctx.Engine().ActionBlockCreate("div", scope)
}

func Button(ctx *Context, text string, action func()) {
	ctx.Engine().ActionBlockCreate("button", func(eh engine.ElementHandler) {
		eh.SetContent(text)
	})
}

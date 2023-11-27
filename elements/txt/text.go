package txt

import "Felyp-Henrique/webflow"

func El(ctx *webflow.Context, el string, content string) {
	el_ := ctx.Engine().ActionInlineCreate(el)
	el_.SetContent(content)
}

func Text(ctx *webflow.Context, text string) {
	p := ctx.Engine().ActionInlineCreate("p")
	p.SetContent(text)
}

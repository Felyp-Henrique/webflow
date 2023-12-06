package els

import "Felyp-Henrique/webflow"

func Inline(ctx *webflow.Context, el string, content string) {
	el_ := ctx.Engine().ActionInlineCreate(el)
	el_.SetContent(content)
}

type TitleLevel string

const (
	Title1 TitleLevel = "h1"
	Title2 TitleLevel = "h2"
	Title3 TitleLevel = "h3"
	Title4 TitleLevel = "h4"
	Title5 TitleLevel = "h5"
)

func Title(ctx *webflow.Context, level TitleLevel, text string) {
	h := ctx.Engine().ActionInlineCreate(string(level))
	h.SetContent(text)
}

func Text(ctx *webflow.Context, text string) {
	p := ctx.Engine().ActionInlineCreate("p")
	p.SetContent(text)
}

func Span(ctx *webflow.Context, content string) {
	span := ctx.Engine().ActionInlineCreate("span")
	span.SetContent(content)
}

package els

import (
	"Felyp-Henrique/webflow"
	"Felyp-Henrique/webflow/engine"
)

func Button(ctx *webflow.Context, text string, action func()) {
	ctx.Engine().ActionBlockCreate("button", func(eh engine.ElementHandler) {
		eh.SetContent(text)
	})
}

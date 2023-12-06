package els

import (
	"Felyp-Henrique/webflow"
	"Felyp-Henrique/webflow/engine"
)

func Block(ctx *webflow.Context, el string, scope engine.Scope) {
	ctx.Engine().ActionBlockCreate(el, scope)
}

func Section(ctx *webflow.Context, scope engine.Scope) {
	ctx.Engine().ActionBlockCreate("section", scope)
}

func Container(ctx *webflow.Context, scope engine.Scope) {
	ctx.Engine().ActionBlockCreate("div", scope)
}

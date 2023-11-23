package components

import "Felyp-Henrique/webflow/dom"

type Component interface {
	Render() *dom.Element
}

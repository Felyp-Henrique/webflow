package engines

type Style map[string]any

type Attributes map[string]any

type Element struct {
	Id         string
	Class      string
	Attributes Attributes
	Style      Style
	Elements   []*Element
}

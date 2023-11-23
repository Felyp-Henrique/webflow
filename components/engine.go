package components

import "fmt"

func Render(el string, root Component) {
	fmt.Println(root.Render())
}

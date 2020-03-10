package svg

import (
	"encoding/xml"
)

// Group represents a G SVG element
// See: https://developer.mozilla.org/en-US/docs/Web/SVG/Element/g
type Group struct {
	XMLName  xml.Name
	Children []interface{}
}

// NewGroup constructs new Group element
func NewGroup(children ...interface{}) Group {
	g := Group{
		XMLName: xml.Name{Local: "g"},
	}

	g.Children = append(g.Children, children...)

	return g
}

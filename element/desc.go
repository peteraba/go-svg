package element

import (
	"encoding/xml"
)

// Desc represents a Desc SVG element
// See: https://developer.mozilla.org/en-US/docs/Web/SVG/Element/desc
type Desc struct {
	XMLName  xml.Name
	Text     string `xml:",innerxml"`
	Children []interface{}
}

// NewDesc constructs new Desc element
func NewDesc(text string, children ...interface{}) Desc {
	ts := Desc{
		XMLName: xml.Name{Local: "desc"},
		Text:    text,
	}

	ts.Children = append(ts.Children, children...)

	return ts
}

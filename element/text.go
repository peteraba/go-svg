package element

import (
	"encoding/xml"

	"../attribute"
)

// Text represents a Text SVG element
// See: https://developer.mozilla.org/en-US/docs/Web/SVG/Element/text
type Text struct {
	XMLName    xml.Name
	X          attribute.Length      `xml:"x,attr,omitempty"`
	Y          attribute.Length      `xml:"y,attr,omitempty"`
	TextAnchor *attribute.TextAnchor `xml:"text-anchor,attr,omitempty"`
	Fill       *attribute.Color      `xml:"stroke,attr,omitempty"`
	Children   []interface{}
}

// T constructs new Text element (shortcut)
func T(x, y float64, children ...interface{}) Text {
	return NewText(
		attribute.Length{Number: x},
		attribute.Length{Number: y},
		children...,
	)
}

// NewText constructs new Text element
func NewText(x, y attribute.Length, children ...interface{}) Text {
	t := Text{
		XMLName: xml.Name{Local: "text"},
		X:       x,
		Y:       y,
	}

	t.Children = append(t.Children, children...)

	return t
}

// SetFill sets the fill color of a Text
func (t Text) SetFill(fill attribute.Color) Text {
	t.Fill = &fill

	return t
}

// UnsetFill removes the previously set fill color of a Text
func (t Text) UnsetFill() Text {
	t.Fill = nil

	return t
}

// SetTextAnchor sets the text anchor of a Text
func (t Text) SetTextAnchor(ta attribute.TextAnchor) Text {
	t.TextAnchor = &ta

	return t
}

// UnsetFill removes the previously set text anchor of a Text
func (t Text) UnsetTextAnchor() Text {
	t.TextAnchor = nil

	return t
}

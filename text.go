package svg

import (
	"encoding/xml"
)

// Text represents a Text SVG element
// See: https://developer.mozilla.org/en-US/docs/Web/SVG/Element/text
type Text struct {
	XMLName    xml.Name
	X          *Length     `xml:"x,attr,omitempty"`
	Y          *Length     `xml:"y,attr,omitempty"`
	TextAnchor *TextAnchor `xml:"text-anchor,attr,omitempty"`
	Fill       *Color      `xml:"stroke,attr,omitempty"`
	Children   []interface{}
}

// T constructs new Text element (shortcut)
func T(x, y float64, children ...interface{}) Text {
	var (
		pX *Length
		pY *Length
	)

	if x != 0.0 {
		pX = &Length{Number: x}
	}

	if y != 0.0 {
		pY = &Length{Number: y}
	}

	return NewText(
		pX,
		pY,
		children...,
	)
}

// NewText constructs new Text element
func NewText(x, y *Length, children ...interface{}) Text {
	t := Text{
		XMLName: xml.Name{Local: "text"},
		X:       x,
		Y:       y,
	}

	t.Children = append(t.Children, children...)

	return t
}

// SetFill sets the fill color of a Text
func (t Text) SetFill(fill Color) Text {
	t.Fill = &fill

	return t
}

// UnsetFill removes the previously set fill color of a Text
func (t Text) UnsetFill() Text {
	t.Fill = nil

	return t
}

// SetTextAnchor sets the text anchor of a Text
func (t Text) SetTextAnchor(ta TextAnchor) Text {
	t.TextAnchor = &ta

	return t
}

// UnsetFill removes the previously set text anchor of a Text
func (t Text) UnsetTextAnchor() Text {
	t.TextAnchor = nil

	return t
}

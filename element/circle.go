package element

import (
	"encoding/xml"

	"../attribute"
)

// Circle represents a Circle SVG element
// See: https://developer.mozilla.org/en-US/docs/Web/SVG/Element/circle
type Circle struct {
	XMLName     xml.Name
	CX          attribute.Length `xml:"cx,attr,omitempty"`
	CY          attribute.Length `xml:"cy,attr,omitempty"`
	R           attribute.Length `xml:"r,attr,omitempty"`
	Stroke      *attribute.Color `xml:"stroke,attr,omitempty"`
	StrokeWidth *uint8           `xml:"stroke-width,attr,omitempty"`
	Children    []interface{}
}

// C constructs new Circle element (shortcut)
func C(cx, cy, r float64, children ...interface{}) Circle {
	return NewCircle(
		attribute.Length{Number: cx},
		attribute.Length{Number: cy},
		attribute.Length{Number: r},
		children...,
	)
}

// NewCircle constructs new Circle element
func NewCircle(cx, cy, r attribute.Length, children ...interface{}) Circle {
	c := Circle{
		XMLName: xml.Name{Local: "circle"},
		CX:      cx,
		CY:      cy,
		R:       r,
	}

	c.Children = append(c.Children, children...)

	return c
}

// SetStrokeWidth sets the stroke width of a Circle
func (c Circle) SetStrokeWidth(strokeWidth uint8) Circle {
	c.StrokeWidth = &strokeWidth

	return c
}

// UnsetStrokeWidth removes the previously set stroke width of a Circle
func (c Circle) UnsetStrokeWidth() Circle {
	c.StrokeWidth = nil

	return c
}

// SetStroke sets the stroke color of a Circle
func (c Circle) SetStroke(stroke attribute.Color) Circle {
	c.Stroke = &stroke

	return c
}

// UnsetStroke removes the previously set stroke color of a Circle
func (c Circle) UnsetStroke() Circle {
	c.Stroke = nil

	return c
}

package element

import (
	"encoding/xml"

	"../attribute"
)

// Ellipse represents a Ellipse SVG element
// See: https://developer.mozilla.org/en-US/docs/Web/SVG/Element/ellipse
type Ellipse struct {
	XMLName     xml.Name
	CX          attribute.Length `xml:"cx,attr,omitempty"`
	CY          attribute.Length `xml:"cy,attr,omitempty"`
	RX          attribute.Length `xml:"rx,attr,omitempty"`
	RY          attribute.Length `xml:"ry,attr,omitempty"`
	Stroke      *attribute.Color `xml:"stroke,attr,omitempty"`
	StrokeWidth *uint8           `xml:"stroke-width,attr,omitempty"`
	Children    []interface{}
}

// E constructs new Ellipse element (shortcut)
func E(cx, cy, rx, ry float64, children ...interface{}) Ellipse {
	return NewEllipse(
		attribute.Length{Number: cx},
		attribute.Length{Number: cy},
		attribute.Length{Number: rx},
		attribute.Length{Number: ry},
		children...,
	)
}

// NewEllipse constructs new Ellipse element
func NewEllipse(cx, cy, rx, ry attribute.Length, children ...interface{}) Ellipse {
	c := Ellipse{
		XMLName: xml.Name{Local: "ellipse"},
		CX:      cx,
		CY:      cy,
		RX:      rx,
		RY:      ry,
	}

	c.Children = append(c.Children, children...)

	return c
}

// SetStrokeWidth sets the stroke width of a Ellipse
func (e Ellipse) SetStrokeWidth(strokeWidth uint8) Ellipse {
	e.StrokeWidth = &strokeWidth

	return e
}

// UnsetStrokeWidth removes the previously set stroke width of a Ellipse
func (e Ellipse) UnsetStrokeWidth() Ellipse {
	e.StrokeWidth = nil

	return e
}

// SetStroke sets the stroke color of a Ellipse
func (e Ellipse) SetStroke(stroke attribute.Color) Ellipse {
	e.Stroke = &stroke

	return e
}

// UnsetStroke removes the previously set stroke color of a Ellipse
func (e Ellipse) UnsetStroke() Ellipse {
	e.Stroke = nil

	return e
}

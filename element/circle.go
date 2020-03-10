package element

import (
	"encoding/xml"

	"../attribute"
)

// Circle represents a Circle SVG element
// See: https://developer.mozilla.org/en-US/docs/Web/SVG/Element/circle
type Circle struct {
	XMLName       xml.Name
	CX            attribute.Length   `xml:"cx,attr,omitempty"`
	CY            attribute.Length   `xml:"cy,attr,omitempty"`
	R             attribute.Length   `xml:"r,attr,omitempty"`
	Stroke        *attribute.Color   `xml:"stroke,attr,omitempty"`
	StrokeWidth   *uint8             `xml:"stroke-width,attr,omitempty"`
	StrokeOpacity *attribute.Opacity `xml:"stroke-opacity,attr,omitempty"`
	Fill          *attribute.Color   `xml:"fill,attr,omitempty"`
	FillOpacity   *attribute.Opacity `xml:"fill-opacity,attr,omitempty"`
	Opacity       float64            `xml:"opacity,attr,omitempty"`
	Children      []interface{}
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

// SetStrokeOpacity sets the stroke opacity of a Circle
func (c Circle) SetStrokeOpacity(so attribute.Opacity) Circle {
	c.StrokeOpacity = &so

	return c
}

// SetStrokeOpacity removes the stroke opacity of a Circle
func (c Circle) UnsetStrokeOpacity() Circle {
	c.StrokeOpacity = nil

	return c
}

// SetStroke sets the fill color of a Circle
func (c Circle) SetFill(fill attribute.Color) Circle {
	c.Fill = &fill

	return c
}

// UnsetStroke removes the previously set fill color of a Circle
func (c Circle) UnsetFill() Circle {
	c.Fill = nil

	return c
}

// SetStrokeOpacity sets the fill opacity of a Circle
func (c Circle) SetFillOpacity(fo attribute.Opacity) Circle {
	c.FillOpacity = &fo

	return c
}

// SetStrokeOpacity removes the stroke opacity of a Circle
func (c Circle) UnsetFillOpacity() Circle {
	c.FillOpacity = nil

	return c
}

// SetOpacity sets the opacity of a Circle
func (c Circle) SetOpacity(o float64) Circle {
	if o < 0 {
		c.Opacity = 0
	} else if o > 1 {
		c.Opacity = 1
	} else {
		c.Opacity = o
	}

	return c
}

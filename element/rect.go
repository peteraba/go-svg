package element

import (
	"encoding/xml"

	"../attribute"
)

// Rect represents a Rect SVG element
// See: https://developer.mozilla.org/en-US/docs/Web/SVG/Element/rect
type Rect struct {
	XMLName       xml.Name
	X             attribute.Length   `xml:"x,attr,omitempty"`
	Y             attribute.Length   `xml:"y,attr,omitempty"`
	Width         attribute.Length   `xml:"width,attr,omitempty"`
	Height        attribute.Length   `xml:"height,attr,omitempty"`
	RX            attribute.Length   `xml:"rx,attr,omitempty"`
	RY            attribute.Length   `xml:"ry,attr,omitempty"`
	Stroke        *attribute.Color   `xml:"stroke,attr,omitempty"`
	StrokeWidth   *uint8             `xml:"stroke-width,attr,omitempty"`
	StrokeOpacity *attribute.Opacity `xml:"stroke-opacity,attr,omitempty"`
	Fill          *attribute.Color   `xml:"fill,attr,omitempty"`
	FillOpacity   *attribute.Opacity `xml:"fill-opacity,attr,omitempty"`
	Opacity       float64            `xml:"opacity,attr,omitempty"`
	Children      []interface{}
}

// R constructs new Rect element (shortcut)
func R(x, y, width, height float64, children ...interface{}) Rect {
	return NewRect(
		attribute.Length{Number: x},
		attribute.Length{Number: y},
		attribute.Length{Number: width},
		attribute.Length{Number: height},
		attribute.Length{},
		attribute.Length{},
		children...,
	)
}

// NewRect constructs new Rect element
func NewRect(x, y, width, height, rx, ry attribute.Length, children ...interface{}) Rect {
	c := Rect{
		XMLName: xml.Name{Local: "rect"},
		X:       x,
		Y:       y,
		Width:   width,
		Height:  height,
		RX:      rx,
		RY:      ry,
	}

	c.Children = append(c.Children, children...)

	return c
}

// SetStrokeWidth sets the stroke width of a Rect
func (r Rect) SetStrokeWidth(strokeWidth uint8) Rect {
	r.StrokeWidth = &strokeWidth

	return r
}

// UnsetStrokeWidth removes the previously set stroke width of a Rect
func (r Rect) UnsetStrokeWidth() Rect {
	r.StrokeWidth = nil

	return r
}

// SetStroke sets the stroke color of a Rect
func (r Rect) SetStroke(stroke attribute.Color) Rect {
	r.Stroke = &stroke

	return r
}

// UnsetStroke removes the previously set stroke color of a Rect
func (r Rect) UnsetStroke() Rect {
	r.Stroke = nil

	return r
}

// SetStrokeOpacity sets the stroke opacity of a Rect
func (r Rect) SetStrokeOpacity(so attribute.Opacity) Rect {
	r.StrokeOpacity = &so

	return r
}

// SetStrokeOpacity removes the stroke opacity of a Rect
func (r Rect) UnsetStrokeOpacity() Rect {
	r.StrokeOpacity = nil

	return r
}

// SetStroke sets the fill color of a Rect
func (r Rect) SetFill(fill attribute.Color) Rect {
	r.Fill = &fill

	return r
}

// UnsetStroke removes the previously set fill color of a Rect
func (r Rect) UnsetFill() Rect {
	r.Fill = nil

	return r
}

// SetStrokeOpacity sets the fill opacity of a Rect
func (r Rect) SetFillOpacity(fo attribute.Opacity) Rect {
	r.FillOpacity = &fo

	return r
}

// SetStrokeOpacity removes the stroke opacity of a Rect
func (r Rect) UnsetFillOpacity() Rect {
	r.FillOpacity = nil

	return r
}

// SetOpacity sets the opacity of a Rect
func (r Rect) SetOpacity(o float64) Rect {
	if o < 0 {
		r.Opacity = 0
	} else if o > 1 {
		r.Opacity = 1
	} else {
		r.Opacity = o
	}

	return r
}

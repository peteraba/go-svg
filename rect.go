package svg

import (
	"encoding/xml"
)

// Rect represents a Rect SVG element
// See: https://developer.mozilla.org/en-US/docs/Web/SVG/Element/rect
type Rect struct {
	XMLName       xml.Name
	X             Length   `xml:"x,attr,omitempty"`
	Y             Length   `xml:"y,attr,omitempty"`
	Width         Length   `xml:"width,attr,omitempty"`
	Height        Length   `xml:"height,attr,omitempty"`
	RX            Length   `xml:"rx,attr,omitempty"`
	RY            Length   `xml:"ry,attr,omitempty"`
	Stroke        *Color   `xml:"stroke,attr,omitempty"`
	StrokeWidth   *uint8   `xml:"stroke-width,attr,omitempty"`
	StrokeOpacity *Opacity `xml:"stroke-opacity,attr,omitempty"`
	Fill          *Color   `xml:"fill,attr,omitempty"`
	FillOpacity   *Opacity `xml:"fill-opacity,attr,omitempty"`
	Opacity       float64  `xml:"opacity,attr,omitempty"`
	Children      []interface{}
}

// R constructs new Rect element (shortcut)
func R(x, y, width, height float64, children ...interface{}) Rect {
	return NewRect(
		Length{Number: x},
		Length{Number: y},
		Length{Number: width},
		Length{Number: height},
		Length{},
		Length{},
		children...,
	)
}

// NewRect constructs new Rect element
func NewRect(x, y, width, height, rx, ry Length, children ...interface{}) Rect {
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
func (r Rect) SetStroke(stroke Color) Rect {
	r.Stroke = &stroke

	return r
}

// UnsetStroke removes the previously set stroke color of a Rect
func (r Rect) UnsetStroke() Rect {
	r.Stroke = nil

	return r
}

// SetStrokeOpacity sets the stroke opacity of a Rect
func (r Rect) SetStrokeOpacity(so Opacity) Rect {
	r.StrokeOpacity = &so

	return r
}

// SetStrokeOpacity removes the stroke opacity of a Rect
func (r Rect) UnsetStrokeOpacity() Rect {
	r.StrokeOpacity = nil

	return r
}

// SetStroke sets the fill color of a Rect
func (r Rect) SetFill(fill Color) Rect {
	r.Fill = &fill

	return r
}

// UnsetStroke removes the previously set fill color of a Rect
func (r Rect) UnsetFill() Rect {
	r.Fill = nil

	return r
}

// SetStrokeOpacity sets the fill opacity of a Rect
func (r Rect) SetFillOpacity(fo Opacity) Rect {
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

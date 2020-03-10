package svg

import (
	"encoding/xml"
)

// Ellipse represents a Ellipse SVG element
// See: https://developer.mozilla.org/en-US/docs/Web/SVG/Element/ellipse
type Ellipse struct {
	XMLName       xml.Name
	CX            Length   `xml:"cx,attr,omitempty"`
	CY            Length   `xml:"cy,attr,omitempty"`
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

// E constructs new Ellipse element (shortcut)
func E(cx, cy, rx, ry float64, children ...interface{}) Ellipse {
	return NewEllipse(
		Length{Number: cx},
		Length{Number: cy},
		Length{Number: rx},
		Length{Number: ry},
		children...,
	)
}

// NewEllipse constructs new Ellipse element
func NewEllipse(cx, cy, rx, ry Length, children ...interface{}) Ellipse {
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
func (e Ellipse) SetStroke(stroke Color) Ellipse {
	e.Stroke = &stroke

	return e
}

// UnsetStroke removes the previously set stroke color of a Ellipse
func (e Ellipse) UnsetStroke() Ellipse {
	e.Stroke = nil

	return e
}

// SetStrokeOpacity sets the stroke opacity of a Ellipse
func (e Ellipse) SetStrokeOpacity(so Opacity) Ellipse {
	e.StrokeOpacity = &so

	return e
}

// SetStrokeOpacity removes the stroke opacity of a Ellipse
func (e Ellipse) UnsetStrokeOpacity() Ellipse {
	e.StrokeOpacity = nil

	return e
}

// SetStroke sets the fill color of a Ellipse
func (e Ellipse) SetFill(fill Color) Ellipse {
	e.Fill = &fill

	return e
}

// UnsetStroke removes the previously set fill color of a Ellipse
func (e Ellipse) UnsetFill() Ellipse {
	e.Fill = nil

	return e
}

// SetStrokeOpacity sets the fill opacity of a Ellipse
func (e Ellipse) SetFillOpacity(fo Opacity) Ellipse {
	e.FillOpacity = &fo

	return e
}

// SetStrokeOpacity removes the stroke opacity of a Ellipse
func (e Ellipse) UnsetFillOpacity() Ellipse {
	e.FillOpacity = nil

	return e
}

// SetOpacity sets the opacity of a Ellipse
func (e Ellipse) SetOpacity(o float64) Ellipse {
	if o < 0 {
		e.Opacity = 0
	} else if o > 1 {
		e.Opacity = 1
	} else {
		e.Opacity = o
	}

	return e
}

package svg

import (
	"encoding/xml"
	"sync"
)

// Circle represents a Circle SVG element
// See: https://developer.mozilla.org/en-US/docs/Web/SVG/Element/circle
type Circle struct {
	XMLName       xml.Name
	CX            *Length    `xml:"cx,attr,omitempty"`
	CY            *Length    `xml:"cy,attr,omitempty"`
	R             *Length    `xml:"r,attr,omitempty"`
	Stroke        *Color     `xml:"stroke,attr,omitempty"`
	StrokeWidth   *uint8     `xml:"stroke-width,attr,omitempty"`
	StrokeOpacity *Opacity   `xml:"stroke-opacity,attr,omitempty"`
	Fill          *Color     `xml:"fill,attr,omitempty"`
	FillOpacity   *Opacity   `xml:"fill-opacity,attr,omitempty"`
	Opacity       float64    `xml:"opacity,attr,omitempty"`
	Attrs         []xml.Attr `xml:",attr"`
	Children      []interface{}
	lock          *sync.Mutex
}

// C constructs new Circle element (shortcut)
func C(cx, cy, r float64, children ...interface{}) Circle {
	var (
		pCx, pCy, pR *Length
	)

	if cx != 0.0 {
		pCx = &Length{Number: cx}
	}

	if cy != 0.0 {
		pCy = &Length{Number: cy}
	}

	if r != 0.0 {
		pR = &Length{Number: r}
	}

	return NewCircle(
		pCx,
		pCy,
		pR,
		children...,
	)
}

// NewCircle constructs new Circle element
func NewCircle(cx, cy, r *Length, children ...interface{}) Circle {
	c := Circle{
		XMLName: xml.Name{Local: "circle"},
		CX:      cx,
		CY:      cy,
		R:       r,
		lock:    &sync.Mutex{},
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
func (c Circle) SetStroke(stroke Color) Circle {
	c.Stroke = &stroke

	return c
}

// UnsetStroke removes the previously set stroke color of a Circle
func (c Circle) UnsetStroke() Circle {
	c.Stroke = nil

	return c
}

// SetStrokeOpacity sets the stroke opacity of a Circle
func (c Circle) SetStrokeOpacity(so Opacity) Circle {
	c.StrokeOpacity = &so

	return c
}

// SetStrokeOpacity removes the stroke opacity of a Circle
func (c Circle) UnsetStrokeOpacity() Circle {
	c.StrokeOpacity = nil

	return c
}

// SetStroke sets the fill color of a Circle
func (c Circle) SetFill(fill Color) Circle {
	c.Fill = &fill

	return c
}

// UnsetStroke removes the previously set fill color of a Circle
func (c Circle) UnsetFill() Circle {
	c.Fill = nil

	return c
}

// SetStrokeOpacity sets the fill opacity of a Circle
func (c Circle) SetFillOpacity(fo Opacity) Circle {
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

// AddAttr adds a new attribute to a Circle
func (c Circle) AddAttr(name, value string) Circle {
	c.lock.Lock()
	c.Attrs = append(c.Attrs, xml.Attr{Name: xml.Name{Local: name}, Value: value})
	c.lock.Unlock()

	return c
}

// RemoveAttr removes all attributes of a given name of a Circle
func (c Circle) RemoveAttr(name string) Circle {
	c.lock.Lock()
	attrs := []xml.Attr{}
	for _, attr := range c.Attrs {
		if attr.Name.Local != name {
			attrs = append(attrs, attr)
		}
	}
	c.Attrs = attrs
	c.lock.Unlock()

	return c
}

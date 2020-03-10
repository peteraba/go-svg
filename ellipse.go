package svg

import (
	"encoding/xml"
	"sync"
)

// Ellipse represents a Ellipse SVG element
// See: https://developer.mozilla.org/en-US/docs/Web/SVG/Element/ellipse
type Ellipse struct {
	XMLName       xml.Name
	CX            *Length    `xml:"cx,attr,omitempty"`
	CY            *Length    `xml:"cy,attr,omitempty"`
	RX            *Length    `xml:"rx,attr,omitempty"`
	RY            *Length    `xml:"ry,attr,omitempty"`
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

// El constructs new Ellipse element (shortcut)
func El(cx, cy, rx, ry float64, children ...interface{}) Ellipse {
	var (
		pCx, pCy, pRx, pRy *Length
	)

	if cx != 0.0 {
		pCx = &Length{Number: cx}
	}

	if cy != 0.0 {
		pCy = &Length{Number: cy}
	}

	if rx != 0.0 {
		pRx = &Length{Number: rx}
	}

	if ry != 0.0 {
		pRy = &Length{Number: ry}
	}

	return NewEllipse(
		pCx,
		pCy,
		pRx,
		pRy,
		children...,
	)
}

// NewEllipse constructs new Ellipse element
func NewEllipse(cx, cy, rx, ry *Length, children ...interface{}) Ellipse {
	c := Ellipse{
		XMLName: xml.Name{Local: "ellipse"},
		CX:      cx,
		CY:      cy,
		RX:      rx,
		RY:      ry,
		lock:    &sync.Mutex{},
	}

	c.Children = append(c.Children, children...)

	return c
}

// SetStrokeWidth sets the stroke width of a Ellipse
func (el Ellipse) SetStrokeWidth(strokeWidth uint8) Ellipse {
	el.StrokeWidth = &strokeWidth

	return el
}

// UnsetStrokeWidth removes the previously set stroke width of a Ellipse
func (el Ellipse) UnsetStrokeWidth() Ellipse {
	el.StrokeWidth = nil

	return el
}

// SetStroke sets the stroke color of a Ellipse
func (el Ellipse) SetStroke(stroke Color) Ellipse {
	el.Stroke = &stroke

	return el
}

// UnsetStroke removes the previously set stroke color of a Ellipse
func (el Ellipse) UnsetStroke() Ellipse {
	el.Stroke = nil

	return el
}

// SetStrokeOpacity sets the stroke opacity of a Ellipse
func (el Ellipse) SetStrokeOpacity(so Opacity) Ellipse {
	el.StrokeOpacity = &so

	return el
}

// SetStrokeOpacity removes the stroke opacity of a Ellipse
func (el Ellipse) UnsetStrokeOpacity() Ellipse {
	el.StrokeOpacity = nil

	return el
}

// SetStroke sets the fill color of a Ellipse
func (el Ellipse) SetFill(fill Color) Ellipse {
	el.Fill = &fill

	return el
}

// UnsetStroke removes the previously set fill color of a Ellipse
func (el Ellipse) UnsetFill() Ellipse {
	el.Fill = nil

	return el
}

// SetStrokeOpacity sets the fill opacity of a Ellipse
func (el Ellipse) SetFillOpacity(fo Opacity) Ellipse {
	el.FillOpacity = &fo

	return el
}

// SetStrokeOpacity removes the stroke opacity of a Ellipse
func (el Ellipse) UnsetFillOpacity() Ellipse {
	el.FillOpacity = nil

	return el
}

// SetOpacity sets the opacity of a Ellipse
func (el Ellipse) SetOpacity(o float64) Ellipse {
	if o < 0 {
		el.Opacity = 0
	} else if o > 1 {
		el.Opacity = 1
	} else {
		el.Opacity = o
	}

	return el
}

// AddAttr adds a new attribute of a Ellipse
func (el Ellipse) AddAttr(name, value string) Ellipse {
	el.lock.Lock()
	el.Attrs = append(el.Attrs, xml.Attr{Name: xml.Name{Local: name}, Value: value})
	el.lock.Unlock()

	return el
}

// RemoveAttr removes all attributes of a given name of a Ellipse
func (el Ellipse) RemoveAttr(name string) Ellipse {
	el.lock.Lock()
	var attrs []xml.Attr
	for _, attr := range el.Attrs {
		if attr.Name.Local != name {
			attrs = append(attrs, attr)
		}
	}
	el.Attrs = attrs
	el.lock.Unlock()

	return el
}

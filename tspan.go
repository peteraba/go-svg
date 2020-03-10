package svg

import (
	"encoding/xml"
	"sync"
)

// TSpan represents a TSpan SVG element
// See: https://developer.mozilla.org/en-US/docs/Web/SVG/Element/tspan
type TSpan struct {
	XMLName  xml.Name
	X        *Length    `xml:"x,attr,omitempty"`
	Y        *Length    `xml:"y,attr,omitempty"`
	DX       *Length    `xml:"dx,attr,omitempty"`
	DY       *Length    `xml:"dy,attr,omitempty"`
	Attrs    []xml.Attr `xml:",attr"`
	Text     string     `xml:",innerxml"`
	Children []interface{}
	l        sync.Mutex
}

// TS constructs new TSpan element (shortcut)
func TS(text string, children ...interface{}) TSpan {
	return NewTSpan(text, children...)
}

// NewTSpan constructs new TSpan element
func NewTSpan(text string, children ...interface{}) TSpan {
	ts := TSpan{
		XMLName: xml.Name{Local: "tspan"},
		Text:    text,
	}

	ts.Children = append(ts.Children, children...)

	return ts
}

// SetX sets the X attribute of a TSpan
func (ts TSpan) SetX(x Length) TSpan {
	ts.X = &x

	return ts
}

// SetY sets the Y attribute of a TSpan
func (ts TSpan) SetY(y Length) TSpan {
	ts.Y = &y

	return ts
}

// SetDx sets the DX attribute of a TSpan
func (ts TSpan) SetDx(dx Length) TSpan {
	ts.DX = &dx

	return ts
}

// SetDy sets the DY attribute of a TSpan
func (ts TSpan) SetDy(dy Length) TSpan {
	ts.DY = &dy

	return ts
}

// SX sets the X attribute of a TSpan (shortcut)
func (ts TSpan) SX(x float64) TSpan {
	ts.X = &Length{Number: x}

	return ts
}

// SY sets the Y attribute of a TSpan (shortcut)
func (ts TSpan) SY(y float64) TSpan {
	ts.Y = &Length{Number: y}

	return ts
}

// SDx sets the DX attribute of a TSpan (shortcut)
func (ts TSpan) SDx(dx float64) TSpan {
	ts.DX = &Length{Number: dx}

	return ts
}

// SDy sets the DY attribute of a TSpan (shortcut)
func (ts TSpan) SDy(dy float64) TSpan {
	ts.DY = &Length{Number: dy}

	return ts
}

// UnsetDx removes the previously set DX attribute of a TSpan
func (ts TSpan) UnsetDx() TSpan {
	ts.DX = nil

	return ts
}

// UnsetDy removes the previously set DY attribute of a TSpan
func (ts TSpan) UnsetDy() TSpan {
	ts.DY = nil

	return ts
}

// AddAttr adds a new attribute of a TSpan
func (ts TSpan) AddAttr(name, value string) TSpan {
	ts.l.Lock()
	ts.Attrs = append(ts.Attrs, xml.Attr{Name: xml.Name{Local: name}, Value: value})
	ts.l.Unlock()

	return ts
}

// RemoveAttr removes all attributes of a given name of a TSpan
func (ts TSpan) RemoveAttr(name string) TSpan {
	ts.l.Lock()
	attrs := []xml.Attr{}
	for _, attr := range ts.Attrs {
		if attr.Name.Local != name {
			attrs = append(attrs, attr)
		}
	}
	ts.Attrs = attrs
	ts.l.Unlock()

	return ts
}

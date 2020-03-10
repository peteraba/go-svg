package svg

import (
	"encoding/xml"
	"sync"
)

type SVG struct {
	XMLName  xml.Name
	Width    float64    `xml:"width,attr,omitempty"`
	Height   float64    `xml:"height,attr,omitempty"`
	Version  string     `xml:"version,attr,omitempty"`
	Attrs    []xml.Attr `xml:",attr"`
	Children []interface{}
	l        sync.Mutex
}

func NewSVG(width, height float64, children ...interface{}) SVG {
	s := SVG{
		XMLName: xml.Name{Space: "http://www.w3.org/2000/svg", Local: "svg"},
		Width:   width,
		Height:  height,
		Version: "1.1",
	}

	s.Children = append(s.Children, children...)

	return s
}

// AddAttr adds a new attribute of an SVG tag
func (s SVG) AddAttr(name, value string) SVG {
	s.l.Lock()
	s.Attrs = append(s.Attrs, xml.Attr{Name: xml.Name{Local: name}, Value: value})
	s.l.Unlock()

	return s
}

// RemoveAttr removes all attributes of a given name of an SVG tag
func (s SVG) RemoveAttr(name string) SVG {
	s.l.Lock()
	var attrs []xml.Attr
	for _, attr := range s.Attrs {
		if attr.Name.Local != name {
			attrs = append(attrs, attr)
		}
	}
	s.Attrs = attrs
	s.l.Unlock()

	return s
}

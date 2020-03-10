package svg

import (
	"encoding/xml"
	"sync"
)

// Group represents a G SVG element
// See: https://developer.mozilla.org/en-US/docs/Web/SVG/Element/g
type Group struct {
	XMLName  xml.Name
	Attrs    []xml.Attr `xml:",attr"`
	Children []interface{}
	l        sync.Mutex
}

// NewGroup constructs new Group element
func NewGroup(children ...interface{}) Group {
	g := Group{
		XMLName: xml.Name{Local: "g"},
	}

	g.Children = append(g.Children, children...)

	return g
}

// AddAttr adds a new attribute of a Group
func (g Group) AddAttr(name, value string) Group {
	g.l.Lock()
	g.Attrs = append(g.Attrs, xml.Attr{Name: xml.Name{Local: name}, Value: value})
	g.l.Unlock()

	return g
}

// RemoveAttr removes all attributes of a given name of a Group
func (g Group) RemoveAttr(name string) Group {
	g.l.Lock()
	var attrs []xml.Attr
	for _, attr := range g.Attrs {
		if attr.Name.Local != name {
			attrs = append(attrs, attr)
		}
	}
	g.Attrs = attrs
	g.l.Unlock()

	return g
}

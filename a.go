package svg

import (
	"encoding/xml"
	"sync"
)

// A represents an A SVG element
// See: https://developer.mozilla.org/en-US/docs/Web/SVG/Element/a
type A struct {
	XMLName  xml.Name
	Href     string     `xml:"href,attr,omitempty"`
	Target   string     `xml:"target,attr,omitempty"`
	Type     string     `xml:"type,attr,omitempty"`
	Attrs    []xml.Attr `xml:",attr"`
	Children []interface{}
	lock     *sync.Mutex
}

// NewA constructs new A element
func NewA(href string, children ...interface{}) A {
	a := A{
		XMLName: xml.Name{Local: "a"},
		Href:    href,
		lock:    &sync.Mutex{},
	}

	a.Children = append(a.Children, children...)

	return a
}

// SetHrefLang sets the href attribute of an A tag
func (a A) SetTarget(target string) A {
	a.Target = target

	return a
}

// SetType sets the type attribute of an A tag
func (a A) SetType(t string) A {
	a.Type = t

	return a
}

// AddAttr adds a new attribute of an A tag
func (a A) AddAttr(name, value string) A {
	a.lock.Lock()
	a.Attrs = append(a.Attrs, xml.Attr{Name: xml.Name{Local: name}, Value: value})
	a.lock.Unlock()

	return a
}

// RemoveAttr removes all attributes of a given name of an A tag
func (a A) RemoveAttr(name string) A {
	a.lock.Lock()
	attrs := []xml.Attr{}
	for _, attr := range a.Attrs {
		if attr.Name.Local != name {
			attrs = append(attrs, attr)
		}
	}
	a.Attrs = attrs
	a.lock.Unlock()

	return a
}

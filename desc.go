package svg

import (
	"encoding/xml"
	"sync"
)

// Desc represents a Desc SVG element
// See: https://developer.mozilla.org/en-US/docs/Web/SVG/Element/desc
type Desc struct {
	XMLName  xml.Name
	Text     string     `xml:",innerxml"`
	Attrs    []xml.Attr `xml:",attr"`
	Children []interface{}
	lock     *sync.Mutex
}

// NewDesc constructs new Desc element
func NewDesc(text string, children ...interface{}) Desc {
	ts := Desc{
		XMLName: xml.Name{Local: "desc"},
		Text:    text,
		lock:    &sync.Mutex{},
	}

	ts.Children = append(ts.Children, children...)

	return ts
}

// AddAttr adds a new attribute of a Desc
func (d Desc) AddAttr(name, value string) Desc {
	d.lock.Lock()
	d.Attrs = append(d.Attrs, xml.Attr{Name: xml.Name{Local: name}, Value: value})
	d.lock.Unlock()

	return d
}

// RemoveAttr removes all attributes of a given name of a Desc
func (d Desc) RemoveAttr(name string) Desc {
	d.lock.Lock()
	var attrs []xml.Attr
	for _, attr := range d.Attrs {
		if attr.Name.Local != name {
			attrs = append(attrs, attr)
		}
	}
	d.Attrs = attrs
	d.lock.Unlock()

	return d
}

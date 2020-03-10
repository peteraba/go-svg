package svg

import (
	"encoding/xml"
	"sync"
)

type Element struct {
	XMLName  xml.Name
	Text     string     `xml:",innerxml"`
	Attrs    []xml.Attr `xml:",attr"`
	Children []interface{}
	lock     *sync.Mutex
}

func E(local, space, text string, attrs map[string]string, children ...interface{}) Element {
	element := Element{Text: text, lock: &sync.Mutex{}}

	if space == "" {
		element.XMLName = xml.Name{Local: local}
	} else {
		element.XMLName = xml.Name{Local: local, Space: space}
	}

	for n, v := range attrs {
		element.Attrs = append(element.Attrs, xml.Attr{Name: xml.Name{Local: n}, Value: v})
	}

	element.Children = append(element.Children, children...)

	return element
}

// AddAttr adds a new attribute of an Element
func (e Element) AddAttr(name, value string) Element {
	e.lock.Lock()
	e.Attrs = append(e.Attrs, xml.Attr{Name: xml.Name{Local: name}, Value: value})
	e.lock.Unlock()

	return e
}

// RemoveAttr removes all attributes of a given name of an Element
func (e Element) RemoveAttr(name string) Element {
	e.lock.Lock()
	attrs := []xml.Attr{}
	for _, attr := range e.Attrs {
		if attr.Name.Local != name {
			attrs = append(attrs, attr)
		}
	}
	e.Attrs = attrs
	e.lock.Unlock()

	return e
}

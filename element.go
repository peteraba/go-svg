package svg

import "encoding/xml"

type Element struct {
	XMLName  xml.Name
	Text     string     `xml:",innerxml,omitempty"`
	Attrs    []xml.Attr `xml:",attr"`
	Children []interface{}
}

func E(local, space, text string, attrs map[string]string, children []interface{}) Element {
	element := Element{Text: text}

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

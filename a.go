package svg

import (
	"encoding/xml"
)

// A represents an A SVG element
// See: https://developer.mozilla.org/en-US/docs/Web/SVG/Element/a
type A struct {
	XMLName  xml.Name
	Href     string `xml:"href,attr,omitempty"`
	HrefLang string `xml:"hreflang,attr,omitempty"`
	Target   string `xml:"target,attr,omitempty"`
	Type     string `xml:"type,attr,omitempty"`
	Children []interface{}
}

// NewA constructs new A element
func NewA(href string, children ...interface{}) A {
	a := A{
		XMLName: xml.Name{Local: "a"},
		Href:    href,
	}

	a.Children = append(a.Children, children...)

	return a
}

func (a A) SetHrefLang(hrefLang string) A {
	a.HrefLang = hrefLang

	return a
}

func (a A) SetTarget(target string) A {
	a.Target = target

	return a
}

func (a A) SetType(t string) A {
	a.Type = t

	return a
}

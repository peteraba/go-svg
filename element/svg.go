package element

import "encoding/xml"

type SVG struct {
	XMLName  xml.Name
	Width    float64 `xml:"width,attr,omitempty"`
	Height   float64 `xml:"height,attr,omitempty"`
	Version  string  `xml:"version,attr,omitempty"`
	Children []interface{}
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

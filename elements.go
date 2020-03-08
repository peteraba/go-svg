package svg

import (
	"encoding/xml"
)

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

type Line struct {
	XMLName     xml.Name
	X1          float64 `xml:"x1,attr,omitempty"`
	Y1          float64 `xml:"y1,attr,omitempty"`
	X2          float64 `xml:"x2,attr,omitempty"`
	Y2          float64 `xml:"y2,attr,omitempty"`
	StrokeWidth *uint8  `xml:"stroke-width,attr,omitempty"`
	Stroke      *Color  `xml:"stroke,attr,omitempty"`
	Children    []interface{}
}

func NewLine(x1, y1, x2, y2 float64, children ...interface{}) Line {
	l := Line{
		XMLName: xml.Name{Local: "line"},
		X1:      x1,
		Y1:      y1,
		X2:      x2,
		Y2:      y2,
	}

	l.Children = append(l.Children, children...)

	return l
}

func (l Line) SetStrokeWidth(strokeWidth uint8) Line {
	l.StrokeWidth = &strokeWidth

	return l
}

func (l Line) UnsetStrokeWidth() Line {
	l.StrokeWidth = nil

	return l
}

func (l Line) SetStroke(stroke Color) Line {
	l.Stroke = &stroke

	return l
}

func (l Line) UnsetStroke() Line {
	l.Stroke = nil

	return l
}

type Text struct {
	XMLName    xml.Name
	X          float64     `xml:"x,attr,omitempty"`
	Y          float64     `xml:"y,attr,omitempty"`
	TextAnchor *TextAnchor `xml:"text-anchor,attr,omitempty"`
	Fill       *Color      `xml:"stroke,attr,omitempty"`
	Children   []interface{}
}

func NewText(x, y float64, children ...interface{}) Text {
	t := Text{
		XMLName: xml.Name{Local: "text"},
		X:       x,
		Y:       y,
	}

	t.Children = append(t.Children, children...)

	return t
}

func (t Text) SetFill(fill Color) Text {
	t.Fill = &fill

	return t
}

func (t Text) UnsetFill() Text {
	t.Fill = nil

	return t
}

func (t Text) SetTextAnchor(ta TextAnchor) Text {
	t.TextAnchor = &ta

	return t
}

func (t Text) UnsetTextAnchor() Text {
	t.TextAnchor = nil

	return t
}

type TSpan struct {
	XMLName  xml.Name
	X        float64  `xml:"x,attr,omitempty"`
	Y        float64  `xml:"y,attr,omitempty"`
	DX       *float64 `xml:"dx,attr,omitempty"`
	DY       *float64 `xml:"dy,attr,omitempty"`
	Text     string   `xml:",innerxml"`
	Children []interface{}
}

func NewTSpan(text string, children ...interface{}) TSpan {
	ts := TSpan{
		XMLName: xml.Name{Local: "tspan"},
		Text:    text,
	}

	ts.Children = append(ts.Children, children...)

	return ts
}

func (ts TSpan) SetX(x float64) TSpan {
	ts.X = x

	return ts
}

func (ts TSpan) SetY(y float64) TSpan {
	ts.Y = y

	return ts
}

func (ts TSpan) SetDx(dx float64) TSpan {
	ts.DX = &dx

	return ts
}

func (ts TSpan) SetDy(dy float64) TSpan {
	ts.DY = &dy

	return ts
}

func (ts TSpan) UnsetDx() TSpan {
	ts.DX = nil

	return ts
}

func (ts TSpan) UnsetDy() TSpan {
	ts.DY = nil

	return ts
}


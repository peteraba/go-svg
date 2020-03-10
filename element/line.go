package element

import (
	"encoding/xml"

	"../attribute"
)

// Line represents a Line SVG element
// See: https://developer.mozilla.org/en-US/docs/Web/SVG/Element/line
type Line struct {
	XMLName       xml.Name
	X1            attribute.Length   `xml:"x1,attr,omitempty"`
	Y1            attribute.Length   `xml:"y1,attr,omitempty"`
	X2            attribute.Length   `xml:"x2,attr,omitempty"`
	Y2            attribute.Length   `xml:"y2,attr,omitempty"`
	StrokeWidth   *uint8             `xml:"stroke-width,attr,omitempty"`
	Stroke        *attribute.Color   `xml:"stroke,attr,omitempty"`
	StrokeOpacity *attribute.Opacity `xml:"stroke-opacity,attr,omitempty"`
	Fill          *attribute.Color   `xml:"fill,attr,omitempty"`
	FillOpacity   *attribute.Opacity `xml:"fill-opacity,attr,omitempty"`
	Opacity       float64            `xml:"opacity,attr,omitempty"`
	Children      []interface{}
}

// L constructs new Line element (shortcut)
func L(x1, y1, x2, y2 float64, children ...interface{}) Line {
	return NewLine(
		attribute.Length{Number: x1},
		attribute.Length{Number: y1},
		attribute.Length{Number: x2},
		attribute.Length{Number: y2},
		children...,
	)
}

// NewLine constructs new Line element
func NewLine(x1, y1, x2, y2 attribute.Length, children ...interface{}) Line {
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

// SetStrokeWidth sets the stroke width of a Line
func (l Line) SetStrokeWidth(strokeWidth uint8) Line {
	l.StrokeWidth = &strokeWidth

	return l
}

// UnsetStrokeWidth removes the previously set stroke width of a Line
func (l Line) UnsetStrokeWidth() Line {
	l.StrokeWidth = nil

	return l
}

// SetStroke sets the stroke color of a Line
func (l Line) SetStroke(stroke attribute.Color) Line {
	l.Stroke = &stroke

	return l
}

// UnsetStroke removes the previously set stroke color of a Line
func (l Line) UnsetStroke() Line {
	l.Stroke = nil

	return l
}

// SetStrokeOpacity sets the stroke opacity of a Line
func (l Line) SetStrokeOpacity(so attribute.Opacity) Line {
	l.StrokeOpacity = &so

	return l
}

// SetStrokeOpacity removes the stroke opacity of a Line
func (l Line) UnsetStrokeOpacity() Line {
	l.StrokeOpacity = nil

	return l
}

// SetStroke sets the fill color of a Line
func (l Line) SetFill(fill attribute.Color) Line {
	l.Fill = &fill

	return l
}

// UnsetStroke removes the previously set fill color of a Line
func (l Line) UnsetFill() Line {
	l.Fill = nil

	return l
}

// SetStrokeOpacity sets the fill opacity of a Line
func (l Line) SetFillOpacity(fo attribute.Opacity) Line {
	l.FillOpacity = &fo

	return l
}

// SetStrokeOpacity removes the stroke opacity of a Line
func (l Line) UnsetFillOpacity() Line {
	l.FillOpacity = nil

	return l
}

// SetOpacity sets the opacity of a Line
func (l Line) SetOpacity(o float64) Line {
	if o < 0 {
		l.Opacity = 0
	} else if o > 1 {
		l.Opacity = 1
	} else {
		l.Opacity = o
	}

	return l
}

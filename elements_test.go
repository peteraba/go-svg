package svg

import (
	"encoding/xml"
	"image/color"
	"reflect"
	"strings"
	"testing"
)

func TestNewSVG(t *testing.T) {
	type args struct {
		width    float64
		height   float64
		children []interface{}
	}
	tests := []struct {
		name string
		args args
		want SVG
	}{
		{
			"simple svg",
			args{width: 200, height: 50},
			SVG{XMLName: xml.Name{Space: "http://www.w3.org/2000/svg", Local: "svg"}, Width: 200, Height: 50, Version: "1.1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSVG(tt.args.width, tt.args.height, tt.args.children...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSVG() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewLine(t *testing.T) {
	type args struct {
		x1       float64
		y1       float64
		x2       float64
		y2       float64
		children []interface{}
	}
	tests := []struct {
		name string
		args args
		want Line
	}{
		{
			"simple line",
			args{x1: 1, y1: 2, x2: 3, y2: 4},
			Line{XMLName: xml.Name{Local: "line"}, X1: 1, Y1: 2, X2: 3, Y2: 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLine(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2, tt.args.children...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewText(t *testing.T) {
	type args struct {
		x        float64
		y        float64
		children []interface{}
	}
	tests := []struct {
		name string
		args args
		want Text
	}{
		{
			"simple text",
			args{},
			Text{XMLName: xml.Name{Local: "text"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewText(tt.args.x, tt.args.y, tt.args.children...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTSpan(t *testing.T) {
	type args struct {
		text     string
		children []interface{}
	}
	tests := []struct {
		name string
		args args
		want TSpan
	}{
		{
			"simple tspan",
			args{"Foo", nil},
			TSpan{XMLName: xml.Name{Local: "tspan"}, Text: "Foo"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTSpan(tt.args.text, tt.args.children...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTSpan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSVG_MarshalText(t *testing.T) {
	red := Color{color.RGBA{255, 0, 0, 255}}
	navy := Color{color.RGBA{0, 0, 128, 255}}

	tests := []struct {
		name string
		svg SVG
		wantLines []string
		wantErr bool
	}{
		{
			"empty svg",
			NewSVG(200, 100),
			[]string{`<svg xmlns="http://www.w3.org/2000/svg" width="200" height="100" version="1.1"></svg>`},
			false,
		},
		{
			"complex svg",
			NewSVG(
				200,
				100,
				NewLine(30, 30, 170, 30).SetStroke(red).SetStrokeWidth(2),
				NewLine(170, 30, 170, 70).SetStroke(red).SetStrokeWidth(2),
				NewLine(170, 70, 30, 70).SetStroke(navy).SetStrokeWidth(2),
				NewLine(30, 70, 30, 30).SetStroke(red).SetStrokeWidth(2),
				NewText(30, 40, NewTSpan("foo")).SetTextAnchor(Middle).SetFill(red),
				NewText(30, 40, NewTSpan("bar")).SetTextAnchor(Start).SetFill(navy),
			),
			[]string{
				`<svg xmlns="http://www.w3.org/2000/svg" width="200" height="100" version="1.1">`,
				`<line x1="30" y1="30" x2="170" y2="30" stroke-width="2" stroke="#ff0000"></line>`,
				`<line x1="170" y1="30" x2="170" y2="70" stroke-width="2" stroke="#ff0000"></line>`,
				`<line x1="170" y1="70" x2="30" y2="70" stroke-width="2" stroke="#000080"></line>`,
				`<line x1="30" y1="70" x2="30" y2="30" stroke-width="2" stroke="#ff0000"></line>`,
				`<text x="30" y="40" text-anchor="middle" stroke="#ff0000"><tspan>foo</tspan></text>`,
				`<text x="30" y="40" text-anchor="start" stroke="#000080"><tspan>bar</tspan></text>`,
				`</svg>`,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := strings.Join(tt.wantLines, "")
			gotBytes, err := xml.Marshal(tt.svg)
			if (err != nil) != tt.wantErr {
				t.Errorf("xml.Marshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := string(gotBytes)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("xml.Marshal() got = %v, want %v", got, want)
			}
		})
	}
}

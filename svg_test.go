package svg

import (
	"encoding/xml"
	"image/color"
	"reflect"
	"strings"
	"sync"
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
			SVG{XMLName: xml.Name{Space: "http://www.w3.org/2000/svg", Local: "svg"}, Width: 200, Height: 50, Version: "1.1", lock: &sync.Mutex{}},
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

func TestSVG_MarshalText(t *testing.T) {
	red := Color{color.RGBA{255, 0, 0, 255}}
	navy := Color{color.RGBA{0, 0, 128, 255}}

	tests := []struct {
		name      string
		svg       SVG
		wantLines []string
		wantErr   bool
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
				L(0, 30, 170, 30).SetStroke(red).SetStrokeWidth(2),
				L(170, 30, 170, 70).SetStroke(red).SetStrokeWidth(2),
				L(170, 70, 30, 70).SetStroke(navy).SetStrokeWidth(2),
				L(30, 70, 30, 30).SetStroke(red).SetStrokeWidth(2),
				T(0, 40, TS("foo")).SetTextAnchor(Middle).SetFill(red),
				T(30, 40, TS("bar").SX(30)).SetTextAnchor(Start).SetFill(navy),
			),
			[]string{
				`<svg xmlns="http://www.w3.org/2000/svg" width="200" height="100" version="1.1">`,
				`<line y1="30" x2="170" y2="30" stroke-width="2" stroke="#ff0000"></line>`,
				`<line x1="170" y1="30" x2="170" y2="70" stroke-width="2" stroke="#ff0000"></line>`,
				`<line x1="170" y1="70" x2="30" y2="70" stroke-width="2" stroke="#000080"></line>`,
				`<line x1="30" y1="70" x2="30" y2="30" stroke-width="2" stroke="#ff0000"></line>`,
				`<text y="40" text-anchor="middle" stroke="#ff0000"><tspan>foo</tspan></text>`,
				`<text x="30" y="40" text-anchor="start" stroke="#000080"><tspan x="30">bar</tspan></text>`,
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

func TestSVG_AddAttr(t *testing.T) {
	tests := []struct {
		name string
		c    SVG
		want string
	}{
		{
			"single attribute",
			NewSVG(200, 400).AddAttr("foo", "Foo"),
			`<svg xmlns="http://www.w3.org/2000/svg" width="200" height="400" version="1.1" foo="Foo"></svg>`,
		},
		{
			"multiple attributes",
			NewSVG(200, 400).AddAttr("foo", "Foo").AddAttr("bar", "Bar"),
			`<svg xmlns="http://www.w3.org/2000/svg" width="200" height="400" version="1.1" foo="Foo" bar="Bar"></svg>`,
		},
		{
			"single attribute repeated",
			NewSVG(200, 400).AddAttr("foo", "Foo").AddAttr("foo", "Bar"),
			`<svg xmlns="http://www.w3.org/2000/svg" width="200" height="400" version="1.1" foo="Foo" foo="Bar"></svg>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBytes, err := xml.Marshal(tt.c)
			if err != nil {
				t.Errorf("xml.Marshal() error = %v, wantErr %v", err, false)
				return
			}

			got := string(gotBytes)
			if got != tt.want {
				t.Errorf("xml.Marshal() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSVG_RemoveAttr(t *testing.T) {
	tests := []struct {
		name string
		c    SVG
		want string
	}{
		{
			"single attribute",
			NewSVG(200, 400).AddAttr("foo", "Foo").RemoveAttr("foo"),
			`<svg xmlns="http://www.w3.org/2000/svg" width="200" height="400" version="1.1"></svg>`,
		},
		{
			"multiple attributes",
			NewSVG(200, 400).AddAttr("foo", "Foo").AddAttr("bar", "Bar").RemoveAttr("foo"),
			`<svg xmlns="http://www.w3.org/2000/svg" width="200" height="400" version="1.1" bar="Bar"></svg>`,
		},
		{
			"single attribute repeated",
			NewSVG(200, 400).AddAttr("foo", "Foo").AddAttr("foo", "Bar").RemoveAttr("foo"),
			`<svg xmlns="http://www.w3.org/2000/svg" width="200" height="400" version="1.1"></svg>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBytes, err := xml.Marshal(tt.c)
			if err != nil {
				t.Errorf("xml.Marshal() error = %v, wantErr %v", err, false)
				return
			}

			got := string(gotBytes)
			if got != tt.want {
				t.Errorf("xml.Marshal() got = %v, want %v", got, tt.want)
			}
		})
	}
}

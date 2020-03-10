package svg

import (
	"encoding/xml"
	"reflect"
	"strings"
	"testing"
)

func TestNewText(t *testing.T) {
	type args struct {
		x        *Length
		y        *Length
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
		{
			"text width x and y",
			args{x: &Length{Number: 23.43}, y: &Length{Number: -43, Type: Em}},
			Text{XMLName: xml.Name{Local: "text"}, X: &Length{Number: 23.43}, Y: &Length{Number: -43, Type: Em}},
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

func TestT(t *testing.T) {
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
		{
			"text width x and y",
			args{x: 23.45, y: -34},
			Text{XMLName: xml.Name{Local: "text"}, X: &Length{23.45, ""}, Y: &Length{-34, ""}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := T(tt.args.x, tt.args.y, tt.args.children...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("T() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestText_UnsetFill(t1 *testing.T) {
	red, _ := ColorFromHexaString("#f00")

	type fields struct {
		XMLName    xml.Name
		X          *Length
		Y          *Length
		TextAnchor *TextAnchor
		Fill       *Color
		Children   []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   Text
	}{
		{
			"default",
			fields{Fill: &red},
			Text{},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := Text{
				XMLName:    tt.fields.XMLName,
				X:          tt.fields.X,
				Y:          tt.fields.Y,
				TextAnchor: tt.fields.TextAnchor,
				Fill:       tt.fields.Fill,
				Children:   tt.fields.Children,
			}
			if got := t.UnsetFill(); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("UnsetFill() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestText_UnsetTextAnchor(t1 *testing.T) {
	var middle = Middle
	type fields struct {
		XMLName    xml.Name
		X          *Length
		Y          *Length
		TextAnchor *TextAnchor
		Fill       *Color
		Children   []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   Text
	}{
		{
			"default",
			fields{TextAnchor: &middle},
			Text{},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := Text{
				XMLName:    tt.fields.XMLName,
				X:          tt.fields.X,
				Y:          tt.fields.Y,
				TextAnchor: tt.fields.TextAnchor,
				Fill:       tt.fields.Fill,
				Children:   tt.fields.Children,
			}
			if got := t.UnsetTextAnchor(); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("UnsetTextAnchor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestText_MarshalText(t *testing.T) {
	tests := []struct {
		name      string
		text      Text
		wantLines []string
		wantErr   bool
	}{
		{
			"simple text",
			T(0, 100),
			[]string{`<text y="100"></text>`},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := strings.Join(tt.wantLines, "")
			gotBytes, err := xml.Marshal(tt.text)
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

package svg

import (
	"encoding/xml"
	"reflect"
	"strings"
	"testing"
)

func TestNewLine(t *testing.T) {
	type args struct {
		x1       *Length
		y1       *Length
		x2       *Length
		y2       *Length
		children []interface{}
	}
	tests := []struct {
		name string
		args args
		want Line
	}{
		{
			"simple line",
			args{x1: &Length{Number: 1}, y1: &Length{Number: 2}, x2: &Length{Number: -4.2}, y2: &Length{Number: 4, Type: Em}},
			Line{XMLName: xml.Name{Local: "line"}, X1: &Length{Number: 1}, Y1: &Length{Number: 2}, X2: &Length{Number: -4.2}, Y2: &Length{Number: 4, Type: Em}},
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

func TestLine_MarshalText(t *testing.T) {
	tests := []struct {
		name      string
		line      Line
		wantLines []string
		wantErr   bool
	}{
		{
			"simple line",
			L(0, 100, 200, 150),
			[]string{`<line y1="100" x2="200" y2="150"></line>`},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := strings.Join(tt.wantLines, "")
			gotBytes, err := xml.Marshal(tt.line)
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

func TestLine_AddAttr(t *testing.T) {
	tests := []struct {
		name string
		c    Line
		want string
	}{
		{
			"single attribute",
			Line{}.AddAttr("foo", "Foo"),
			`<Line foo="Foo"></Line>`,
		},
		{
			"multiple attributes",
			Line{}.AddAttr("foo", "Foo").AddAttr("bar", "Bar"),
			`<Line foo="Foo" bar="Bar"></Line>`,
		},
		{
			"single attribute repeated",
			Line{}.AddAttr("foo", "Foo").AddAttr("foo", "Bar"),
			`<Line foo="Foo" foo="Bar"></Line>`,
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

func TestLine_RemoveAttr(t *testing.T) {
	tests := []struct {
		name string
		c    Line
		want string
	}{
		{
			"single attribute",
			Line{}.AddAttr("foo", "Foo").RemoveAttr("foo"),
			`<Line></Line>`,
		},
		{
			"multiple attributes",
			Line{}.AddAttr("foo", "Foo").AddAttr("bar", "Bar").RemoveAttr("foo"),
			`<Line bar="Bar"></Line>`,
		},
		{
			"single attribute repeated",
			Line{}.AddAttr("foo", "Foo").AddAttr("foo", "Bar").RemoveAttr("foo"),
			`<Line></Line>`,
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

package svg

import (
	"encoding/xml"
	"reflect"
	"strings"
	"sync"
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
			Line{XMLName: xml.Name{Local: "line"}, X1: &Length{Number: 1}, Y1: &Length{Number: 2}, X2: &Length{Number: -4.2}, Y2: &Length{Number: 4, Type: Em}, lock: &sync.Mutex{}},
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
			L(2, 4, 6, 8).AddAttr("foo", "Foo"),
			`<line x1="2" y1="4" x2="6" y2="8" foo="Foo"></line>`,
		},
		{
			"multiple attributes",
			L(2, 4, 6, 8).AddAttr("foo", "Foo").AddAttr("bar", "Bar"),
			`<line x1="2" y1="4" x2="6" y2="8" foo="Foo" bar="Bar"></line>`,
		},
		{
			"single attribute repeated",
			L(2, 4, 6, 8).AddAttr("foo", "Foo").AddAttr("foo", "Bar"),
			`<line x1="2" y1="4" x2="6" y2="8" foo="Foo" foo="Bar"></line>`,
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
			L(2, 4, 6, 8).AddAttr("foo", "Foo").RemoveAttr("foo"),
			`<line x1="2" y1="4" x2="6" y2="8"></line>`,
		},
		{
			"multiple attributes",
			L(2, 4, 6, 8).AddAttr("foo", "Foo").AddAttr("bar", "Bar").RemoveAttr("foo"),
			`<line x1="2" y1="4" x2="6" y2="8" bar="Bar"></line>`,
		},
		{
			"single attribute repeated",
			L(2, 4, 6, 8).AddAttr("foo", "Foo").AddAttr("foo", "Bar").RemoveAttr("foo"),
			`<line x1="2" y1="4" x2="6" y2="8"></line>`,
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

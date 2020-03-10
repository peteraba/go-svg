package svg

import (
	"encoding/xml"
	"reflect"
	"strings"
	"testing"
)

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

func TestTSpan_MarshalText(t *testing.T) {
	tests := []struct {
		name      string
		tspan     TSpan
		wantLines []string
		wantErr   bool
	}{
		{
			"simple tspan",
			TS("foo"),
			[]string{`<tspan>foo</tspan>`},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := strings.Join(tt.wantLines, "")
			gotBytes, err := xml.Marshal(tt.tspan)
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

func TestTSpan_AddAttr(t *testing.T) {
	tests := []struct {
		name string
		c    TSpan
		want string
	}{
		{
			"single attribute",
			TSpan{}.AddAttr("foo", "Foo"),
			`<TSpan foo="Foo"></TSpan>`,
		},
		{
			"multiple attributes",
			TSpan{}.AddAttr("foo", "Foo").AddAttr("bar", "Bar"),
			`<TSpan foo="Foo" bar="Bar"></TSpan>`,
		},
		{
			"single attribute repeated",
			TSpan{}.AddAttr("foo", "Foo").AddAttr("foo", "Bar"),
			`<TSpan foo="Foo" foo="Bar"></TSpan>`,
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

func TestTSpan_RemoveAttr(t *testing.T) {
	tests := []struct {
		name string
		c    TSpan
		want string
	}{
		{
			"single attribute",
			TSpan{}.AddAttr("foo", "Foo").RemoveAttr("foo"),
			`<TSpan></TSpan>`,
		},
		{
			"multiple attributes",
			TSpan{}.AddAttr("foo", "Foo").AddAttr("bar", "Bar").RemoveAttr("foo"),
			`<TSpan bar="Bar"></TSpan>`,
		},
		{
			"single attribute repeated",
			TSpan{}.AddAttr("foo", "Foo").AddAttr("foo", "Bar").RemoveAttr("foo"),
			`<TSpan></TSpan>`,
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

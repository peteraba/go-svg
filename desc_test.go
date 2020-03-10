package svg

import (
	"encoding/xml"
	"reflect"
	"strings"
	"sync"
	"testing"
)

func TestNewDesc(t *testing.T) {
	type args struct {
		text     string
		children []interface{}
	}
	tests := []struct {
		name string
		args args
		want Desc
	}{
		{
			"simple desc",
			args{"Foo", nil},
			Desc{XMLName: xml.Name{Local: "desc"}, Text: "Foo", lock: &sync.Mutex{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDesc(tt.args.text, tt.args.children...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDesc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDesc_MarshalText(t *testing.T) {
	tests := []struct {
		name      string
		tspan     Desc
		wantLines []string
		wantErr   bool
	}{
		{
			"simple desc",
			NewDesc("foo"),
			[]string{`<desc>foo</desc>`},
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

func TestDesc_AddAttr(t *testing.T) {
	tests := []struct {
		name string
		c    Desc
		want string
	}{
		{
			"single attribute",
			NewDesc("baz").AddAttr("foo", "Foo"),
			`<desc foo="Foo">baz</desc>`,
		},
		{
			"multiple attributes",
			NewDesc("baz").AddAttr("foo", "Foo").AddAttr("bar", "Bar"),
			`<desc foo="Foo" bar="Bar">baz</desc>`,
		},
		{
			"single attribute repeated",
			NewDesc("baz").AddAttr("foo", "Foo").AddAttr("foo", "Bar"),
			`<desc foo="Foo" foo="Bar">baz</desc>`,
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

func TestDesc_RemoveAttr(t *testing.T) {
	tests := []struct {
		name string
		c    Desc
		want string
	}{
		{
			"single attribute",
			NewDesc("baz").AddAttr("foo", "Foo").RemoveAttr("foo"),
			`<desc>baz</desc>`,
		},
		{
			"multiple attributes",
			NewDesc("baz").AddAttr("foo", "Foo").AddAttr("bar", "Bar").RemoveAttr("foo"),
			`<desc bar="Bar">baz</desc>`,
		},
		{
			"single attribute repeated",
			NewDesc("baz").AddAttr("foo", "Foo").AddAttr("foo", "Bar").RemoveAttr("foo"),
			`<desc>baz</desc>`,
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

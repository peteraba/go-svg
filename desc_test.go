package svg

import (
	"encoding/xml"
	"reflect"
	"strings"
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
			Desc{XMLName: xml.Name{Local: "desc"}, Text: "Foo"},
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
			Desc{}.AddAttr("foo", "Foo"),
			`<Desc foo="Foo"></Desc>`,
		},
		{
			"multiple attributes",
			Desc{}.AddAttr("foo", "Foo").AddAttr("bar", "Bar"),
			`<Desc foo="Foo" bar="Bar"></Desc>`,
		},
		{
			"single attribute repeated",
			Desc{}.AddAttr("foo", "Foo").AddAttr("foo", "Bar"),
			`<Desc foo="Foo" foo="Bar"></Desc>`,
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
			Desc{}.AddAttr("foo", "Foo").RemoveAttr("foo"),
			`<Desc></Desc>`,
		},
		{
			"multiple attributes",
			Desc{}.AddAttr("foo", "Foo").AddAttr("bar", "Bar").RemoveAttr("foo"),
			`<Desc bar="Bar"></Desc>`,
		},
		{
			"single attribute repeated",
			Desc{}.AddAttr("foo", "Foo").AddAttr("foo", "Bar").RemoveAttr("foo"),
			`<Desc></Desc>`,
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

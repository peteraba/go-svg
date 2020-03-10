package svg

import (
	"encoding/xml"
	"reflect"
	"strings"
	"sync"
	"testing"
)

func TestNewGroup(t *testing.T) {
	type args struct {
		children []interface{}
	}
	tests := []struct {
		name string
		args args
		want Group
	}{
		{
			"simple group",
			args{},
			Group{XMLName: xml.Name{Local: "g"}, lock: &sync.Mutex{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGroup(tt.args.children...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroup_MarshalText(t *testing.T) {
	tests := []struct {
		name      string
		tspan     Group
		wantLines []string
		wantErr   bool
	}{
		{
			"simple group",
			NewGroup(),
			[]string{`<g></g>`},
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

func TestGroup_AddAttr(t *testing.T) {
	tests := []struct {
		name string
		c    Group
		want string
	}{
		{
			"single attribute",
			NewGroup().AddAttr("foo", "Foo"),
			`<g foo="Foo"></g>`,
		},
		{
			"multiple attributes",
			NewGroup().AddAttr("foo", "Foo").AddAttr("bar", "Bar"),
			`<g foo="Foo" bar="Bar"></g>`,
		},
		{
			"single attribute repeated",
			NewGroup().AddAttr("foo", "Foo").AddAttr("foo", "Bar"),
			`<g foo="Foo" foo="Bar"></g>`,
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

func TestGroup_RemoveAttr(t *testing.T) {
	tests := []struct {
		name string
		c    Group
		want string
	}{
		{
			"single attribute",
			NewGroup().AddAttr("foo", "Foo").RemoveAttr("foo"),
			`<g></g>`,
		},
		{
			"multiple attributes",
			NewGroup().AddAttr("foo", "Foo").AddAttr("bar", "Bar").RemoveAttr("foo"),
			`<g bar="Bar"></g>`,
		},
		{
			"single attribute repeated",
			NewGroup().AddAttr("foo", "Foo").AddAttr("foo", "Bar").RemoveAttr("foo"),
			`<g></g>`,
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

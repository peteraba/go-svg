package svg

import (
	"encoding/xml"
	"reflect"
	"strings"
	"testing"
)

func TestNewA(t *testing.T) {
	type args struct {
		href     string
		children []interface{}
	}
	tests := []struct {
		name string
		args args
		want A
	}{
		{
			"simple anchor",
			args{href: "http://foo.com/"},
			A{XMLName: xml.Name{Local: "a"}, Href: "http://foo.com/"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewA(tt.args.href, tt.args.children...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestA_MarshalText(t *testing.T) {
	tests := []struct {
		name      string
		a         A
		wantLines []string
		wantErr   bool
	}{
		{
			"simple anchor",
			NewA("http://foo.com/"),
			[]string{`<a href="http://foo.com/"></a>`},
			false,
		},
		{
			"complex anchor",

			NewA("http://foo.com/").AddAttr("foo", "Foo").AddAttr("bar", "Bar"),
			[]string{`<a href="http://foo.com/" foo="Foo" bar="Bar"></a>`},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := strings.Join(tt.wantLines, "")
			gotBytes, err := xml.Marshal(tt.a)
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

func TestA_AddAttr(t *testing.T) {
	tests := []struct {
		name string
		a    A
		want string
	}{
		{
			"single attribute",
			A{}.AddAttr("foo", "Foo"),
			`<A foo="Foo"></A>`,
		},
		{
			"multiple attributes",
			A{}.AddAttr("foo", "Foo").AddAttr("bar", "Bar"),
			`<A foo="Foo" bar="Bar"></A>`,
		},
		{
			"single attribute repeated",
			A{}.AddAttr("foo", "Foo").AddAttr("foo", "Bar"),
			`<A foo="Foo" foo="Bar"></A>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBytes, err := xml.Marshal(tt.a)
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

func TestA_RemoveAttr(t *testing.T) {
	tests := []struct {
		name string
		a    A
		want string
	}{
		{
			"single attribute",
			A{}.AddAttr("foo", "Foo").RemoveAttr("foo"),
			`<A></A>`,
		},
		{
			"multiple attributes",
			A{}.AddAttr("foo", "Foo").AddAttr("bar", "Bar").RemoveAttr("foo"),
			`<A bar="Bar"></A>`,
		},
		{
			"single attribute repeated",
			A{}.AddAttr("foo", "Foo").AddAttr("foo", "Bar").RemoveAttr("foo"),
			`<A></A>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBytes, err := xml.Marshal(tt.a)
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

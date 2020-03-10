package element

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

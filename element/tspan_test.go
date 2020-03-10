package element

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
			"one line",
			TS("foo"),
			[]string{`<tspan x="0" y="0">foo</tspan>`},
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

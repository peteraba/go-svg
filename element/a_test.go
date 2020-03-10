package element

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
			"simple line",
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
			"simple a",
			NewA("http://foo.com/"),
			[]string{`<a href="http://foo.com/"></a>`},
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

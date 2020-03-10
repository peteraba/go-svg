package svg

import (
	"encoding/xml"
	"reflect"
	"strings"
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
			Group{XMLName: xml.Name{Local: "g"}},
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

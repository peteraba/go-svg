package element

import (
	"encoding/xml"
	"reflect"
	"strings"
	"testing"

	"../attribute"
)

func TestNewLine(t *testing.T) {
	type args struct {
		x1       attribute.Length
		y1       attribute.Length
		x2       attribute.Length
		y2       attribute.Length
		children []interface{}
	}
	tests := []struct {
		name string
		args args
		want Line
	}{
		{
			"simple line",
			args{x1: attribute.Length{Number: 1}, y1: attribute.Length{Number: 2}, x2: attribute.Length{Number: -4.2}, y2: attribute.Length{Number: 4, Type: attribute.Em}},
			Line{XMLName: xml.Name{Local: "line"}, X1: attribute.Length{Number: 1}, Y1: attribute.Length{Number: 2}, X2: attribute.Length{Number: -4.2}, Y2: attribute.Length{Number: 4, Type: attribute.Em}},
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
			"one line",
			L(0, 100, 200, 150),
			[]string{`<line x1="0" y1="100" x2="200" y2="150"></line>`},
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

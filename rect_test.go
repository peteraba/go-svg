package svg

import (
	"encoding/xml"
	"reflect"
	"strings"
	"testing"
)

func TestR(t *testing.T) {
	type args struct {
		x        float64
		y        float64
		width    float64
		height   float64
		children []interface{}
	}
	tests := []struct {
		name string
		args args
		want Rect
	}{
		{
			"simple rect",
			args{x: 1, y: 2, width: 4.2, height: 3.1},
			Rect{XMLName: xml.Name{Local: "rect"}, X: Length{Number: 1}, Y: Length{Number: 2}, Width: Length{Number: 4.2}, Height: Length{Number: 3.1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := R(tt.args.x, tt.args.y, tt.args.width, tt.args.height, tt.args.children...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("E() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRect(t *testing.T) {
	type args struct {
		x        Length
		y        Length
		width    Length
		height   Length
		rx       Length
		ry       Length
		children []interface{}
	}
	tests := []struct {
		name string
		args args
		want Rect
	}{
		{
			"simple rect",
			args{x: Length{Number: 1}, y: Length{Number: 2}, width: Length{Number: 20}, height: Length{Number: 10}, rx: Length{Number: 4.2}, ry: Length{Number: 3.1}},
			Rect{XMLName: xml.Name{Local: "rect"}, X: Length{Number: 1}, Y: Length{Number: 2}, Width: Length{Number: 20}, Height: Length{Number: 10}, RX: Length{Number: 4.2}, RY: Length{Number: 3.1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRect(tt.args.x, tt.args.y, tt.args.width, tt.args.height, tt.args.rx, tt.args.ry, tt.args.children...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRect_MarshalText(t *testing.T) {
	tests := []struct {
		name      string
		line      Rect
		wantLines []string
		wantErr   bool
	}{
		{
			"simple rect",
			R(0, 100, 50, 20),
			[]string{`<rect x="0" y="100" width="50" height="20" rx="0" ry="0"></rect>`},
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
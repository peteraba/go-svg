package svg

import (
	"encoding/xml"
	"reflect"
	"strings"
	"testing"
)

func TestE(t *testing.T) {
	type args struct {
		cx       float64
		cy       float64
		rx       float64
		ry       float64
		children []interface{}
	}
	tests := []struct {
		name string
		args args
		want Ellipse
	}{
		{
			"simple ellipse",
			args{cx: 1, cy: 2, rx: 4.2, ry: 3.1},
			Ellipse{XMLName: xml.Name{Local: "ellipse"}, CX: &Length{Number: 1}, CY: &Length{Number: 2}, RX: &Length{Number: 4.2}, RY: &Length{Number: 3.1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := E(tt.args.cx, tt.args.cy, tt.args.rx, tt.args.ry, tt.args.children...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("E() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewEllipse(t *testing.T) {
	type args struct {
		cx       *Length
		cy       *Length
		rx       *Length
		ry       *Length
		children []interface{}
	}
	tests := []struct {
		name string
		args args
		want Ellipse
	}{
		{
			"simple ellipse",
			args{cx: &Length{Number: 1}, cy: &Length{Number: 2}, rx: &Length{Number: 4.2}, ry: &Length{Number: 3.1}},
			Ellipse{XMLName: xml.Name{Local: "ellipse"}, CX: &Length{Number: 1}, CY: &Length{Number: 2}, RX: &Length{Number: 4.2}, RY: &Length{Number: 3.1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEllipse(tt.args.cx, tt.args.cy, tt.args.rx, tt.args.ry, tt.args.children...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEllipse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEllipse_MarshalText(t *testing.T) {
	tests := []struct {
		name      string
		line      Ellipse
		wantLines []string
		wantErr   bool
	}{
		{
			"simple ellipse",
			E(0, 100, 50, 20),
			[]string{`<ellipse cy="100" rx="50" ry="20"></ellipse>`},
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

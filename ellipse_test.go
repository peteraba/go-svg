package svg

import (
	"encoding/xml"
	"reflect"
	"strings"
	"sync"
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
			Ellipse{XMLName: xml.Name{Local: "ellipse"}, CX: &Length{Number: 1}, CY: &Length{Number: 2}, RX: &Length{Number: 4.2}, RY: &Length{Number: 3.1}, lock: &sync.Mutex{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := El(tt.args.cx, tt.args.cy, tt.args.rx, tt.args.ry, tt.args.children...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("El() = %v, want %v", got, tt.want)
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
			Ellipse{XMLName: xml.Name{Local: "ellipse"}, CX: &Length{Number: 1}, CY: &Length{Number: 2}, RX: &Length{Number: 4.2}, RY: &Length{Number: 3.1}, lock: &sync.Mutex{}},
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
			El(0, 100, 50, 20),
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

func TestEllipse_AddAttr(t *testing.T) {
	tests := []struct {
		name string
		c    Ellipse
		want string
	}{
		{
			"single attribute",
			El(2, 4, 6, 8).AddAttr("foo", "Foo"),
			`<ellipse cx="2" cy="4" rx="6" ry="8" foo="Foo"></ellipse>`,
		},
		{
			"multiple attributes",
			El(2, 4, 6, 8).AddAttr("foo", "Foo").AddAttr("bar", "Bar"),
			`<ellipse cx="2" cy="4" rx="6" ry="8" foo="Foo" bar="Bar"></ellipse>`,
		},
		{
			"single attribute repeated",
			El(2, 4, 6, 8).AddAttr("foo", "Foo").AddAttr("foo", "Bar"),
			`<ellipse cx="2" cy="4" rx="6" ry="8" foo="Foo" foo="Bar"></ellipse>`,
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

func TestEllipse_RemoveAttr(t *testing.T) {
	tests := []struct {
		name string
		c    Ellipse
		want string
	}{
		{
			"single attribute",
			El(2, 4, 6, 8).AddAttr("foo", "Foo").RemoveAttr("foo"),
			`<ellipse cx="2" cy="4" rx="6" ry="8"></ellipse>`,
		},
		{
			"multiple attributes",
			El(2, 4, 6, 8).AddAttr("foo", "Foo").AddAttr("bar", "Bar").RemoveAttr("foo"),
			`<ellipse cx="2" cy="4" rx="6" ry="8" bar="Bar"></ellipse>`,
		},
		{
			"single attribute repeated",
			El(2, 4, 6, 8).AddAttr("foo", "Foo").AddAttr("foo", "Bar").RemoveAttr("foo"),
			`<ellipse cx="2" cy="4" rx="6" ry="8"></ellipse>`,
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

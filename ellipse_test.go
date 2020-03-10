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
			Ellipse{}.AddAttr("foo", "Foo"),
			`<Ellipse foo="Foo"></Ellipse>`,
		},
		{
			"multiple attributes",
			Ellipse{}.AddAttr("foo", "Foo").AddAttr("bar", "Bar"),
			`<Ellipse foo="Foo" bar="Bar"></Ellipse>`,
		},
		{
			"single attribute repeated",
			Ellipse{}.AddAttr("foo", "Foo").AddAttr("foo", "Bar"),
			`<Ellipse foo="Foo" foo="Bar"></Ellipse>`,
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
			Ellipse{}.AddAttr("foo", "Foo").RemoveAttr("foo"),
			`<Ellipse></Ellipse>`,
		},
		{
			"multiple attributes",
			Ellipse{}.AddAttr("foo", "Foo").AddAttr("bar", "Bar").RemoveAttr("foo"),
			`<Ellipse bar="Bar"></Ellipse>`,
		},
		{
			"single attribute repeated",
			Ellipse{}.AddAttr("foo", "Foo").AddAttr("foo", "Bar").RemoveAttr("foo"),
			`<Ellipse></Ellipse>`,
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

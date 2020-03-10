package svg

import (
	"encoding/xml"
	"reflect"
	"strings"
	"testing"
)

func TestC(t *testing.T) {
	type args struct {
		cx       float64
		cy       float64
		r        float64
		children []interface{}
	}
	tests := []struct {
		name string
		args args
		want Circle
	}{
		{
			"simple circle",
			args{cx: 1, cy: 2, r: 4.2},
			Circle{XMLName: xml.Name{Local: "circle"}, CX: &Length{Number: 1}, CY: &Length{Number: 2}, R: &Length{Number: 4.2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := C(tt.args.cx, tt.args.cy, tt.args.r, tt.args.children...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("C() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCircle(t *testing.T) {
	type args struct {
		cx       *Length
		cy       *Length
		r        *Length
		children []interface{}
	}
	tests := []struct {
		name string
		args args
		want Circle
	}{
		{
			"simple circle",
			args{cx: &Length{Number: 1}, cy: &Length{Number: 2}, r: &Length{Number: 4.2}},
			Circle{XMLName: xml.Name{Local: "circle"}, CX: &Length{Number: 1}, CY: &Length{Number: 2}, R: &Length{Number: 4.2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCircle(tt.args.cx, tt.args.cy, tt.args.r, tt.args.children...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCircle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCircle_MarshalText(t *testing.T) {
	tests := []struct {
		name      string
		line      Circle
		wantLines []string
		wantErr   bool
	}{
		{
			"simple circle",
			C(0, 100, 50),
			[]string{`<circle cy="100" r="50"></circle>`},
			false,
		},
		{
			"complex circle",
			C(0, 100, 50),
			[]string{`<circle cy="100" r="50"></circle>`},
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

func TestCircle_AddAttr(t *testing.T) {
	tests := []struct {
		name string
		c    Circle
		want string
	}{
		{
			"single attribute",
			Circle{}.AddAttr("foo", "Foo"),
			`<Circle foo="Foo"></Circle>`,
		},
		{
			"multiple attributes",
			Circle{}.AddAttr("foo", "Foo").AddAttr("bar", "Bar"),
			`<Circle foo="Foo" bar="Bar"></Circle>`,
		},
		{
			"single attribute repeated",
			Circle{}.AddAttr("foo", "Foo").AddAttr("foo", "Bar"),
			`<Circle foo="Foo" foo="Bar"></Circle>`,
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

func TestCircle_RemoveAttr(t *testing.T) {
	tests := []struct {
		name string
		c    Circle
		want string
	}{
		{
			"single attribute",
			Circle{}.AddAttr("href", "Foo").RemoveAttr("href"),
			`<Circle></Circle>`,
		},
		{
			"multiple attributes",
			Circle{}.AddAttr("href", "Foo").AddAttr("bar", "Bar").RemoveAttr("href"),
			`<Circle bar="Bar"></Circle>`,
		},
		{
			"single attribute repeated",
			Circle{}.AddAttr("href", "Foo").AddAttr("href", "Bar").RemoveAttr("href"),
			`<Circle></Circle>`,
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

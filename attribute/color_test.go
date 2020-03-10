package attribute

import (
	"image/color"
	"reflect"
	"testing"
)

func TestColor_MarshalText(t *testing.T) {
	type fields struct {
		RGBA color.RGBA
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			"red",
			fields{color.RGBA{255, 0, 0, 0}},
			[]byte("#ff0000"),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Color{
				RGBA: tt.fields.RGBA,
			}
			got, err := c.MarshalText()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalText() got = %v, want %v", got, tt.want)
			}
		})
	}

	t.Run("nil receiver handled", func(t *testing.T) {
		var c *Color

		got, err := c.MarshalText()
		if err != nil {
			t.Errorf("MarshalText() error = %v, wantErr %v", err, false)
			return
		}
		if !reflect.DeepEqual(got, []byte{}) {
			t.Errorf("MarshalText() got = %v, want %v", got, []byte{})
		}
	})
}

func TestColor_String(t *testing.T) {
	type fields struct {
		RGBA color.RGBA
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"red",
			fields{color.RGBA{255, 0, 0, 255}},
			"#ff0000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Color{
				RGBA: tt.fields.RGBA,
			}
			if got := c.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestColor_UnmarshalText(t *testing.T) {
	type fields struct {
		RGBA color.RGBA
	}
	type args struct {
		text []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Color
		wantErr bool
	}{
		{
			"red long",
			fields{},
			args{[]byte("#ff0000")},
			Color{color.RGBA{255, 0, 0, 255}},
			false,
		},
		{
			"red short",
			fields{},
			args{[]byte("#f00")},
			Color{color.RGBA{255, 0, 0, 255}},
			false,
		},
		{
			"red word",
			fields{},
			args{[]byte("red")},
			Color{color.RGBA{255, 0, 0, 255}},
			false,
		},
		{
			"red word",
			fields{},
			args{[]byte("ezüst")},
			Color{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Color{
				RGBA: tt.fields.RGBA,
			}
			if err := c.UnmarshalText(tt.args.text); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalText() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(*c, tt.want) {
				t.Errorf("MarshalText() got = %v, want %v", c, tt.want)
			}
		})
	}
}

func TestColorFromHexaString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    Color
		wantErr bool
	}{
		{
			"short color code",
			args{s: "#25f"},
			Color{color.RGBA{34, 85, 255, 255}},
			false,
		},
		{
			"long color code",
			args{s: "#2255ff"},
			Color{color.RGBA{34, 85, 255, 255}},
			false,
		},
		{
			"invalid color code length",
			args{s: "#2255ff8"},
			Color{},
			true,
		},
		{
			"invalid color code char",
			args{s: "#25z"},
			Color{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ColorFromHexaString(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseHexaColor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseHexaColor() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_charsToUint8(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    [3]uint8
		wantErr bool
	}{
		{
			"short color code",
			args{s: "25f"},
			[3]uint8{34, 85, 255},
			false,
		},
		{
			"long color code",
			args{s: "2255ff"},
			[3]uint8{34, 85, 255},
			false,
		},
		{
			"invalid color code length",
			args{s: "2255ff8"},
			[3]uint8{},
			true,
		},
		{
			"invalid color code char",
			args{s: "25z"},
			[3]uint8{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := charsToUint8(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("charsToUint8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("charsToUint8() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_twoDigitHexa(t *testing.T) {
	type args struct {
		i uint8
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"single digit",
			args{i: 0xe},
			"0e",
		},
		{
			"valid double digit",
			args{i: 0xef},
			"ef",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoDigitHexa(tt.args.i); got != tt.want {
				t.Errorf("twoDigitHexa() = %v, want %v", got, tt.want)
			}
		})
	}
}

package svg

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
		want Color
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

func TestTextAnchor_MarshalText(t *testing.T) {
	tests := []struct {
		name    string
		t       TextAnchor
		want    []byte
		wantErr bool
	}{
		{
			"middle",
			Middle,
			[]byte("middle"),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.t.MarshalText()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalText() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTextAnchor_UnmarshalText(t *testing.T) {
	type args struct {
		text []byte
	}
	tests := []struct {
		name    string
		t       TextAnchor
		args    args
		want       TextAnchor
		wantErr bool
	}{
		{
			"start to middle",
			Start,
			args{[]byte("middle")},
			Middle,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.t.UnmarshalText(tt.args.text); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalText() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.t, tt.want) {
				t.Errorf("MarshalText() got = %v, want %v", tt.t, tt.want)
			}
		})
	}
}


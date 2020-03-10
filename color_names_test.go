package svg

import (
	"image/color"
	"reflect"
	"testing"
)

func TestNewColorName(t *testing.T) {
	type args struct {
		cn string
	}
	tests := []struct {
		name    string
		args    args
		want    ColorName
		wantErr bool
	}{
		{
			"silver",
			args{cn: "silver"},
			ColorName("silver"),
			false,
		},
		{
			"ezüst",
			args{cn: "ezüst"},
			"",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewColorName(tt.args.cn)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewColorName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewColorName() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestColorName_ToColor(t *testing.T) {
	tests := []struct {
		name string
		cn   ColorName
		want Color
	}{
		{
			"silver",
			ColorName("silver"),
			Color{color.RGBA{192, 192, 192, 255}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.cn.ToColor()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToColor() got = %v, want %v", got, tt.want)
			}
		})
	}

	t.Run("panic on invalid color name", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()

		cn := ColorName("ezüst")

		cn.ToColor()
	})
}

func TestColorName_ToHexa(t *testing.T) {
	tests := []struct {
		name string
		cn   ColorName
		want string
	}{
		{
			"silver",
			ColorName("silver"),
			"#c0c0c0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.cn.ToHexa()
			if got != tt.want {
				t.Errorf("ToHexa() got = %v, want %v", got, tt.want)
			}
		})
	}

	t.Run("panic on invalid color name", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()

		cn := ColorName("ezüst")

		cn.ToHexa()
	})
}

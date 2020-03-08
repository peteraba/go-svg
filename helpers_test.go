package svg

import (
	"image/color"
	"reflect"
	"testing"
)

func TestParseHexaColor(t *testing.T) {
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
			Color{ },
			true,
		},
		{
			"invalid color code char",
			args{s: "#25z"},
			Color{ },
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseHexaColor(tt.args.s)
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

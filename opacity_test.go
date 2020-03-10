package svg

import (
	"reflect"
	"testing"
)

func TestOpacityType_MarshalText(t *testing.T) {
	tests := []struct {
		name    string
		lt      OpacityType
		want    []byte
		wantErr bool
	}{
		{
			"empty",
			"",
			[]byte{},
			false,
		},
		{
			"%",
			"%",
			[]byte("%"),
			false,
		},
		{
			"invalid is treated like empty",
			"invalid",
			[]byte{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.lt.MarshalText()
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

func TestOpacityType_UnmarshalText(t *testing.T) {
	type args struct {
		text []byte
	}
	tests := []struct {
		name    string
		lt      OpacityType
		lt2     OpacityType
		args    args
		wantErr bool
	}{
		{
			"% to empty",
			"%",
			"",
			args{[]byte{}},
			false,
		},
		{
			"empty to %",
			"",
			"%",
			args{[]byte("%")},
			false,
		},
		{
			"invalid is treated like empty",
			"%",
			"",
			args{[]byte("invalid")},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.lt.UnmarshalText(tt.args.text); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalText() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.lt, tt.lt2) {
				t.Errorf("MarshalText() got = %v, want %v", tt.lt, tt.lt2)
			}
		})
	}
}

func TestO(t *testing.T) {
	type args struct {
		n       float64
		lengths []OpacityType
	}
	tests := []struct {
		name string
		args args
		want Opacity
	}{
		{
			"no length type",
			args{26.4, nil},
			Opacity{26.4, ""},
		},
		{
			"first length type is used",
			args{26.4, []OpacityType{OPercent}},
			Opacity{26.4, "%"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := O(tt.args.n, tt.args.lengths...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("L() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOpacity_MarshalText(t *testing.T) {
	type fields struct {
		Number float64
		Type   OpacityType
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			"number only",
			fields{65.23, ""},
			[]byte("65.23"),
			false,
		},
		{
			"percent",
			fields{-17.2321, "%"},
			[]byte("-17.2321%"),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Opacity{
				Number: tt.fields.Number,
				Type:   tt.fields.Type,
			}
			got, err := l.MarshalText()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalText() got = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}

func TestOpacity_UnmarshalText(t *testing.T) {
	type args struct {
		text []byte
	}
	tests := []struct {
		name    string
		length  Opacity
		args    args
		want    Opacity
		wantErr bool
	}{
		{
			"percent to number",
			Opacity{5426432.9382, "%"},
			args{[]byte("65.23")},
			Opacity{65.23, ""},
			false,
		},
		{
			"number to percent",
			Opacity{6845458, ""},
			args{[]byte("3.33%")},
			Opacity{3.33, "%"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &tt.length
			if err := l.UnmarshalText(tt.args.text); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalText() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

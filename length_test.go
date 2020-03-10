package svg

import (
	"reflect"
	"testing"
)

func TestLengthType_MarshalText(t *testing.T) {
	tests := []struct {
		name    string
		lt      LengthType
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
			"em",
			"em",
			[]byte("em"),
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

func TestLengthType_UnmarshalText(t *testing.T) {
	type args struct {
		text []byte
	}
	tests := []struct {
		name    string
		lt      LengthType
		lt2     LengthType
		args    args
		wantErr bool
	}{
		{
			"em to empty",
			"em",
			"",
			args{[]byte{}},
			false,
		},
		{
			"em to ex",
			"em",
			"ex",
			args{[]byte("ex")},
			false,
		},
		{
			"invalid is treated like empty",
			"em",
			"",
			args{[]byte{}},
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

func TestLth(t *testing.T) {
	type args struct {
		n       float64
		lengths []LengthType
	}
	tests := []struct {
		name string
		args args
		want Length
	}{
		{
			"no length type",
			args{26.4, nil},
			Length{26.4, ""},
		},
		{
			"first length type is used",
			args{26.4, []LengthType{Em, Percent}},
			Length{26.4, "em"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Lth(tt.args.n, tt.args.lengths...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("L() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLength_MarshalText(t *testing.T) {
	type fields struct {
		Number float64
		Type   LengthType
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
		{
			"em",
			fields{30, "em"},
			[]byte("30em"),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Length{
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

func TestLength_UnmarshalText(t *testing.T) {
	type args struct {
		text []byte
	}
	tests := []struct {
		name    string
		length  Length
		args    args
		want    Length
		wantErr bool
	}{
		{
			"percent to number",
			Length{5426432.9382, "%"},
			args{[]byte("65.23")},
			Length{65.23, ""},
			false,
		},
		{
			"percent to em",
			Length{-17.2321, "%"},
			args{[]byte("25em")},
			Length{25, "em"},
			false,
		},
		{
			"number to em",
			Length{6845458, ""},
			args{[]byte("3.33em")},
			Length{3.33, "em"},
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

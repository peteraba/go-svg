package attribute

import (
	"reflect"
	"testing"
)

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
		want    TextAnchor
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

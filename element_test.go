package svg

import (
	"encoding/xml"
	"sync"
	"testing"
)

func TestE(t *testing.T) {
	type args struct {
		local    string
		space    string
		text     string
		attrs    map[string]string
		children []Element
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"tag /wo long name",
			args{
				"x",
				"",
				"lorem ipsum",
				map[string]string{"foo": "Foo"},
				[]Element{
					{XMLName: xml.Name{Local: "e"}, Text: "merol muspi", lock: &sync.Mutex{}},
				},
			},
			`<x foo="Foo">lorem ipsum<e>merol muspi</e></x>`,
		},
		{
			"tag with long name",
			args{
				"x",
				"https://example.com/x",
				"lorem ipsum",
				map[string]string{"foo": "Foo"},
				[]Element{
					{XMLName: xml.Name{Local: "e"}, Text: "merol muspi", lock: &sync.Mutex{}},
				},
			},
			`<x xmlns="https://example.com/x" foo="Foo">lorem ipsum<e>merol muspi</e></x>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := E(tt.args.local, tt.args.space, tt.args.text, tt.args.attrs, tt.args.children)

			gotBytes, err := xml.Marshal(e)
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

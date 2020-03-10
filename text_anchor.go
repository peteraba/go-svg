package svg

import (
	"strings"
)

type TextAnchor int

const (
	Start TextAnchor = iota
	Middle
	End
)

func (t *TextAnchor) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {
	default:
		*t = Start
	case "middle":
		*t = Middle
	case "end":
		*t = End
	}

	return nil
}

func (t *TextAnchor) MarshalText() ([]byte, error) {
	if t == nil {
		return []byte{}, nil
	}

	var name string

	switch *t {
	default:
		name = "end"
	case Middle:
		name = "middle"
	case Start:
		name = "start"
	}

	return []byte(name), nil
}

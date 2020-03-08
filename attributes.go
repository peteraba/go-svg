package svg

import (
	"fmt"
	"image/color"
	"strings"
)

type Color struct {
	color.RGBA
}

func (c Color) String() string {
	return fmt.Sprintf("#%s%s%s", twoDigitHexa(c.R), twoDigitHexa(c.G), twoDigitHexa(c.B))
}

func (c *Color) UnmarshalText(text []byte) error {
	t := string(text)

	newColorName, err := NewColorName(t)
	if err == nil {
		*c = newColorName.ToColor()

		return nil
	}

	newColor, err := ParseHexaColor(t)
	if err != nil {
		return err
	}

	*c = newColor

	return nil
}

func (c *Color) MarshalText() ([]byte, error) {
	if c == nil {
		return []byte{}, nil
	}

	s := fmt.Sprintf("#%s%s%s", twoDigitHexa(c.R), twoDigitHexa(c.G), twoDigitHexa(c.B))

	return []byte(s), nil
}

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

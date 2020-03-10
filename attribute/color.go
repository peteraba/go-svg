package attribute

import (
	"errors"
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

	newColor, err := ColorFromHexaString(t)
	if err != nil {
		return err
	}

	*c = newColor

	return nil
}

func ColorFromHexaString(s string) (Color, error) {
	if len(s) != 4 && len(s) != 7 {
		return Color{}, errors.New("invalid hexa color length")
	}

	if s[0] != '#' {
		return Color{}, errors.New("invalid first character for hexa color")
	}

	us, err := charsToUint8(s[1:])
	if err != nil {
		return Color{}, err
	}

	return Color{RGBA: color.RGBA{R: us[0], G: us[1], B: us[2], A: 255}}, nil
}

func (c *Color) MarshalText() ([]byte, error) {
	if c == nil {
		return []byte{}, nil
	}

	s := fmt.Sprintf("#%s%s%s", twoDigitHexa(c.R), twoDigitHexa(c.G), twoDigitHexa(c.B))

	return []byte(s), nil
}

func charsToUint8(s string) ([3]uint8, error) {
	if len(s) != 3 && len(s) != 6 {
		return [3]uint8{}, errors.New("invalid hexadecimal color string")
	}

	s = strings.ToLower(s)

	var tmp []int
	for _, runeValue := range s {
		if idx := strings.IndexRune("0123456789abcdef", runeValue); idx > -1 {
			tmp = append(tmp, idx)
			if len(s) == 3 {
				tmp = append(tmp, idx)
			}
		}
	}

	if len(tmp) < 6 {
		return [3]uint8{}, errors.New("error in charsToUint8")
	}

	res := [3]uint8{}
	res[0] = uint8(tmp[0]*16 + tmp[1])
	res[1] = uint8(tmp[2]*16 + tmp[3])
	res[2] = uint8(tmp[4]*16 + tmp[5])

	return res, nil
}

func twoDigitHexa(i uint8) string {
	if i > 0xf {
		return fmt.Sprintf("%x", i)
	}

	return fmt.Sprintf("0%x", i)
}

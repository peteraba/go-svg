package svg

import (
	"fmt"
	"strconv"
	"strings"
)

type OpacityType string

const (
	OPercent OpacityType = "%"
)

func (ot OpacityType) String() string {
	t := strings.ToLower(string(ot))

	switch t {
	case string(OPercent):
		return string(OPercent)
	}

	return ""
}

func (ot *OpacityType) UnmarshalText(text []byte) error {
	t := strings.ToLower(string(text))
	switch t {
	default:
		*ot = ""
	case string(OPercent):
		*ot = OPercent
	}
	return nil
}

func (ot OpacityType) MarshalText() ([]byte, error) {
	s := ot.String()

	return []byte(s), nil
}

type Opacity struct {
	Number float64
	Type   OpacityType
}

func O(n float64, lengths ...OpacityType) Opacity {
	var t OpacityType
	if len(lengths) > 0 {
		t = lengths[0]
	}

	return Opacity{n, t}
}

func (o *Opacity) UnmarshalText(text []byte) error {
	t := strings.ToLower(string(text))

	if len(t) < 1 {
		o.Type = ""
		o.Number = 0

		return nil
	}

	if t[len(t)-1] == '%' {
		o.Type = OPercent
		t = t[:len(t)-1]
	}

	n, err := strconv.ParseFloat(t, 64)
	if err != nil {
		return err
	}

	o.Number = n

	return nil
}

func (o Opacity) String() string {
	return fmt.Sprintf("%v%s", o.Number, o.Type)
}

func (o Opacity) MarshalText() ([]byte, error) {
	s := o.String()

	return []byte(s), nil
}

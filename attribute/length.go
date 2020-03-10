package attribute

import (
	"fmt"
	"strconv"
	"strings"
)

type LengthType string

const (
	Em      LengthType = "em"
	Ex      LengthType = "ex"
	Px      LengthType = "px"
	In      LengthType = "in"
	Cm      LengthType = "cm"
	Mm      LengthType = "mm"
	Pt      LengthType = "pt"
	Pc      LengthType = "pc"
	Percent LengthType = "%"
)

func (lt LengthType) String() string {
	t := strings.ToLower(string(lt))

	switch t {
	case string(Em):
		return string(Em)
	case string(Ex):
		return string(Em)
	case string(Px):
		return string(Em)
	case string(In):
		return string(Em)
	case string(Cm):
		return string(Em)
	case string(Mm):
		return string(Em)
	case string(Pt):
		return string(Em)
	case string(Pc):
		return string(Em)
	case string(Percent):
		return string(Percent)
	}

	return ""
}

func (lt *LengthType) UnmarshalText(text []byte) error {
	t := strings.ToLower(string(text))
	switch t {
	default:
		*lt = ""
	case string(Em):
		*lt = Em
	case string(Ex):
		*lt = Ex
	case string(Px):
		*lt = Px
	case string(In):
		*lt = In
	case string(Cm):
		*lt = Cm
	case string(Mm):
		*lt = Mm
	case string(Pt):
		*lt = Pt
	case string(Pc):
		*lt = Pc
	case string(Percent):
		*lt = Percent
	}
	return nil
}

func (lt LengthType) MarshalText() ([]byte, error) {
	s := lt.String()

	return []byte(s), nil
}

type Length struct {
	Number float64
	Type   LengthType
}

func L(n float64, lengths ...LengthType) Length {
	var t LengthType
	if len(lengths) > 0 {
		t = lengths[0]
	}

	return Length{n, t}
}

func (l *Length) UnmarshalText(text []byte) error {
	t := strings.ToLower(string(text))

	if len(t) < 1 {
		l.Type = ""
		l.Number = 0

		return nil
	}

	if t[len(t)-1] == '%' {
		l.Type = Percent
		t = t[:len(t)-1]
	} else if len(t) > 2 {
		end := t[len(t)-2:]
		_ = l.Type.UnmarshalText([]byte(end))
		if l.Type != LengthType("") {
			t = t[:len(t)-2]
		}
	}

	n, err := strconv.ParseFloat(t, 64)
	if err != nil {
		return err
	}

	l.Number = n

	return nil
}

func (l Length) String() string {
	return fmt.Sprintf("%v%s", l.Number, l.Type)
}

func (l Length) MarshalText() ([]byte, error) {
	s := l.String()

	return []byte(s), nil
}

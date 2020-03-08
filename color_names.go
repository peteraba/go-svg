package svg

import (
	"fmt"
)

type ColorName string

const (
	LightSalmon          ColorName = "lightsalmon"
	Salmon                         = "salmon"
	DarkSalmon                     = "darksalmon"
	LightCoral                     = "lightcoral"
	IndianRed                      = "indianred"
	Crimson                        = "crimson"
	Firebrick                      = "firebrick"
	Red                            = "red"
	DarkRed                        = "darkred"
	Coral                          = "coral"
	Tomato                         = "tomato"
	OrangeRed                      = "orangered"
	Gold                           = "gold"
	Orange                         = "orange"
	DarkOrange                     = "darkorange"
	LightYellow                    = "lightyellow"
	Lemonchiffon                   = "lemonchiffon"
	LightGoldenrodYellow           = "lightgoldenrodyellow"
	PapayaWhip                     = "papayawhip"
	Moccasin                       = "moccasin"
	Peachpuff                      = "peachpuff"
	PaleGoldenrod                  = "palegoldenrod"
	Khaki                          = "khaki"
	DarkKhaki                      = "darkkhaki"
	Yellow                         = "yellow"
	LawnGreen                      = "lawngreen"
	Chartreuse                     = "chartreuse"
	LimeGreen                      = "limegreen"
	Lime                           = "lime"
	ForestGreen                    = "forestgreen"
	Green                          = "green"
	DarkGreen                      = "darkgreen"
	GreenYellow                    = "greenyellow"
	YellowGreen                    = "yellowgreen"
	SpringGreen                    = "springgreen"
	MediumSpringGreen              = "mediumspringgreen"
	LightGreen                     = "lightgreen"
	PaleGreen                      = "palegreen"
	DarkSeaGreen                   = "darkseagreen"
	MediumSeaGreen                 = "mediumseagreen"
	SeaGreen                       = "seagreen"
	Olive                          = "olive"
	DarkOliveGreen                 = "darkolivegreen"
	OliveDrab                      = "olivedrab"
	LightCyan                      = "lightcyan"
	Cyan                           = "cyan"
	Aqua                           = "aqua"
	Aquamarine                     = "aquamarine"
	MediumAquamarine               = "mediumaquamarine"
	PaleTurquoise                  = "paleturquoise"
	Turquoise                      = "turquoise"
	MediumTurquoise                = "mediumturquoise"
	DarkTurquoise                  = "darkturquoise"
	LightSeaGreen                  = "lightseagreen"
	CadetBlue                      = "cadetblue"
	DarkCyan                       = "darkcyan"
	Teal                           = "teal"
	PowderBlue                     = "powderblue"
	LightBlue                      = "lightblue"
	LightskyBlue                   = "lightskyblue"
	SkyBlue                        = "skyblue"
	DeepskyBlue                    = "deepskyblue"
	LightsteelBlue                 = "lightsteelblue"
	DodgerBlue                     = "dodgerblue"
	CornflowerBlue                 = "cornflowerblue"
	SteelBlue                      = "steelblue"
	RoyalBlue                      = "royalblue"
	Blue                           = "blue"
	MediumBlue                     = "mediumblue"
	DarkBlue                       = "darkblue"
	Navy                           = "navy"
	MidnightBlue                   = "midnightblue"
	MediumslateBlue                = "mediumslateblue"
	SlateBlue                      = "slateblue"
	DarkslateBlue                  = "darkslateblue"
	Lavender                       = "lavender"
	Thistle                        = "thistle"
	Plum                           = "plum"
	Violet                         = "violet"
	Orchid                         = "orchid"
	Fuchsia                        = "fuchsia"
	Magenta                        = "magenta"
	MediumOrchid                   = "mediumorchid"
	Mediumpurple                   = "mediumpurple"
	BlueViolet                     = "blueviolet"
	DarkViolet                     = "darkviolet"
	DarkOrchid                     = "darkorchid"
	Darkmagenta                    = "darkmagenta"
	Purple                         = "purple"
	Indigo                         = "indigo"
	Pink                           = "pink"
	LightPink                      = "lightpink"
	HotPink                        = "hotpink"
	DeepPink                       = "deeppink"
	PaleVioletRed                  = "palevioletred"
	MediumVioletRed                = "mediumvioletred"
	White                          = "white"
	Snow                           = "snow"
	Honeydew                       = "honeydew"
	MintCream                      = "mintcream"
	Azure                          = "azure"
	AliceBlue                      = "aliceblue"
	GhostWhite                     = "ghostwhite"
	WhiteSmoke                     = "whitesmoke"
	Seashell                       = "seashell"
	Beige                          = "beige"
	OldLace                        = "oldlace"
	FloralWhite                    = "floralwhite"
	Ivory                          = "ivory"
	AntiqueWhite                   = "antiquewhite"
	Linen                          = "linen"
	LavenderBlush                  = "lavenderblush"
	MistyRose                      = "mistyrose"
	Gainsboro                      = "gainsboro"
	LightGray                      = "lightgray"
	Silver                         = "silver"
	DarkGray                       = "darkgray"
	Gray                           = "gray"
	DimGray                        = "dimgray"
	LightslateGray                 = "lightslategray"
	SlateGray                      = "slategray"
	DarkslateGray                  = "darkslategray"
	Black                          = "black"
	Cornsilk                       = "cornsilk"
	BlanchedAlmond                 = "blanchedalmond"
	Bisque                         = "bisque"
	NavajoWhite                    = "navajowhite"
	Wheat                          = "wheat"
	BurlyWood                      = "burlywood"
	Tan                            = "tan"
	RosyBrown                      = "rosybrown"
	SandyBrown                     = "sandybrown"
	Goldenrod                      = "goldenrod"
	Peru                           = "peru"
	Chocolate                      = "chocolate"
	SaddleBrown                    = "saddlebrown"
	Sienna                         = "sienna"
	Brown                          = "brown"
	Maroon                         = "maroon"
)

var nameToHexa = map[ColorName]string{
	// Reds
	LightSalmon:          "#ffa07a",
	Salmon:               "#fa8072",
	DarkSalmon:           "#e9967a",
	LightCoral:           "#f08080",
	IndianRed:            "#cd5c5c",
	Crimson:              "#dc143c",
	Firebrick:            "#b22222",
	Red:                  "#ff0000",
	DarkRed:              "#8b0000",
	// Oranges
	Coral:                "#ff7f50",
	Tomato:               "#ff6347",
	OrangeRed:            "#ff4500",
	Gold:                 "#ffd700",
	Orange:               "#ffa500",
	DarkOrange:           "#ff8c00",
	// Yellows
	LightYellow:          "#ffffe0",
	Lemonchiffon:         "#fffacd",
	LightGoldenrodYellow: "#fafad2",
	PapayaWhip:           "#ffefd5",
	Moccasin:             "#ffe4b5",
	Peachpuff:            "#ffdab9",
	PaleGoldenrod:        "#eee8aa",
	Khaki:                "#f0e68c",
	DarkKhaki:            "#bdb76b",
	Yellow:               "#ffff00",
	// Greens
	LawnGreen:            "#7cfc00",
	Chartreuse:           "#7fff00",
	LimeGreen:            "#32cd32",
	Lime:                 "#00ff00",
	ForestGreen:          "#228b22",
	Green:                "#008000",
	DarkGreen:            "#006400",
	GreenYellow:          "#adff2f",
	YellowGreen:          "#9acd32",
	SpringGreen:          "#00ff7f",
	MediumSpringGreen:    "#00fa9a",
	LightGreen:           "#90ee90",
	PaleGreen:            "#98fb98",
	DarkSeaGreen:         "#8fbc8f",
	MediumSeaGreen:       "#3cb371",
	SeaGreen:             "#2e8b57",
	Olive:                "#808000",
	DarkOliveGreen:       "#556b2f",
	OliveDrab:            "#6b8e23",
	// Cyans
	LightCyan:            "#e0ffff",
	Cyan:                 "#00ffff",
	Aqua:                 "#00ffff",
	Aquamarine:           "#7fffd4",
	MediumAquamarine:     "#66cdaa",
	PaleTurquoise:        "#afeeee",
	Turquoise:            "#40e0d0",
	MediumTurquoise:      "#48d1cc",
	DarkTurquoise:        "#00ced1",
	LightSeaGreen:        "#20b2aa",
	CadetBlue:            "#5f9ea0",
	DarkCyan:             "#008b8b",
	Teal:                 "#008080",
	// Blues
	PowderBlue:           "#b0e0e6",
	LightBlue:            "#add8e6",
	LightskyBlue:         "#87cefa",
	SkyBlue:              "#87ceeb",
	DeepskyBlue:          "#00bfff",
	LightsteelBlue:       "#b0c4de",
	DodgerBlue:           "#1e90ff",
	CornflowerBlue:       "#6495ed",
	SteelBlue:            "#4682b4",
	RoyalBlue:            "#4169e1",
	Blue:                 "#0000ff",
	MediumBlue:           "#0000cd",
	DarkBlue:             "#00008b",
	Navy:                 "#000080",
	MidnightBlue:         "#191970",
	MediumslateBlue:      "#7b68ee",
	SlateBlue:            "#6a5acd",
	DarkslateBlue:        "#483d8b",
	// Purples
	Lavender:             "#e6e6fa",
	Thistle:              "#d8bfd8",
	Plum:                 "#dda0dd",
	Violet:               "#ee82ee",
	Orchid:               "#da70d6",
	Fuchsia:              "#ff00ff",
	Magenta:              "#ff00ff",
	MediumOrchid:         "#ba55d3",
	Mediumpurple:         "#9370db",
	BlueViolet:           "#8a2be2",
	DarkViolet:           "#9400d3",
	DarkOrchid:           "#9932cc",
	Darkmagenta:          "#8b008b",
	Purple:               "#800080",
	Indigo:               "#4b0082",
	// Pinks
	Pink:                 "#ffc0cb",
	LightPink:            "#ffb6c1",
	HotPink:              "#ff69b4",
	DeepPink:             "#ff1493",
	PaleVioletRed:        "#db7093",
	MediumVioletRed:      "#c71585",
	// Whites
	White:                "#ffffff",
	Snow:                 "#fffafa",
	Honeydew:             "#f0fff0",
	MintCream:            "#f5fffa",
	Azure:                "#f0ffff",
	AliceBlue:            "#f0f8ff",
	GhostWhite:           "#f8f8ff",
	WhiteSmoke:           "#f5f5f5",
	Seashell:             "#fff5ee",
	Beige:                "#f5f5dc",
	OldLace:              "#fdf5e6",
	FloralWhite:          "#fffaf0",
	Ivory:                "#fffff0",
	AntiqueWhite:         "#faebd7",
	Linen:                "#faf0e6",
	LavenderBlush:        "#fff0f5",
	MistyRose:            "#ffe4e1",
	// Grays
	Gainsboro:            "#dcdcdc",
	LightGray:            "#d3d3d3",
	Silver:               "#c0c0c0",
	DarkGray:             "#a9a9a9",
	Gray:                 "#808080",
	DimGray:              "#696969",
	LightslateGray:       "#778899",
	SlateGray:            "#708090",
	DarkslateGray:        "#2f4f4f",
	Black:                "#000000",
	// Browns
	Cornsilk:             "#fff8dc",
	BlanchedAlmond:       "#ffebcd",
	Bisque:               "#ffe4c4",
	NavajoWhite:          "#ffdead",
	Wheat:                "#f5deb3",
	BurlyWood:            "#deb887",
	Tan:                  "#d2b48c",
	RosyBrown:            "#bc8f8f",
	SandyBrown:           "#f4a460",
	Goldenrod:            "#daa520",
	Peru:                 "#cd853f",
	Chocolate:            "#d2691e",
	SaddleBrown:          "#8b4513",
	Sienna:               "#a0522d",
	Brown:                "#a52a2a",
	Maroon:               "#800000",
}

func NewColorName(cn string) (ColorName, error) {
	_, ok := nameToHexa[ColorName(cn)]
	if !ok {
		return "", fmt.Errorf("invalid color name: %s", cn)
	}

	return ColorName(cn), nil
}

func (cn ColorName) ToColor() Color {
	h := cn.ToHexa()

	c, err := ParseHexaColor(h)
	if err != nil {
		panic(err)
	}

	return c
}

func (cn ColorName) ToHexa() string {
	h, ok := nameToHexa[cn]
	if !ok {
		panic(fmt.Errorf("invalid color name: %s", string(cn)))
	}

	return h
}

package skittles

import (
	"fmt"
)

type Attribute int

const (
	black Attribute = iota + 30
	red
	green
	yellow
	blue
	magenta
	cyan
	white
	regular   = ""
	bold      = "1;"
	blink     = "5;"
	underline = "4;"
	inverse   = "7;"
)

func makeString(attr, text string, color Attribute) string {
	return fmt.Sprintf("\033[%s%dm%s\033[0m", attr, color, text)
}

func Black(text string) string {
	return makeString(regular, text, black)
}

func Red(text string) string {
	return makeString(regular, text, red)
}

func Green(text string) string {
	return makeString(regular, text, green)
}

func Yellow(text string) string {
	return makeString(regular, text, yellow)
}

func Blue(text string) string {
	return makeString(regular, text, blue)
}

func Magenta(text string) string {
	return makeString(regular, text, magenta)
}

func Cyan(text string) string {
	return makeString(regular, text, cyan)
}

func White(text string) string {
	return makeString(regular, text, white)
}

func BoldBlack(text string) string {
	return makeString(bold, text, black)
}

func BoldRed(text string) string {
	return makeString(bold, text, red)
}

func BoldGreen(text string) string {
	return makeString(bold, text, green)
}

func BoldYellow(text string) string {
	return makeString(bold, text, yellow)
}

func BoldBlue(text string) string {
	return makeString(bold, text, blue)
}

func BoldMagenta(text string) string {
	return makeString(bold, text, magenta)
}

func BoldCyan(text string) string {
	return makeString(bold, text, cyan)
}

func BoldWhite(text string) string {
	return makeString(bold, text, white)
}

func BlinkBlack(text string) string {
	return makeString(blink, text, black)
}

func BlinkRed(text string) string {
	return makeString(blink, text, red)
}

func BlinkGreen(text string) string {
	return makeString(blink, text, green)
}

func BlinkYellow(text string) string {
	return makeString(blink, text, yellow)
}

func BlinkBlue(text string) string {
	return makeString(blink, text, blue)
}

func BlinkMagenta(text string) string {
	return makeString(blink, text, magenta)
}

func BlinkCyan(text string) string {
	return makeString(blink, text, cyan)
}

func BlinkWhite(text string) string {
	return makeString(blink, text, white)
}

func UnderlineBlack(text string) string {
	return makeString(underline, text, black)
}

func UnderlineRed(text string) string {
	return makeString(underline, text, red)
}

func UnderlineGreen(text string) string {
	return makeString(underline, text, green)
}

func UnderlineYellow(text string) string {
	return makeString(underline, text, yellow)
}

func UnderlineBlue(text string) string {
	return makeString(underline, text, blue)
}

func UnderlineMagenta(text string) string {
	return makeString(underline, text, magenta)
}

func UnderlineCyan(text string) string {
	return makeString(underline, text, cyan)
}

func UnderlineWhite(text string) string {
	return makeString(underline, text, white)
}

func InverseBlack(text string) string {
	return makeString(inverse, text, black)
}

func InverseRed(text string) string {
	return makeString(inverse, text, red)
}

func InverseGreen(text string) string {
	return makeString(inverse, text, green)
}

func InverseYellow(text string) string {
	return makeString(inverse, text, yellow)
}

func InverseBlue(text string) string {
	return makeString(inverse, text, blue)
}

func InverseMagenta(text string) string {
	return makeString(inverse, text, magenta)
}

func InverseCyan(text string) string {
	return makeString(inverse, text, cyan)
}

func InverseWhite(text string) string {
	return makeString(inverse, text, white)
}

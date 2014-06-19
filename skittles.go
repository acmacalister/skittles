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

// black text \033[30m
func Black(text string) string {
	return makeString(regular, text, black)
}

// red text \033[30m
func Red(text string) string {
	return makeString(regular, text, red)
}

// green text \033[30m
func Green(text string) string {
	return makeString(regular, text, green)
}

// yellow text \033[30m
func Yellow(text string) string {
	return makeString(regular, text, yellow)
}

// blue text \033[30m
func Blue(text string) string {
	return makeString(regular, text, blue)
}

// magenta text \033[30m
func Magenta(text string) string {
	return makeString(regular, text, magenta)
}

// cyan text \033[30m
func Cyan(text string) string {
	return makeString(regular, text, cyan)
}

// white text \033[30m
func White(text string) string {
	return makeString(regular, text, white)
}

// bold black text \033[1;30m
func BoldBlack(text string) string {
	return makeString(bold, text, black)
}

// bold red text \1;033[30m
func BoldRed(text string) string {
	return makeString(bold, text, red)
}

// bold green text \033[1;30m
func BoldGreen(text string) string {
	return makeString(bold, text, green)
}

// bold yellow text \033[1;30m
func BoldYellow(text string) string {
	return makeString(bold, text, yellow)
}

// bold blue text \033[1;30m
func BoldBlue(text string) string {
	return makeString(bold, text, blue)
}

// bold magenta text \033[1;30m
func BoldMagenta(text string) string {
	return makeString(bold, text, magenta)
}

// bold cyan text \033[1;30m
func BoldCyan(text string) string {
	return makeString(bold, text, cyan)
}

// bold white text \033[1;30m
func BoldWhite(text string) string {
	return makeString(bold, text, white)
}

// blink black text \033[5;30m
func BlinkBlack(text string) string {
	return makeString(blink, text, black)
}

// blink red text \033[5;30m
func BlinkRed(text string) string {
	return makeString(blink, text, red)
}

// blink green text \033[5;30m
func BlinkGreen(text string) string {
	return makeString(blink, text, green)
}

// blink yellow text \033[5;30m
func BlinkYellow(text string) string {
	return makeString(blink, text, yellow)
}

// blink blue text \033[5;30m
func BlinkBlue(text string) string {
	return makeString(blink, text, blue)
}

// blink magenta text \033[5;30m
func BlinkMagenta(text string) string {
	return makeString(blink, text, magenta)
}

// blink cyan text \033[5;30m
func BlinkCyan(text string) string {
	return makeString(blink, text, cyan)
}

// blink white text \033[5;30m
func BlinkWhite(text string) string {
	return makeString(blink, text, white)
}

// underline black text \033[4;30m
func UnderlineBlack(text string) string {
	return makeString(underline, text, black)
}

// underline red text \033[4;30m
func UnderlineRed(text string) string {
	return makeString(underline, text, red)
}

// underline green text \033[4;30m
func UnderlineGreen(text string) string {
	return makeString(underline, text, green)
}

// underline yellow text \033[4;30m
func UnderlineYellow(text string) string {
	return makeString(underline, text, yellow)
}

// underline blue text \033[4;30m
func UnderlineBlue(text string) string {
	return makeString(underline, text, blue)
}

// underline magenta text \033[4;30m
func UnderlineMagenta(text string) string {
	return makeString(underline, text, magenta)
}

// underline cyan text \033[4;30m
func UnderlineCyan(text string) string {
	return makeString(underline, text, cyan)
}

// underline white text \033[4;30m
func UnderlineWhite(text string) string {
	return makeString(underline, text, white)
}

// inverse black text \033[7;30m
func InverseBlack(text string) string {
	return makeString(inverse, text, black)
}

// inverse red text \033[7;30m
func InverseRed(text string) string {
	return makeString(inverse, text, red)
}

// inverse green text \033[7;30m
func InverseGreen(text string) string {
	return makeString(inverse, text, green)
}

// inverse yellow text \033[7;30m
func InverseYellow(text string) string {
	return makeString(inverse, text, yellow)
}

// inverse blue text \033[7;30m
func InverseBlue(text string) string {
	return makeString(inverse, text, blue)
}

// inverse magenta text \033[7;30m
func InverseMagenta(text string) string {
	return makeString(inverse, text, magenta)
}

// inverse cyan text \033[7;30m
func InverseCyan(text string) string {
	return makeString(inverse, text, cyan)
}

// inverse white text \033[7;30m
func InverseWhite(text string) string {
	return makeString(inverse, text, white)
}

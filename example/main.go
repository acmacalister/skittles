package main

import (
	"fmt"
	"github.com/acmacalister/skittles"
)

func main() {
	fmt.Printf("%s%s%s%s%s%s%s%s\n", skittles.Red("S"), skittles.Magenta("k"), skittles.Yellow("i"),
		skittles.Green("t"), skittles.Blue("t"), skittles.Cyan("l"), skittles.Red("e"), skittles.Magenta("s"))
	fmt.Printf("%s %s %s%s\n", skittles.BoldRed("Print"), skittles.BoldGreen("the"),
		skittles.BoldCyan("rainbow"), skittles.BoldMagenta("!!"))
}

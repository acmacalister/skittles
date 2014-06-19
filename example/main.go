package main

import (
	"fmt"
	"github.com/acmacalister/skittles"
	"log"
)

func main() {
	fmt.Printf("%s %s\n", skittles.Green("Logging some stuff"), skittles.Blue("and other stuff"))
	log.Println(skittles.BlinkRed("Error!!"))
}

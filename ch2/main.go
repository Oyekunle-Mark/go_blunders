package main

import (
	"flag"
	"fmt"
	"go_blunders/ch2/tempconv"
)

var temperature = flag.Float64("temperature", 0.0, "The temperature to be converted")

func main() {
	flag.Parse()

	fmt.Printf(
		"%s = %s, %s = %s\n",
		tempconv.Celsius(*temperature),
		tempconv.CToF(tempconv.Celsius(*temperature)),
		tempconv.Fahrenheit(*temperature),
		tempconv.FToC(tempconv.Fahrenheit(*temperature)),
	)
}

package main

import (
	"flag"
	"fmt"
	"go_blunders/ch2/tempconv"
)

var temp = flag.Float64("temp", 0.0, "The temperature to be converted")

func main() {
	flag.Parse()

	fmt.Printf(
		"%s = %s, %s = %s\n",
		tempconv.Celsius(*temp),
		tempconv.CToF(tempconv.Celsius(*temp)),
		tempconv.Fahrenheit(*temp),
		tempconv.FToC(tempconv.Fahrenheit(*temp)),
	)
}

package main

import (
	"conv/lenconv"
	"conv/tempconv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	numArgs := len(os.Args[1:])
	for idx, arg := range os.Args[1:] {
		x, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		tempC := tempconv.Celsius(x)
		tempF := tempconv.Fahrenheit(x)
		lenM := lenconv.Meter(x)
		lenF := lenconv.Foot(x)
		fmt.Printf("%s = %s, %s = %s\n",
			tempC, tempconv.CToF(tempC), tempF, tempconv.FToC(tempF))
		fmt.Printf("%s = %s, %s = %s\n",
			lenM, lenconv.MToF(lenM), lenF, lenconv.FToM(lenF))
		if idx+1 < numArgs {
			fmt.Println("-----")
		}
	}
}

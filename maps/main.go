package main

import "fmt"

func main() {
	colors := map[string]string{
		"red": "#FF0000",
		"green": "#008000",
	}

	printColorMap(colors)
}

func printColorMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex value for color", color, "is", hex)
	}
}
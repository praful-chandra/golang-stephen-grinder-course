package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
		"white": "#ffffff",
		"black": "#000000",
	}

	// var colors1 map[string]string

	// colors2 := make(map[string]string)

	// colors2["new Property"] = "Hey"
	// colors2["second Property"] = "Hello"

	// fmt.Println(colors)
	// fmt.Println(colors1)
	// fmt.Println(colors2)

	printColorsMap(colors)

}

func printColorsMap(cols map[string]string) {
	for color, hex := range cols {
		fmt.Println("Hex code for color \t", color, "\t is \t", hex)
	}
}

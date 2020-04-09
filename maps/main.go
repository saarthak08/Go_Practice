package main

import (
	"fmt"
)

func main() {
	colors:=map[string]string{
		"red":"##ff0000",
		"blue":"4bf745",
	}

	map2:= make(map[string]int)
	map2["hello"]=123
	map2["bye"]=12
	delete(map2,"bye")
	fmt.Println(map2)

	fmt.Println(colors)

	printMap(colors)
}

func printMap(m map[string]string) {
	for color, code := range m {
		fmt.Println("Color:",color," Code:", code)
	}
}
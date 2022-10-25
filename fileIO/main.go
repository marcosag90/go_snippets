package main

import "fmt"

func main() {
	// var colors map[string]string
	// colors := map[string]string {
	//	"red": "#ff0000",
	//	"white": "#ffffff"
	//}
	colors := make(map[string]string)

	colors["white"] = "#ffffff"
	colors["red"] = "#ff0000"
	fmt.Println(colors["white"])
	fmt.Printf("%+v\n", colors)
	delete(colors, "red")
	fmt.Printf("%+v\n", colors)

}

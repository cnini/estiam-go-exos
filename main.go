package main

import "fmt"

func Add(d map[string]string, key string, value string) {
	d[key] = value
}

func main() {
	d := make(map[string]string)

	// Add to map
	Add(d, "pomme", "apple")   // d["pomme"] = "apple"
	Add(d, "banane", "banana") // d["banane"] = "banana"
	Add(d, "voiture", "car")   // d["voiture"] = "car"
	Add(d, "poire", "pear")    // d["poire"] = "pear"

	fmt.Println("d:", d)
}

package main

import (
	"estiam-go-exos/dictionary"
	"fmt"
)

func main() {
	d := make(map[string]string)

	// Add to map
	dictionary.Add(d, "pomme", "apple")   // d["pomme"] = "apple"
	dictionary.Add(d, "banane", "banana") // d["banane"] = "banana"
	dictionary.Add(d, "voiture", "car")   // d["voiture"] = "car"
	dictionary.Add(d, "poire", "pear")    // d["poire"] = "pear"

	// Get value from existing key
	getBananeVal, getBananeErr := dictionary.Get(d, "banane")
	if getBananeErr != nil {
		fmt.Println("Get banane (error):", getBananeErr.Error())
	} else {
		fmt.Println("Get banane:", getBananeVal)
	}

	// Get value from non-existing key
	getFraiseVal, getFraiseErr := dictionary.Get(d, "fraise")
	if getFraiseErr != nil {
		fmt.Println("Get fraise (error):", getFraiseErr.Error())
	} else {
		fmt.Println("Get fraise:", getFraiseVal)
	}

	// Remove key/value pair from existing key
	removeVoitureMess, removeVoitureErr := dictionary.Remove(d, "voiture")
	if removeVoitureErr != nil {
		fmt.Println("Remove voiture (error):", removeVoitureErr.Error())
	} else {
		fmt.Println("Remove voiture:", removeVoitureMess)
	}

	// Remove key/value pair from non-existing key
	removeFraiseMess, removeFraiseErr := dictionary.Remove(d, "fraise")
	if removeFraiseErr != nil {
		fmt.Println("Remove fraise (error):", removeFraiseErr.Error())
	} else {
		fmt.Println("Remove fraise:", removeFraiseMess)
	}
}

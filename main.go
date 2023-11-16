package main

import "fmt"

func Add(d map[string]string, key string, value string) {
	d[key] = value
}

func Get(d map[string]string, key string) (string, error) {
	if value, exists := d[key]; exists {
		return value, nil
	} else {
		return "", fmt.Errorf("Value for %s key does not exist in dictionary", key)
	}
}

func main() {
	d := make(map[string]string)

	// Add to map
	Add(d, "pomme", "apple")   // d["pomme"] = "apple"
	Add(d, "banane", "banana") // d["banane"] = "banana"
	Add(d, "voiture", "car")   // d["voiture"] = "car"
	Add(d, "poire", "pear")    // d["poire"] = "pear"

	// Get value from existing key
	bVal, bErr := Get(d, "banane")
	if bErr != nil {
		fmt.Println("banane (error):", bErr.Error())
	} else {
		fmt.Println("banane:", bVal)
	}

	// Get value from non-existing key
	fVal, fErr := Get(d, "fraise")
	if fErr != nil {
		fmt.Println("fraise (error):", fErr.Error())
	} else {
		fmt.Println("fraise:", fVal)
	}
}

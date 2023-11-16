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

func Remove(d map[string]string, key string) (string, error) {
	_, exists := Get(d, key)
	if exists != nil {
		return "", exists
	} else {
		delete(d, key)
		return key + " key deleted", nil
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
	getBananeVal, getBananeErr := Get(d, "banane")
	if getBananeErr != nil {
		fmt.Println("Get banane (error):", getBananeErr.Error())
	} else {
		fmt.Println("Get banane:", getBananeVal)
	}

	// Get value from non-existing key
	getFraiseVal, getFraiseErr := Get(d, "fraise")
	if getFraiseErr != nil {
		fmt.Println("Get fraise (error):", getFraiseErr.Error())
	} else {
		fmt.Println("Get fraise:", getFraiseVal)
	}

	// Remove key/value pair from existing key
	removeVoitureMess, removeVoitureErr := Remove(d, "voiture")
	if removeVoitureErr != nil {
		fmt.Println("Remove voiture (error):", removeVoitureErr.Error())
	} else {
		fmt.Println("Remove voiture:", removeVoitureMess)
	}

	// Remove key/value pair from non-existing key
	removeFraiseMess, removeFraiseErr := Remove(d, "fraise")
	if removeFraiseErr != nil {
		fmt.Println("Remove fraise (error):", removeFraiseErr.Error())
	} else {
		fmt.Println("Remove fraise:", removeFraiseMess)
	}
}

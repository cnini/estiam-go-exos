package main

import (
	"estiam-go-exos/dictionary"
	"fmt"
	"os"
)

func main() {
	file, openErr := os.OpenFile("dictionary.txt", os.O_RDWR|os.O_APPEND, 0777)
	dictionary.Check(openErr)

	// Add keys and values to a file
	dictionary.Add(file, "pomme", "apple")
	dictionary.Add(file, "banane", "banana")
	dictionary.Add(file, "voiture", "car")
	dictionary.Add(file, "poire", "pear")

	// Get value from existing key
	getBananeVal, _, getBananeErr := dictionary.Get(file, "banane")
	if getBananeErr != nil {
		fmt.Println("Get 'banana' from key (error):", getBananeErr.Error())
	} else {
		fmt.Println("Get 'banana' from key:", getBananeVal)
	}

	// Get key from existing value
	getAppleKey, _, getAppleErr := dictionary.Get(file, "apple")
	if getAppleErr != nil {
		fmt.Println("Get 'pomme' from value (error):", getAppleErr.Error())
	} else {
		fmt.Println("Get 'pomme' from value:", getAppleKey)
	}

	// Get value from non-existing key
	getFraiseVal, _, getFraiseErr := dictionary.Get(file, "fraise")
	if getFraiseErr != nil {
		fmt.Println("Get 'strawberry' from key (error):", getFraiseErr.Error())
	} else {
		fmt.Println("Get 'strawberry' from key:", getFraiseVal)
	}

	file.Close()

	// Remove key/value pair from existing key
	// removeVoitureMess, removeVoitureErr := dictionary.Remove(d, "voiture")
	// if removeVoitureErr != nil {
	// 	fmt.Println("Remove voiture (error):", removeVoitureErr.Error())
	// } else {
	// 	fmt.Println("Remove voiture:", removeVoitureMess)
	// }

	// Remove key/value pair from non-existing key
	// removeFraiseMess, removeFraiseErr := dictionary.Remove(d, "fraise")
	// if removeFraiseErr != nil {
	// 	fmt.Println("Remove fraise (error):", removeFraiseErr.Error())
	// } else {
	// 	fmt.Println("Remove fraise:", removeFraiseMess)
	// }

	// List separately the keys and the values of the dictionary
	// keys, values := dictionary.List(d)
	// fmt.Println("List all the keys:", keys)
	// fmt.Println("List all the values:", values)
}

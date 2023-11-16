package main

import (
	"estiam-go-exos/dictionary"
	"os"
)

func main() {
	file, openErr := os.OpenFile("dictionary.txt", os.O_RDWR|os.O_APPEND, 0666)
	dictionary.Check(openErr)

	// Add keys and values to a file
	dictionary.Add(file, "pomme", "apple")
	dictionary.Add(file, "banane", "banana")
	dictionary.Add(file, "voiture", "car")
	dictionary.Add(file, "poire", "pear")

	file.Close()

	// Get value from existing key
	// getBananeVal, getBananeErr := dictionary.Get(d, "banane")
	// if getBananeErr != nil {
	// 	fmt.Println("Get banane (error):", getBananeErr.Error())
	// } else {
	// 	fmt.Println("Get banane:", getBananeVal)
	// }

	// Get value from non-existing key
	// getFraiseVal, getFraiseErr := dictionary.Get(d, "fraise")
	// if getFraiseErr != nil {
	// 	fmt.Println("Get fraise (error):", getFraiseErr.Error())
	// } else {
	// 	fmt.Println("Get fraise:", getFraiseVal)
	// }

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

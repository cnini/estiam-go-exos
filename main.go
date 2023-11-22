package main

import (
	"estiam-go-exos/server"
)

func main() {
	server.Start()
	server.ListenAndServe()

	// d := make(map[string]string)

	// d["pomme"] = "apple"
	// d["banane"] = "banana"
	// d["voiture"] = "car"
	// d["poire"] = "pear"

	// // Get value from existing key
	// _, getBananeVal, getBananeErr := dictionary.Get(file, "banane")
	// if getBananeErr != nil {
	// 	fmt.Println("Get 'banana' from key (error):", getBananeErr.Error())
	// } else {
	// 	fmt.Println("Get 'banana' from key:", getBananeVal)
	// }

	// // Get key from existing value
	// getAppleKey, _, getAppleErr := dictionary.Get(file, "apple")
	// if getAppleErr != nil {
	// 	fmt.Println("Get 'pomme' from value (error):", getAppleErr.Error())
	// } else {
	// 	fmt.Println("Get 'pomme' from value:", getAppleKey)
	// }

	// // Get value from non-existing key
	// getFraiseVal, _, getFraiseErr := dictionary.Get(file, "fraise")
	// if getFraiseErr != nil {
	// 	fmt.Println("Get 'strawberry' from key (error):", getFraiseErr.Error())
	// } else {
	// 	fmt.Println("Get 'strawberry' from key:", getFraiseVal)
	// }

	// // Remove key/value pair from existing key
	// removeVoitureMess, removeVoitureErr := dictionary.Remove(d, file, "voiture")
	// if removeVoitureErr != nil {
	// 	fmt.Println("Remove voiture (error):", removeVoitureErr.Error())
	// } else {
	// 	fmt.Println("Remove voiture:", removeVoitureMess)
	// }

	// newFile, newOpenErr := os.OpenFile("dictionary.txt", os.O_RDWR|os.O_APPEND, 0777)
	// dictionary.Check(newOpenErr)

	// defer newFile.Close()

	// // List separately the keys and the values of the dictionary
	// keys, values := dictionary.List(d, newFile)
	// fmt.Println("List all the keys:", keys)
	// fmt.Println("List all the values:", values)
}

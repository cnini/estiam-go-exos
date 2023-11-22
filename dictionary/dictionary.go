package dictionary

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func Add(file *os.File, key string, value string) {
	_, err := file.WriteString(key + ":" + value + "\n")

	Check(err)

	file.Sync()
}

func Get(file *os.File, word string) (string, error) {
	// Point to the beginning of the file
	_, seekErr := file.Seek(0, 0)
	Check(seekErr)

	// Read each line
	scanner := bufio.NewScanner(file)

	founded := false

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ":")

		// If the given word is a key
		if line[0] == word {
			founded = true
			return line[1], nil
		}

		// If the given word is a value
		if line[1] == word {
			founded = true
			return line[0], nil
		}
	}

	if !founded {
		return "", fmt.Errorf("%s does not exist in this file", word)
	}

	err := scanner.Err()
	Check(err)

	return "", nil
}

func Remove(d map[string]string, file *os.File, word string) (string, error) {
	translation, err := Get(file, word)

	if err != nil {
		return "", err
	}

	// Create a temporary file
	tmpFile, createErr := os.OpenFile("tmp_dictionary.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	Check(createErr)

	defer tmpFile.Close()

	for k, v := range d {
		if k != word && v != word {
			Add(tmpFile, k, v)
		}
	}

	file.Close()

	removeErr := os.Remove("dictionary.txt")
	Check(removeErr)

	tmpFile.Close()

	renameErr := os.Rename("tmp_dictionary.txt", "dictionary.txt")
	Check(renameErr)

	return "<<" + word + ":" + translation + ">> translation deleted", nil
}

// func List(d map[string]string) ([]string, []string) {
// 	var keys []string
// 	var values []string

// 	for key, value := range d {
// 		_, err := Get(d, key)

// 		if err == nil {
// 			keys = append(keys, key)
// 			values = append(values, value)
// 		}
// 	}

// 	return keys, values
// }

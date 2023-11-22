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

func Get(file *os.File, word string) (string, string, error) {
	// Point to the beginning of the file
	_, seekErr := file.Seek(0, 0)
	Check(seekErr)

	// Read each line
	scanner := bufio.NewScanner(file)

	founded := false

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ":")

		// If the given word is a key or a value
		if line[0] == word || line[1] == word {
			founded = true
			return line[0], line[1], nil
		}
	}

	if !founded {
		return "", "", fmt.Errorf("%s does not exist in this file", word)
	}

	err := scanner.Err()
	Check(err)

	return "", "", nil
}

func Remove(file *os.File, word string) (string, error) {
	wordFound, translationFound, err := Get(file, word)

	if err != nil {
		return "", err
	}

	// Create a temporary file
	tmpFile, createErr := os.OpenFile("tmp_dictionary.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	Check(createErr)

	defer tmpFile.Close()

	// Point to the beginning of the file
	_, seekErr := file.Seek(0, 0)
	Check(seekErr)

	// Read each line
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ":")

		// If the given word is a key or a value, we pass
		if line[0] == word || line[1] == word {
			continue
		}

		Add(tmpFile, line[0], line[1])
	}

	err = scanner.Err()
	Check(err)

	file.Close()

	removeErr := os.Remove("dictionary.txt")
	Check(removeErr)

	tmpFile.Close()

	renameErr := os.Rename("tmp_dictionary.txt", "dictionary.txt")
	Check(renameErr)

	return "<<" + wordFound + ":" + translationFound + ">> translation deleted", nil
}

func List(file *os.File) ([]string, []string) {
	var keys []string
	var values []string

	// Point to the beginning of the file
	_, seekErr := file.Seek(0, 0)
	Check(seekErr)

	// Read each line
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ":")

		keys = append(keys, line[0])
		values = append(values, line[1])
	}

	err := scanner.Err()
	Check(err)

	return keys, values
}

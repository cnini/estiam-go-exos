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

func Get(file *os.File, word string) (string, int, error) {
	// Point to the beginning of the file
	_, seekErr := file.Seek(0, 0)
	Check(seekErr)

	// Read each line
	scanner := bufio.NewScanner(file)

	lineNumber := 1
	founded := false

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ":")

		// If the given word is a key
		if line[0] == word {
			founded = true
			return line[1], lineNumber, nil
		}

		// If the given word is a value
		if line[1] == word {
			founded = true
			return line[0], lineNumber, nil
		}

		lineNumber++
	}

	if !founded {
		return "", 0, fmt.Errorf("%s key does not exist in this file", word)
	}

	err := scanner.Err()
	Check(err)

	return "", 0, nil
}

// func Remove(d map[string]string, key string) (string, error) {
// 	_, err := Get(d, key)

// 	if err != nil {
// 		return "", err
// 	} else {
// 		delete(d, key)
// 		return key + " key deleted", nil
// 	}
// }

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

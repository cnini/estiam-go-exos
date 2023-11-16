package dictionary

import (
	"fmt"
	"os"
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

func Get(d map[string]string, key string) (string, error) {
	if value, exists := d[key]; exists {
		return value, nil
	} else {
		return "", fmt.Errorf("Value for %s key does not exist in dictionary", key)
	}
}

func Remove(d map[string]string, key string) (string, error) {
	_, err := Get(d, key)

	if err != nil {
		return "", err
	} else {
		delete(d, key)
		return key + " key deleted", nil
	}
}

func List(d map[string]string) ([]string, []string) {
	var keys []string
	var values []string

	for key, value := range d {
		_, err := Get(d, key)

		if err == nil {
			keys = append(keys, key)
			values = append(values, value)
		}
	}

	return keys, values
}

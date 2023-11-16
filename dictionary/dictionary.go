package dictionary

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

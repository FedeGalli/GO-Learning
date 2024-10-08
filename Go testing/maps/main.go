package maps

import "errors"

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]
	if !ok {
		return "", errors.New("Could not find the word")
	}
	return definition, nil
}

func (d Dictionary) Add(key string, val string) {
	d[key] = val

}

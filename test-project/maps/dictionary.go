package main

import "errors"

type Dictionary map[string]string

// func Search(d Dictionary, word string) string {
// 	return d[word]
// }

func(d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", errors.New("could not find the word you are looking for")
	}
	return d[word], nil
}
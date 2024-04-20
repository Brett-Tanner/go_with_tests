package main

import "errors"

type Dictionary map[string]string

var NotFoundErr = errors.New("could not find the word you were looking for")

func (d Dictionary) Add(key, value string) {
	d[key] = value
}

func (d Dictionary) Search(word string) (string, error) {
	def, ok := d[word]
	if ok {
		return def, nil
	} else {
		return "", NotFoundErr
	}
}

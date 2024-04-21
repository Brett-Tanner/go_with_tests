package main

type (
	Dictionary    map[string]string
	DictionaryErr string
)

func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	NotFoundErr   = DictionaryErr("could not find the word you were looking for")
	WordExistsErr = DictionaryErr("that word can't be added as it already exists")
)

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case NotFoundErr:
		d[word] = definition
		return nil
	case nil:
		return WordExistsErr
	default:
		return err
	}
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case NotFoundErr:
		return NotFoundErr
	case nil:
		delete(d, word)
		return nil
	default:
		return err
	}
}

func (d Dictionary) Search(word string) (string, error) {
	def, ok := d[word]
	if ok {
		return def, nil
	}

	return "", NotFoundErr
}

func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)

	switch err {
	case NotFoundErr:
		return NotFoundErr
	case nil:
		d[word] = newDefinition
		return nil
	default:
		return err
	}
}

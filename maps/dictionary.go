package maps

const (
	ErrNotFound = DictionaryErr("could not find the word")
	ErrWordExists = DictionaryErr("word already exists")
	ErrWordDoesNotExists = DictionaryErr("word does not exists")
)

type DictionaryErr string

type Dictionary map[string]string

// Search returns the value of given key
func (d Dictionary) Search(word string) (string, error) {
	definition, hasValue := d[word]

	if !hasValue {
		return "", ErrNotFound
	}

	return definition, nil
}
// Add adds a word and definition to dictionary
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	
	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExists
	case nil:
		delete(d, word)
	default:
		return err
	}

	return nil
}

// Update changes word's definition updated and returns an error or nil
func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExists
	case nil: 
		d[word] = newDefinition
	default:
		return err
	}

	return nil
}

func (e DictionaryErr) Error() string {
	return string(e)
}
package ShortID

import (
	"errors"
	"sync"
)

type ShortID struct {
	character *CharacterRuneData
	length    int

	locker sync.Mutex
}

//characters example: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
func CreateShortID(characters string, length int) (*ShortID, error) {
	runeData, err := createCharacterRuneData(characters, GetSeed())

	if err != nil {
		return nil, err
	}

	if length <= 0 {
		return nil, errors.New("length must be greater than 0")
	}

	return &ShortID{
		character: runeData,
		length:    length,
	}, nil
}

//characters example: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
func MustCreateShortID(characters string, length int) *ShortID {
	if shortID, err := CreateShortID(characters, length); err == nil {
		return shortID
	} else {
		panic(err)
	}
}

func (x *ShortID) Generate() string {
	x.locker.Lock()
	defer x.locker.Unlock()

	return x.generate()
}

//Update length and generate id
func (x *ShortID) GenerateL(length int) (string, error) {
	x.locker.Lock()
	defer x.locker.Unlock()

	if length <= 0 {
		return "", errors.New("length must be greater than 0")
	}

	x.length = length

	return x.generate(), nil
}

func (x *ShortID) MustGenerateL(length int) string {
	if id, err := x.GenerateL(length); err == nil {
		return id
	} else {
		panic(err)
	}
}

//No mutex
func (x *ShortID) generate() string {
	newCharacters := make([]rune, 0)

	for len(newCharacters) < x.length {
		x.character.shuffle()
		newCharacters = append(newCharacters, x.character.GetCharacters()...)
	}

	return string(newCharacters[:x.length])
}

package ShortID

import (
	"errors"
	"fmt"
	"math/rand"
)

type CharacterRuneData struct {
	characters []rune
	//seed       int64
	rand *rand.Rand
}

func createCharacterRuneData(characters string, seed int64) (*CharacterRuneData, error) {
	runes := []rune(characters)

	if len(runes) < p_MIN_RUNE_LENGTH {
		return nil, fmt.Errorf("characters must be more than %d", p_MIN_RUNE_LENGTH)
	}

	if !checkUniqueRunes(runes) {
		return nil, errors.New("must contain unique characters only")
	}

	return &CharacterRuneData{
		characters: runes,
		rand:       rand.New(rand.NewSource(seed)),
		//seed:       seed,
	}, nil
}

func (data *CharacterRuneData) shuffle() {
	data.rand.Shuffle(len(data.characters), func(i, j int) {
		data.characters[i], data.characters[j] = data.characters[j], data.characters[i]
	})
}

func (data *CharacterRuneData) GetCharacters() []rune {
	return data.characters
}

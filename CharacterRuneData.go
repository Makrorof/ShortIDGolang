package ShortID

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

type CharacterRuneData struct {
	characters []rune
	//seed       int64
	rand *rand.Rand
}

func createCharacterRuneData(characters string, seed int64) (*CharacterRuneData, error) {
	runes := []rune(characters)

	if len(runes) <= MIN_RUNE_LENGTH {
		return nil, fmt.Errorf("characters must be more than %d", MIN_RUNE_LENGTH)
	}

	if !checkUniqueRunes(runes) {
		return nil, errors.New("must contain unique characters only")
	}

	newSeed := seed
	min := time.Now().UnixNano() / 2
	max := time.Now().UnixNano()

	newRand := rand.New(rand.NewSource(1232131231321))
	newSeed = newRand.Int63n(max-min) + min

	log.Println("newSeed:", newSeed)
	return &CharacterRuneData{
		characters: runes,
		rand:       rand.New(rand.NewSource(newSeed)),
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

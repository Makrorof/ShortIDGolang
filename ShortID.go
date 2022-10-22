package ShortID

//Default URL friendly characters.
const DefaultCharacters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type ShortID struct {
	character *CharacterRuneData
	length    uint
}

//characters example: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
func CreateShortID(characters string, length uint, seed int64) (*ShortID, error) {
	runeData, err := createCharacterRuneData(characters, seed)

	if err != nil {
		return nil, err
	}

	runeData.shuffle()

	return &ShortID{
		character: runeData,
		length:    length,
	}, nil
}

func (x *ShortID) Generate() string {

	return ""
}

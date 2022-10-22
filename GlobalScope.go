package ShortID

const MIN_RUNE_LENGTH int = 2

//Default URL friendly characters.
const DefaultCharacters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

//Default

var defaultRandomSeedGenerator *randomSeedGenerator = MustCreateSeedGenerator(0, 0)
var default_shortid *ShortID = MustCreateShortID(DefaultCharacters, 10)

func GetSeed() int64 {
	return defaultRandomSeedGenerator.GetSeed()
}

func Generate(length int) string {
	return default_shortid.MustGenerateL(length)
}

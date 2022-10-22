package ShortID

func checkUniqueRunes(runes []rune) bool {
	for _, runeLeft := range runes {
		runeCount := 0

		for _, runeRight := range runes {
			if runeLeft == runeRight {
				runeCount++

				if runeCount >= 2 {
					return false
				}
			}
		}
	}

	return true
}

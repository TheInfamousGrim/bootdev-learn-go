package main

func getNameCounts(names []string) map[rune]map[string]int {
	namesByFirstRune := make(map[rune]map[string]int)

	for _, name := range names {
		runes := []rune(name)
		firstLetter := runes[0]

		if _, ok := namesByFirstRune[firstLetter]; !ok {
			namesByFirstRune[firstLetter] = make(map[string]int)
		}

		if _, ok := namesByFirstRune[firstLetter][name]; !ok {
			namesByFirstRune[firstLetter][name] = 0
		}
		namesByFirstRune[firstLetter][name]++

	}

	return namesByFirstRune
}

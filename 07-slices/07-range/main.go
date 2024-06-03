package main

func indexOfFirstBadWord(msg []string, badWords []string) int {
	firstIdxBadWord := -1

	for i, word := range msg {
		badWordFound := false

		for _, badWord := range badWords {
			if word == badWord {
				firstIdxBadWord = i
				badWordFound = true
				break
			}
		}

		if badWordFound {
			break
		}
	}

	return firstIdxBadWord
}

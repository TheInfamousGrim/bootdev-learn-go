package main

func maxMessages(thresh int) int {
	currCost := 0
	maxMessages := 0

	for i := 0; ; i++ {
		currCost += 100 + i
		if currCost > thresh {
			return maxMessages
		}

		maxMessages += 1
	}
}

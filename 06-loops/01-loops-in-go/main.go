package main

func bulkSend(numMessages int) float64 {
	if numMessages == 0 {
		return 0.0
	}

	totalCost := 0.00

	for i := 0; i < numMessages; i++ {
		totalCost += 1.00 + (float64(i) / 100)
	}

	return totalCost
}

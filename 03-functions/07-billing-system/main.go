package main

func calculateFinalBill(costPerMessage float64, numMessages int) float64 {
	baseBill := calculateBaseBill(costPerMessage, numMessages)

	discountRate := calculateDiscount(numMessages)

	return baseBill * (1 - discountRate)
}

func calculateDiscount(messagesSent int) float64 {
	if messagesSent > 5000 {
		return 0.2
	} else if messagesSent > 1000 {
		return 0.1
	} else {
		return 0
	}
}

// don't touch below this line

func calculateBaseBill(costPerMessage float64, messagesSent int) float64 {
	return costPerMessage * float64(messagesSent)
}

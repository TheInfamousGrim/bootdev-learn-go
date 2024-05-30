package main

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	messages := [3]string{primary, secondary, tertiary}

	// Costs
	message1Cost := len(primary)
	message2Cost := len(secondary) + message1Cost
	message3Cost := len(tertiary) + message2Cost
	messageCosts := [3]int{message1Cost, message2Cost, message3Cost}

	return messages, messageCosts
}

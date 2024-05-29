package main

import (
	"fmt"
)

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (int, error) {
	costCust, errCust := sendSMS(msgToCustomer)
	if errCust != nil {
		return 0, errCust
	}

	costSpouse, errSpouse := sendSMS(msgToSpouse)
	if errSpouse != nil {
		return 0, errSpouse
	}

	return costCust + costSpouse, nil
}

// don't edit below this line

func sendSMS(message string) (int, error) {
	const maxTextLen = 25
	const costPerChar = 2
	if len(message) > maxTextLen {
		return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * len(message), nil
}

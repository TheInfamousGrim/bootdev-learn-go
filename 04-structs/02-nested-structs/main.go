package main

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	recipientCanReceive := mToSend.recipient.name != "" && mToSend.recipient.number != 0

	senderCanSend := mToSend.sender.name != "" && mToSend.sender.number != 0

	if senderCanSend && recipientCanReceive {
		return true
	} else {
		return false
	}
}

package main

func addEmailsToQueue(emails []string) chan string {
	emailBuff := make(chan string, len(emails))

	for _, email := range emails {
		emailBuff <- email
	}

	return emailBuff
}

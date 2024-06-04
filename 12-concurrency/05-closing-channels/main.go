package main

func countReports(numSentCh chan int) int {
	reportsSent := 0

	for {
		numSent, ok := <-numSentCh
		if !ok {
			break
		}
		reportsSent += numSent
	}

	return reportsSent
}

// don't touch below this line

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
	}
	close(ch)
}

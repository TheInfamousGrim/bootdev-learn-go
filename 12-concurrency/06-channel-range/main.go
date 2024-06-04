package main

func concurrentFib(n int) []int {
	ch := make(chan int)
	var numSlice []int
	go fibonacci(n, ch)

	for num := range ch {
		numSlice = append(numSlice, num)
	}

	return numSlice
}

// don't touch below this line

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

package main

import "testing"

func TestMaxRequest(t *testing.T) {
	var maxRequest uint64
	maxRequest = 10

	totalConcurrentRequest := 15
	reqChan := make(chan bool, totalConcurrentRequest)

	limiter := New(maxRequest)

	expectedSuccessRequest := 10
	expectedFailedRequest := 5

	for i := 0; i < totalConcurrentRequest; i++ {
		go func() {
			res := limiter.Take()
			reqChan <- res
		}()
	}

	for i := 0; i < totalConcurrentRequest; i++ {
		res := <-reqChan

		if res == true {
			expectedSuccessRequest--
		} else {
			expectedFailedRequest--
		}
	}

	if expectedFailedRequest != 0 ||
		expectedSuccessRequest != 0 {
		t.Fail()
	}

}

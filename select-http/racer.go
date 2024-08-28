package main

import (
	"fmt"
	"net/http"
	"time"
)

// func Racer(websitA, websiteB string) (winner string) {
// 	durationA := measureResponseTime(websitA)
// 	durationB := measureResponseTime(websiteB)

// 	if durationA < durationB {
// 		return websitA
// 	}

// 	return websiteB
// }

// func measureResponseTime(url string) time.Duration {
// 	start := time.Now()
// 	http.Get(url)
// 	return time.Since(start)
// }

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(websitA, websiteB string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(websitA):
		return websitA, nil
	case <-ping(websiteB):
		return websiteB, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", websitA, websiteB)
	}
}

func ping(url string) chan struct{} {
	channel := make(chan struct{})

	go func() {
		http.Get(url)
		close(channel)
	}()

	return channel
}

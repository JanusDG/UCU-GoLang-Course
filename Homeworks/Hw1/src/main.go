package main

import (
	"fmt"

	"github.com/JanusDG/UCU-GoLang-Course/Homeworks/Hw1/retry"
)

func main() {
	var testFunc = func() bool {
		return false
	}

	// Showcase default
	retDefault := retry.NewDefaultRetriable()
	fmt.Printf("Number of attempts: %d\n", retDefault.GetMaxAttempts())
	retDefault.Retry(testFunc)

	retDefaultWithTenAttempts := retry.NewDefaultRetriable(10)
	fmt.Printf("Number of attempts: %d\n", retDefaultWithTenAttempts.GetMaxAttempts())
	retDefaultWithTenAttempts.Retry(testFunc)

	// Showcase with delay
	retDelay := retry.NewRetriableWithDelay(1)
	fmt.Printf("Number of attempts: %d\n", retDelay.GetMaxAttempts())
	retDelay.Retry(testFunc)

	retDelayWithTenAttempts := retry.NewRetriableWithDelay(1, 10)
	fmt.Printf("Number of attempts: %d\n", retDelayWithTenAttempts.GetMaxAttempts())
	retDelayWithTenAttempts.Retry(testFunc)
}

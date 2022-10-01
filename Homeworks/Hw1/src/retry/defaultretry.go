package retry

import (
	"fmt"
)

type DefaultRetriable struct {
	maxAttempts uint
}

func NewDefaultRetriable(maxAttempts ...uint) Retriable {
	var maxAttemptsN uint = 5
	if len(maxAttempts) > 0 {
		maxAttemptsN = maxAttempts[0]
	}
	return &DefaultRetriable{maxAttempts: maxAttemptsN}
}

func (dr DefaultRetriable) GetMaxAttempts() uint {
	return dr.maxAttempts
}

func (dr DefaultRetriable) Retry(action func() bool) {
	if dr.maxAttempts == 0 {
		dr.maxAttempts = 5
	}
	var attempts uint = dr.GetMaxAttempts()
	var i uint = 0
	for i < attempts {
		if action() == true {
			break
		}
		i++
		fmt.Printf("after %d attempts, result: %t\n", i, false)
	}
	return
}

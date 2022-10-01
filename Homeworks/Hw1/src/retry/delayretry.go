package retry

import (
	"fmt"
	"time"
)

type RetriableWithDelay struct {
	maxAttempts uint
	DelayInSec  time.Duration
}

func NewRetriableWithDelay(delayInSec uint, maxAttempts ...uint) Retriable {
	var maxAttemptsN uint = 5
	if len(maxAttempts) > 0 {
		maxAttemptsN = maxAttempts[0]
	}
	return &RetriableWithDelay{maxAttempts: maxAttemptsN, DelayInSec: time.Duration(delayInSec)}
}

func (rd RetriableWithDelay) GetMaxAttempts() uint {
	return rd.maxAttempts
}

func (rd RetriableWithDelay) Retry(action func() bool) {
	if rd.maxAttempts == 0 {
		rd.maxAttempts = 5
	}
	var attempts uint = rd.GetMaxAttempts()
	var i uint = 0
	for i < attempts {
		if action() == true {
			break
		}
		i++
		fmt.Printf("after %d attempts, result: %t\n", i, false)
		time.Sleep(time.Second * rd.DelayInSec)
	}
	return
}

package retry

type Retriable interface {
	GetMaxAttempts() uint
	Retry(action func() bool)
}

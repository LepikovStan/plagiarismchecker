package retryutils

func Retry(ctx context.Context, fn func(), delays []time.Duration) {
	for i := 0; i < len(delays); i++ {
		delay := delays[i]

		select {
		case <-ctx.Done():
			return // returning not to leak the goroutine
		default:
			time.Sleep(delay)
			fn()
		}
	}
}

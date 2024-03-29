package retryutils

import (
	"context"
	"time"
)

func Retry(ctx context.Context, fn func(), delays []time.Duration) {
	for i := 0; i < len(delays); i++ {
		delay := delays[i]

		select {
		case <-ctx.Done():
			return
		default:
			time.Sleep(delay)
			fn()
		}
	}
}

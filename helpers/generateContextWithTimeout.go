package helpers

import (
	"context"
	"time"
)

func GenerateContextWithTimeout(timeout int) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)

	return ctx, cancel
}

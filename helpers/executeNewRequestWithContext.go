package helpers

import (
	"context"
	"net/http"
)

func ExecuteNewRequestWithContext(ctx context.Context, route string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", route, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

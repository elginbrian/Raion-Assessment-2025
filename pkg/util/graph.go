package util

import (
	"context"
	"fmt"
)

func ExtractTokenFromContext(ctx context.Context) (string, error) {
	token, ok := ctx.Value("token").(string)
	if !ok || token == "" {
		return "", fmt.Errorf("missing token")
	}
	return token, nil
}

func Coalesce(values ...string) string {
	for _, v := range values {
		if v != "" {
			return v
		}
	}
	return ""
}

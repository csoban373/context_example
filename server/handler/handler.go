package handler

import (
	"context"
	"fmt"
	"net/http"
)

type specialKey struct{}

func Decorator(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ctx = context.WithValue(ctx, specialKey{}, 42)

		f(w, r.WithContext(ctx))
	}
}

func GetValue(ctx context.Context) (int, error) {
	id, ok := ctx.Value(specialKey{}).(int)
	if !ok {
		return 0, fmt.Errorf("could not find value")
	}
	return id, nil
}

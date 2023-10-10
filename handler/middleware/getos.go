package middleware

import (
	"context"
	"errors"
	"net/http"

	ua "github.com/mileusna/useragent"
)

type contextKey struct{}

var OSNamekey = contextKey{}


func ContextWithOS(parent context.Context, r *http.Request) context.Context {
	return context.WithValue(parent, OSNamekey, ua.Parse(r.UserAgent()).OS)
}

func GetOS(ctx context.Context) (string, error) {
	os, ok := ctx.Value(OSNamekey).(string)
	if !ok {
		return "", errors.New("not found")
	}
	return os, nil
}

func GetOSHandler(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := ContextWithOS(r.Context(), r)
		h.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
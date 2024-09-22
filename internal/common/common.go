package common

import (
	"context"
	"gorm.io/gorm"
	"net/http"
)

type AppContext struct {
	Database *gorm.DB
}

var customContextKey string = "CUSTOM_CONTEXT"

func CreateContext(args *AppContext, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		customContext := &AppContext{
			Database: args.Database,
		}
		requestWithCtx := r.WithContext(context.WithValue(r.Context(), customContextKey, customContext))
		next.ServeHTTP(w, requestWithCtx)
	})
}

func GetContext(ctx context.Context) *AppContext {
	customContext, ok := ctx.Value(customContextKey).(*AppContext)
	if !ok {
		return nil
	}
	return customContext
}

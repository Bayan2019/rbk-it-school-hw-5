package middleware

import (
	"context"
	"errors"

	"github.com/Bayan2019/rbk-it-school-hw-5/internal/model"
)

type contextKey string

const userContextKey contextKey = "user"

func withUser(ctx context.Context, user model.UserContext) context.Context {
	return context.WithValue(ctx, userContextKey, user)
}

func UserFromContext(ctx context.Context) (model.UserContext, error) {
	user, ok := ctx.Value(userContextKey).(model.UserContext)
	if !ok {
		return model.UserContext{}, errors.New("user not found in context")
	}

	return user, nil
}

package ctxutil

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"
)

type contextKey string

const (
	UserIdKey = "userId"
)

func WithUserId(ctx context.Context, id int64) context.Context {
	return withCtxValue(ctx, UserIdKey, fmt.Sprintf("%d", id))
}

func GetUserId(ctx context.Context) int64 {
	id, err := strconv.Atoi(getCtxValue(ctx, UserIdKey))
	if err != nil {
		slog.Error("failed to convert user id to int from ctx value", slog.String("err", err.Error()))
	}
	return int64(id)
}

func withCtxValue(ctx context.Context, key string, value string) context.Context {
	if value == "" {
		return ctx
	}
	return context.WithValue(ctx, contextKey(key), value)
}

func getCtxValue(ctx context.Context, key string) string {
	if ret := ctx.Value(contextKey(key)); ret != nil {
		return ret.(string)
	}
	return ""
}

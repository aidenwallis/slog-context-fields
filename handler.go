package slogctx

import (
	"context"
	"log/slog"
)

type contextKey uint8

const (
	argsKey contextKey = iota
)

// GetArgs returns the args currently found in context, or nil
func GetArgs(ctx context.Context) []any {
	if v, ok := ctx.Value(argsKey).([]any); ok {
		return v
	}
	return nil
}

// WithArgs appends slog args to context
func WithArgs(ctx context.Context, args ...any) context.Context {
	return context.WithValue(ctx, argsKey, append(GetArgs(ctx), args...))
}

// Handler wraps a given slog handler, and applies extra fields to the call
type Handler struct {
	slog.Handler
}

var _ slog.Handler = (*Handler)(nil)

// NewHandler creates a new instance of Handler
func NewHandler(handler slog.Handler) *Handler {
	return &Handler{Handler: handler}
}

// Handle appends any fields from context to the record and calls the nested handler
func (h *Handler) Handle(ctx context.Context, r slog.Record) error {
	if v := GetArgs(ctx); len(v) > 0 {
		r.Add(v...)
	}

	return h.Handler.Handle(ctx, r)
}

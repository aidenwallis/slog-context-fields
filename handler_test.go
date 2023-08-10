package slogctx_test

import (
	"context"
	"log/slog"
	"testing"

	"github.com/aidenwallis/slogctx"
)

type testHandler struct {
	slog.Handler
	lastRecord slog.Record
	ch         chan struct{}
}

func (h *testHandler) wait() {
	<-h.ch
}

func (h *testHandler) Handle(ctx context.Context, r slog.Record) error {
	h.lastRecord = r
	h.ch <- struct{}{}
	return nil
}

var _ slog.Handler = (*testHandler)(nil)

func TestHandler(t *testing.T) {
	ctx := context.Background()

	handler := &testHandler{
		Handler: slog.NewTextHandler(nil, nil),
		ch:      make(chan struct{}, 1), // allows the test to pause for the lastRecord to be written
	}
	logger := slog.New(slogctx.NewHandler(handler))

	// default case: should have 0 attrs
	logger.InfoContext(ctx, "test message")
	handler.wait()
	assert(t, handler.lastRecord.NumAttrs(), 0)

	ctx = slogctx.WithArgs(ctx, "abc", 123)

	// should now have 1 attr from the ctx
	logger.InfoContext(ctx, "test message")
	handler.wait()
	assert(t, handler.lastRecord.NumAttrs(), 1)
}

func assert[T comparable](t *testing.T, in, expected T) {
	t.Helper()

	if in != expected {
		t.Fatalf("expected %#v but got %#v", expected, in)
	}
}

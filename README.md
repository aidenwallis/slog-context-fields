# slogctx

[![codecov](https://codecov.io/gh/aidenwallis/slogctx/branch/main/graph/badge.svg?token=CF9slb1Sjp)](https://codecov.io/gh/aidenwallis/slogctx) [![Go Reference](https://pkg.go.dev/badge/github.com/aidenwallis/slogctx.svg)](https://pkg.go.dev/github.com/aidenwallis/slogctx)

Simple [slog](https://pkg.go.dev/log/slog) wrapping handler that lets you pass fields down to a slog call through context.

## Example

```go
package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/aidenwallis/slogctx"
)

var logger = slog.New(slogctx.NewHandler(slog.NewTextHandler(os.Stdout, nil)))

func main() {
	ctx := context.Background()

	logger.InfoContext(ctx, "this message has no extra fields tied to it!")

	ctx = slogctx.WithArgs(ctx, "passed_arg", "yes!")

	deeplyNestedFunc(ctx)
}

func deeplyNestedFunc(ctx context.Context) {
	logger.InfoContext(ctx, "this will attach the fields from above! particularly useful for passing through things like request URLs etc.")
}
```

Note that this now allows you to pass through fields using your context, rather than either deeply pushing a logger down the stack, and is compatible with the global `slog` functions.

package futures

import (
	"context"
	"github.com/ryanbrainard/go-generics-play/monads/results"
	"github.com/ryanbrainard/go-generics-play/testutil"
	"testing"
	"time"
)

var (
	testTaskWant = "all done"

	testShortTask = func(context.Context) results.Result[string] {
		return results.Ok(testTaskWant)
	}

	testLongTask = func(ctx context.Context) results.Result[string] {
		timer := time.NewTimer(time.Minute)
		for {
			select {
			case <-timer.C:
				return results.Ok(testTaskWant)
			case <-ctx.Done():
				return results.New("", ctx.Err())
			}
		}
	}

	tests = []struct {
		name string
		exec Executor[string]
	}{
		{
			name: "chanFuture",
			exec: NewChanFuture[string],
		},
		{
			name: "mxFuture",
			exec: NewMxFuture[string],
		},
		{
			name: "wgFuture",
			exec: NewWgFuture[string],
		},
		{
			name: "onceFuture",
			exec: NewOnceFuture[string],
		},
	}
)

func TestFutureEverything(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			onSuccess := func(r string) string { return r }
			onError := func(err error) string { return err.Error() }

			ctx, cancel := context.WithCancel(context.Background())

			realizedFuture := tt.exec(ctx, testShortTask)
			testutil.AssertEqual(t, testTaskWant, results.Fold(realizedFuture.Get(), onSuccess, onError))
			testutil.AssertEqual(t, testTaskWant, results.Fold(realizedFuture.Get(), onSuccess, onError))
			if f, ok := realizedFuture.(Running); ok {
				testutil.AssertEqual(t, false, f.Running())
			}

			cancelledFuture := tt.exec(ctx, testLongTask)
			if f, ok := cancelledFuture.(Running); ok {
				testutil.AssertEqual(t, true, f.Running())
			}
			cancel()
			testutil.AssertEqual(t, context.Canceled.Error(), results.Fold(cancelledFuture.Get(), onSuccess, onError))
			testutil.AssertEqual(t, context.Canceled.Error(), results.Fold(cancelledFuture.Get(), onSuccess, onError))
			if f, ok := cancelledFuture.(Running); ok {
				testutil.AssertEqual(t, false, f.Running())
			}
		})
	}
}

func BenchmarkFutureEverything(b *testing.B) {
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				ctx, cancel := context.WithCancel(context.Background())

				realizedFuture := tt.exec(ctx, testShortTask)
				realizedFuture.Get()

				cancelledFuture := tt.exec(ctx, testLongTask)
				cancel()
				cancelledFuture.Get()
			}
		})
	}
}

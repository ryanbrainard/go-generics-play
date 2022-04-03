package futures

import (
	"context"
	"github.com/ryanbrainard/go-generics-play/monads/results"
	"github.com/ryanbrainard/go-generics-play/testutil"
	"testing"
	"time"
)

func TestFutureEverything(t *testing.T) {
	testTaskWant := "all done"
	testShortTask := func(context.Context) results.Result[string] {
		time.Sleep(time.Millisecond)
		return results.Ok(testTaskWant)
	}
	testLongTask := func(ctx context.Context) results.Result[string] {
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

	tests := []struct {
		name string
		exec func(func(ctx context.Context) results.Result[string]) Future[string]
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			onSuccess := func(r string) string { return r }
			onError := func(err error) string { return err.Error() }

			realizedFuture := tt.exec(testShortTask)
			testutil.AssertEqual(t, testTaskWant, results.Fold(realizedFuture.Get(), onSuccess, onError))
			testutil.AssertEqual(t, testTaskWant, results.Fold(realizedFuture.Get(), onSuccess, onError))

			cancelledFuture := tt.exec(testLongTask)
			cancelledFuture.Cancel()
			testutil.AssertEqual(t, context.Canceled.Error(), results.Fold(cancelledFuture.Get(), onSuccess, onError))
			testutil.AssertEqual(t, context.Canceled.Error(), results.Fold(cancelledFuture.Get(), onSuccess, onError))
		})
	}
}

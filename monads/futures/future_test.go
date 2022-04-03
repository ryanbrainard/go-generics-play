package futures

import (
	"context"
	"github.com/ryanbrainard/go-generics-play/monads/results"
	"testing"
	"time"
)

func Test_future_Get(t *testing.T) {
	tests := []struct {
		name string
		exec func(func(ctx context.Context) results.Result[string]) Future[string]
		task func(context.Context) results.Result[string]
		want string
	}{
		{
			name: "chanFuture",
			exec: NewChanFuture[string],
			task: func(context.Context) results.Result[string] {
				time.Sleep(time.Millisecond)
				return results.Ok("all done")
			},
			want: "all done",
		},
		{
			name: "mxFuture",
			exec: NewMxFuture[string],
			task: func(context.Context) results.Result[string] {
				time.Sleep(time.Millisecond)
				return results.Ok("all done")
			},
			want: "all done",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := tt.exec(tt.task)
			if got := f.Get().GetOrElse("fail"); got != tt.want {
				t.Errorf("Get #1 = %v, want %v", got, tt.want)
			}
			if got := f.Get().GetOrElse("fail"); got != tt.want {
				t.Errorf("Get #2 = %v, want %v", got, tt.want)
			}
		})
	}
}

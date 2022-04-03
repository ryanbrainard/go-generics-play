package futures

import (
	"context"
	"github.com/ryanbrainard/go-generics-play/monads/results"
	"testing"
	"time"
)

func Test_future_Get(t *testing.T) {
	testTaskWant := "all done"
	testTask := func(context.Context) results.Result[string] {
		time.Sleep(time.Millisecond)
		return results.Ok(testTaskWant)
	}

	tests := []struct {
		name string
		exec func(func(ctx context.Context) results.Result[string]) Future[string]
		task func(context.Context) results.Result[string]
		want string
	}{
		{
			name: "chanFuture",
			exec: NewChanFuture[string],
			task: testTask,
			want: testTaskWant,
		},
		{
			name: "mxFuture",
			exec: NewMxFuture[string],
			task: testTask,
			want: testTaskWant,
		},
		{
			name: "wgFuture",
			exec: NewWgFuture[string],
			task: testTask,
			want: testTaskWant,
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

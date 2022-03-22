package testutil

import (
	"reflect"
	"testing"
)

func AssertEqual[V any](t *testing.T, want, got V) {
	t.Helper()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want='%+v' got='%+v'", want, got)
	}
}

package results

import (
	"errors"
	"github.com/ryanbrainard/go-generics-play/testutil"
	"testing"
)

func FuzzResultString(f *testing.F) {
	testcases := []string{"hello", " ", "123", "ByE''"}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rOk := New(orig, nil)

		rOkStr, rOkErr := rOk.Split()
		testutil.AssertEqual(t, orig, rOkStr)
		testutil.AssertEqual(t, nil, rOkErr)

		testutil.AssertEqual(t, orig, rOk.GetOrElse("x"))

		testutil.AssertEqual(t, orig, FlatMap(rOk, func(res string) Result[string] {
			return rOk
		}).GetOrElse("x"))

		testutil.AssertEqual(t, true, Fold(rOk, func(res string) bool {
			return true
		}, func(res error) bool {
			return false
		}))

		///

		rErr := New(orig, errors.New(orig))

		rErrStr, rErrErr := rErr.Split()
		testutil.AssertEqual(t, "", rErrStr)
		testutil.AssertEqual(t, orig, rErrErr.Error())

		testutil.AssertEqual(t, "x", rErr.GetOrElse("x"))

		testutil.AssertEqual(t, "x", FlatMap(rErr, func(res string) Result[string] {
			return rOk
		}).GetOrElse("x"))

		testutil.AssertEqual(t, false, Fold(rErr, func(res string) bool {
			return true
		}, func(res error) bool {
			return false
		}))
	})
}

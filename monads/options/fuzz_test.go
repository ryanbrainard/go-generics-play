package options

import "testing"

func FuzzOptionString(f *testing.F) {
	testcases := []string{"hello", " ", "123", "ByE''"}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, orig string) {
		some := Some(orig)

		got := some.Get()
		if orig != got {
			t.Errorf("orig: '%v' got: '%v'", orig, got)
		}

		mapped := some.Map(func(s string) string {
			return s
		}).Get()

		if orig != mapped {
			t.Errorf("orig: '%v' mapped: '%v'", orig, got)
		}
	})
}

func FuzzOptionInt(f *testing.F) {
	testcases := []int{1, 2, 0, -1}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, orig int) {
		some := Some(orig)

		got := some.Get()
		if orig != got {
			t.Errorf("orig: '%v' got: '%v'", orig, got)
		}

		mapped := some.Map(func(s int) int {
			return s
		}).Get()

		if orig != mapped {
			t.Errorf("orig: '%v' mapped: '%v'", orig, got)
		}
	})
}

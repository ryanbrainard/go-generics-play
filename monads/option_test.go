package monads

import (
	"testing"
)

///

func TestOption_Get(t *testing.T) {
	testOption_Get(t, []testcaseOption_Get[int]{
		{
			name:   "some int",
			option: Some[int]{42},
			want:   42,
		},
		{
			name:   "none int",
			option: None[int]{},
			want:   0,
		},
	})
	testOption_Get(t, []testcaseOption_Get[string]{
		{
			name:   "some string",
			option: Some[string]{"hello"},
			want:   "hello",
		},
		{
			name:   "none string",
			option: None[string]{},
			want:   "",
		},
	})
}

func testOption_Get[V any](t *testing.T, tests []testcaseOption_Get[V]) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AssertEqual(t, tt.want, tt.option.Get())
		})
	}
}

type testcaseOption_Get[V any] struct {
	name   string
	option Option[V]
	want   V
}

///

func TestOption_Map(t *testing.T) {
	addOneInt := func(v int) int { return v + 1 }
	testOption_Map(t, []testcaseOption_Map[int]{
		{
			name:   "some int fn same return type",
			option: Some[int]{42},
			fn:     addOneInt,
			want:   Some[int]{43},
		},
		{
			name:   "none int fn same return type",
			option: None[int]{},
			fn:     addOneInt,
			want:   None[int]{},
		},
	})

	concatOneString := func(v string) string { return v + "1" }
	testOption_Map(t, []testcaseOption_Map[string]{
		{
			name:   "some string same return type",
			option: Some[string]{"hello"},
			fn:     concatOneString,
			want:   Some[string]{"hello1"},
		},
		{
			name:   "none string  same return type",
			option: None[string]{},
			fn:     concatOneString,
			want:   None[string]{},
		},
	})
}

func testOption_Map[V any](t *testing.T, tests []testcaseOption_Map[V]) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AssertEqual(t, tt.want, tt.option.Map(tt.fn))
		})
	}
}

type testcaseOption_Map[V any] struct {
	name   string
	option Option[V]
	fn     func(V) V
	want   Option[V]
}

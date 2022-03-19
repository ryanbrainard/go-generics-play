package monads

import (
	"testing"
)

///

func TestOption_Get(t *testing.T) {
	testOption_Get(t, []testcaseOption_Get[int]{
		{
			name:   "some int",
			option: some[int]{42},
			want:   42,
		},
		{
			name:   "none int",
			option: none[int]{},
			want:   0,
		},
	})
	testOption_Get(t, []testcaseOption_Get[string]{
		{
			name:   "some string",
			option: some[string]{"hello"},
			want:   "hello",
		},
		{
			name:   "none string",
			option: none[string]{},
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

func TestOption_GetOrElse(t *testing.T) {
	testOption_GetOrElse(t, []testcaseOption_GetOrElse[int]{
		{
			name:   "some int",
			option: some[int]{42},
			orElse: 3,
			want:   42,
		},
		{
			name:   "none int",
			option: none[int]{},
			orElse: 3,
			want:   3,
		},
	})
	testOption_GetOrElse(t, []testcaseOption_GetOrElse[string]{
		{
			name:   "some string",
			option: some[string]{"hello"},
			orElse: "bye",
			want:   "hello",
		},
		{
			name:   "none string",
			option: none[string]{},
			orElse: "bye",
			want:   "bye",
		},
	})
}

func testOption_GetOrElse[V any](t *testing.T, tests []testcaseOption_GetOrElse[V]) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AssertEqual(t, tt.want, tt.option.GetOrElse(tt.orElse))
		})
	}
}

type testcaseOption_GetOrElse[V any] struct {
	name   string
	option Option[V]
	orElse V
	want   V
}

///

func TestOption_Map(t *testing.T) {
	addOneInt := func(v int) int { return v + 1 }
	testOption_Map(t, []testcaseOption_Map[int]{
		{
			name:   "some int fn same return type",
			option: some[int]{42},
			fn:     addOneInt,
			want:   some[int]{43},
		},
		{
			name:   "none int fn same return type",
			option: none[int]{},
			fn:     addOneInt,
			want:   none[int]{},
		},
	})

	concatOneString := func(v string) string { return v + "1" }
	testOption_Map(t, []testcaseOption_Map[string]{
		{
			name:   "some string same return type",
			option: some[string]{"hello"},
			fn:     concatOneString,
			want:   some[string]{"hello1"},
		},
		{
			name:   "none string  same return type",
			option: none[string]{},
			fn:     concatOneString,
			want:   none[string]{},
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

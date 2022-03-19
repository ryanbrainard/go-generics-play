package monads_test

import (
	"fmt"
	"go.ryanbrainard.com/go-generics-play/monads"
)

func ExampleSome_Int_Chain() {
	fmt.Println(monads.SomeOf(42).Map(func(v int) int { return v + 1 }).Get())
	// Output: 43
}

func ExampleSome_String_Chain() {
	fmt.Println(monads.SomeOf("hello").Map(func(v string) string { return v + "!" }).Get())
	// Output: hello!
}

func ExampleNone_Int_Chain() {
	fmt.Println(monads.NoneOf[int]().Map(func(v int) int { return v + 1 }).Get())
	// Output: 0
}

func ExampleNone_String_Chain() {
	fmt.Println(monads.NoneOf[string]().Map(func(v string) string { return v + "!" }).Get())
	// Output:
}

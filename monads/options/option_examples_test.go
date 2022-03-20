package options_test

import (
	"fmt"
	"go.ryanbrainard.com/go-generics-play/monads/options"
	"strconv"
)

var (
	someInt      = options.Some(42)
	someStr      = options.Some("hello")
	noneInt      = options.None[int]()
	noneStr      = options.None[string]()
	otherStr     = "hola"
	someOtherStr = options.Some(otherStr)
	addOne       = func(v int) int { return v + 1 }
	addBang      = func(v string) string { return v + "!" }
)

func ExampleSome_Fold() {
	fmt.Println(someStr.Fold(otherStr)(addBang))

	// Output: hello!
}

func ExampleNone_Fold() {
	fmt.Println(noneStr.Fold(otherStr)(addBang))

	// Output: hola
}

func ExampleSome_Get() {
	fmt.Println(someInt.Get())

	// Output: 42
}

func ExampleNone_Get() {
	fmt.Println(noneInt.Get())

	// Output: 0
}

func ExampleSome_GetOrElse() {
	fmt.Println(someStr.GetOrElse(otherStr))

	// Output: hello
}

func ExampleNone_GetOrElse() {
	fmt.Println(noneStr.GetOrElse(otherStr))

	// Output: hola
}

func ExampleSome_Map() {
	fmt.Println(someStr.Map(addBang))

	// Output: {hello!}
}

func ExampleNone_Map() {
	fmt.Println(noneStr.Map(addBang))

	// Output: {}
}

func ExampleSome_Else() {
	fmt.Println(someStr.OrElse(someOtherStr))

	// Output: {hello}
}

func ExampleNone_OrElse() {
	fmt.Println(noneStr.OrElse(someOtherStr))

	// Output: {hola}
}

func ExampleMap() {
	opt := options.Some(42)
	fn := func(v int) string { return strconv.Itoa(v) + "!" }
	fmt.Println(options.Map[int, string](opt, fn).Get())

	// Output: 42!
}

func ExampleFold() {
	fn := func(v int) string { return strconv.Itoa(v) + "!" }
	ifNone := "not found"

	some := options.Some(42)
	fmt.Println(options.Fold[int, string](some, ifNone)(fn))

	none := options.None[int]()
	fmt.Println(options.Fold[int, string](none, ifNone)(fn))

	// Output:
	// 42!
	// not found
}

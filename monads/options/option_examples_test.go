package options_test

import (
	"fmt"
	"go.ryanbrainard.com/go-generics-play/monads/options"
	"strconv"
)

var (
	someInt      = options.SomeOf(42)
	someStr      = options.SomeOf("hello")
	noneInt      = options.NoneOf[int]()
	noneStr      = options.NoneOf[string]()
	otherStr     = "hola"
	someOtherStr = options.SomeOf(otherStr)
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
	opt := options.SomeOf(42)
	fn := func(v int) string { return strconv.Itoa(v) + "!" }
	fmt.Println(options.Map[int, string](opt, fn).Get())

	// Output: 42!
}

func ExampleFold() {
	fn := func(v int) string { return strconv.Itoa(v) + "!" }
	ifNone := "not found"

	some := options.SomeOf(42)
	fmt.Println(options.Fold[int, string](some, ifNone)(fn))

	none := options.NoneOf[int]()
	fmt.Println(options.Fold[int, string](none, ifNone)(fn))

	// Output:
	// 42!
	// not found
}

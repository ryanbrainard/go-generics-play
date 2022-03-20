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

func ExampleSomeOf() {
	fmt.Println("A:", someInt)
	fmt.Println("B:", someStr)
	fmt.Println("C:", someInt.Get())
	fmt.Println("D:", someInt.Map(addOne))
	fmt.Println("E:", someStr.Map(addBang))
	fmt.Println("F:", someStr.GetOrElse(otherStr))
	fmt.Println("G:", someStr.OrElse(someOtherStr))
	fmt.Println("H:", someStr.Map(addBang).OrElse(someOtherStr).Get())

	// Output:
	// A: {42}
	// B: {hello}
	// C: 42
	// D: {43}
	// E: {hello!}
	// F: hello
	// G: {hello}
	// H: hello!
}

func ExampleNoneOf() {
	fmt.Println("A:", noneInt)
	fmt.Println("B:", noneStr)
	fmt.Println("C:", noneInt.Get())
	fmt.Println("D:", noneInt.Map(addOne))
	fmt.Println("E:", noneStr.Map(addBang))
	fmt.Println("F:", noneStr.GetOrElse(otherStr))
	fmt.Println("G:", noneStr.OrElse(someOtherStr))
	fmt.Println("H:", noneStr.Map(addBang).OrElse(someOtherStr).Get())

	// Output:
	// A: {}
	// B: {}
	// C: 0
	// D: {}
	// E: {}
	// F: hola
	// G: {hola}
	// H: hola
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

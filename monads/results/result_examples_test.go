package results_test

import (
	"errors"
	"fmt"
	"go.ryanbrainard.com/go-generics-play/monads/results"
)

func doSomethingClassic(raise bool) (string, error) {
	if raise {
		return "", errors.New("boom")
	}
	return "OK", nil
}

func doSomethingGeneric(raise bool) results.Result[string] {
	if raise {
		return results.Err[string](errors.New("boom"))
	}
	return results.Ok("OK")
}

func ExampleNew() {
	fmt.Println(results.New(doSomethingClassic(false)).GetOrElse("failed"))
	fmt.Println(results.New(doSomethingClassic(true)).GetOrElse("failed"))

	// Output:
	// OK
	// failed
}

func ExampleOk() {
	fmt.Println(results.Ok("hello").GetOrElse("failed"))

	// Output:
	// hello
}

func ExampleErr() {
	fmt.Println(results.Err[string](errors.New("boom")).GetOrElse("failed"))

	// Output:
	// failed
}

func ExampleMap() {
	addParens := func(r string) string { return "(" + r + ")" }

	fmt.Println(results.Map(doSomethingGeneric(false), addParens).GetOrElse("failed"))
	fmt.Println(results.Map(doSomethingGeneric(true), addParens).GetOrElse("failed"))

	// Output:
	// (OK)
	// failed
}

func ExampleMapError() {
	addParens := func(r error) error { return errors.New("(" + r.Error() + ")") }

	fmt.Println(results.MapError[string](doSomethingGeneric(false), addParens).GetOrElse("failed"))
	fmt.Println(results.MapError[string](doSomethingGeneric(true), addParens).GetOrElse("failed"))

	// Output:
	// OK
	// failed
}

func ExampleFold() {
	onSuccess := func(r string) string { return "Success: " + r }
	onError := func(l error) string { return "Got an error: " + l.Error() }

	fmt.Println(results.Fold(doSomethingGeneric(false), onSuccess, onError))
	fmt.Println(results.Fold(doSomethingGeneric(true), onSuccess, onError))

	// Output:
	// Success: OK
	// Got an error: boom
}

func ExampleResult_Split() {
	fmt.Println(doSomethingGeneric(false).Split())
	fmt.Println(doSomethingGeneric(true).Split())

	// Output:
	// OK <nil>
	//  boom
}

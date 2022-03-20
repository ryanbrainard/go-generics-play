package eithers_test

import (
	"errors"
	"fmt"
	"go.ryanbrainard.com/go-generics-play/monads/eithers"
)

var (
	rightStr      = eithers.Right[error, string]("hello")
	leftErr       = eithers.Left[error, string](errors.New("boom"))
	otherStr      = "hola"
	rightOtherStr = eithers.Right[error, string](otherStr)
	addOne        = func(v int) int { return v + 1 }
	addBang       = func(v string) string { return v + "!" }
)

func ExampleRight_GetOrElse() {
	fmt.Println(rightStr.GetOrElse(otherStr))

	// Output: hello
}

func ExampleLeft_GetOrElse() {
	fmt.Println(leftErr.GetOrElse(otherStr))

	// Output: hola
}

func ExampleRight_OrElse() {
	fmt.Println(rightStr.OrElse(rightOtherStr))

	// Output: {hello}
}

func ExampleLeft_OrElse() {
	fmt.Println(leftErr.OrElse(rightOtherStr))

	// Output: {hola}
}

func ExampleRight_Map() {
	fmt.Println(rightStr.Map(addBang))

	// Output: {hello!}
}

//func ExampleLeft_Map() {
//	fmt.Println(leftErr.Map(addBang))
//
//	// Output: {hola}
//}

func ExampleFold() {
	ifLeft := func(l error) string { return l.Error() + "!" }
	ifRight := func(r string) string { return r + "!" }

	fmt.Println(eithers.Fold[error, string](rightStr, ifLeft, ifRight))
	fmt.Println(eithers.Fold[error, string](leftErr, ifLeft, ifRight))

	// Output:
	// hello!
	// boom!
}

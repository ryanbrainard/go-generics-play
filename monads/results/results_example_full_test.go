package results_test

import (
	"errors"
	"fmt"
	"go.ryanbrainard.com/go-generics-play/monads/results"
	"strconv"
)

// User model
type User struct {
	Id   int64
	Name string
}

// Fake user database
var userDB = map[int64]User{
	1: {
		Id:   1,
		Name: "Fred",
	},
	2: {
		Id:   2,
		Name: "Jane",
	},
}

// create a new `Result` from `(int64, error)`
func parseUserId(userId string) results.Result[int64] {
	return results.New(strconv.ParseInt(userId, 10, 64))
}

// create a new `Result` as either an `Ok` or `Err`
func loadUser(userId int64) results.Result[User] {
	if user, ok := userDB[userId]; ok {
		return results.Ok(user)
	} else {
		return results.Err[User](errors.New("unknown user"))
	}
}

// convert a User to a string.
// note, this returns the same type as onError
func onSuccess(user User) string {
	return "Hi, " + user.Name
}

// convert an error to a string
// note, this returns the same type as onSuccess
func onError(err error) string {
	return "Sorry, we hit an error: " + err.Error()
}

// Example of a successful method sequence
func Example_all_successful() {
	// parse user id
	// this will return an `Ok` result
	parseUserIdResult := parseUserId("2")

	// instead of first doing error checking, just FlatMap with next function
	loadUserResult := results.FlatMap(parseUserIdResult, loadUser)

	// fold the final result with the success/error handlers to make a string.
	// in this case, it calls the onSuccess handler.
	fmt.Println(results.Fold(loadUserResult, onSuccess, onError))

	// output:
	// Hi, Jane
}

// Example of a method sequence where the first step fails
func Example_first_step_fails() {
	// parse user id
	// this will return an `Err` result
	parseUserIdResult := parseUserId("not a number")

	// instead of first doing error checking, just FlatMap with next function.
	// even though the first step actually failed, we can safely ignore this.
	// internally `FlatMap` won't actually call `loadUser` because of the previous error.
	loadUserResult := results.FlatMap(parseUserIdResult, loadUser)

	// fold the final result with the success/error handlers to make a string.
	// in this case it calls the `onError` handler and outputs the first error.
	fmt.Println(results.Fold(loadUserResult, onSuccess, onError))

	// output:
	// Sorry, we hit an error: strconv.ParseInt: parsing "not a number": invalid syntax
}

// Example of a method sequence where the second step fails
func Example_second_step_fails() {
	// parse user id
	// this returns an `Ok` result
	parseUserIdResult := parseUserId("3")

	// instead of first doing error checking, just FlatMap with next function.
	// the first result was `Ok`, so `loadUser` gets called, but it will fail
	// because the user doesn't exist.
	loadUserResult := results.FlatMap(parseUserIdResult, loadUser)

	// fold the final result with the success/error handlers to make a string.
	// in this case it calls the onError handler and outputs the second error.
	fmt.Println(results.Fold(loadUserResult, onSuccess, onError))

	// output:
	// Sorry, we hit an error: unknown user
}

// Example of a successful method sequence using `Split()`
// to show the equvilent using classic Go error handling
func Example_all_successful_split() {
	// parse user id
	// this will return an `Ok` result
	userId, err := parseUserId("2").Split()
	if err != nil {
		fmt.Println(onError(err))
		return
	}

	// instead of first doing error checking, just FlatMap with next function
	user, err := loadUser(userId).Split()
	if err != nil {
		fmt.Println(onError(err))
		return
	}

	// fold the final result with the success/error handlers to make a string.
	// in this case, it calls the onSuccess handler.
	fmt.Println(onSuccess(user))

	// output:
	// Hi, Jane
}

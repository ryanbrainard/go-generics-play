package results_test

import (
	"errors"
	"fmt"
	"go.ryanbrainard.com/go-generics-play/monads/results"
)

func CallServer(name string, raise bool) results.Result[string] {
	fmt.Println("server: " + name)
	if raise {
		return results.Err[string](errors.New("failed to call server " + name))
	}
	return results.Ok("server " + name + " called successfully")
}

func Example_full() {
	onSuccess := func(msg string) bool {
		fmt.Println("success: " + msg)
		return true
	}

	onError := func(err error) bool {
		fmt.Println("error: " + err.Error())
		return false
	}

	// Calls Red and then Blue
	successfulChain := CallServer("Red", false).
		FlatMap(func(redMsg string) results.Result[string] {
			return CallServer("Blue", false)
		})
	fmt.Println(results.Fold(successfulChain, onSuccess, onError))

	fmt.Println("---")

	// Calls Red, fails, and does not call Blue
	unsuccessfulChain := CallServer("Red", true).
		FlatMap(func(redMsg string) results.Result[string] {
			return CallServer("Blue", false)
		})
	fmt.Println(results.Fold(unsuccessfulChain, onSuccess, onError))

	// server: Red
	// server: Blue
	// success: server Blue called successfully
	// true
	// ---
	// server: Red
	// error: failed to call server Red
	// false
}

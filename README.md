# go-generics-play

Just playing around with generics in Go 1.18.

## Contents

- Monads
  - Option
    - some
    - none
  - Either
    - left
    - right
  - Result
    - (Ok)
    - (Err)

## Examples

### `Result[R any]`

Given a function that returns a `Result[R any]` and success/error handlers that unify to the same type:

```go
func CallServer(name string, raise bool) results.Result[string] {
	fmt.Println("server: " + name)
	if raise {
		return results.Err[string](errors.New("failed to call server " + name))
	}
	return results.Ok("server " + name + " called successfully")
}

func onSuccess(msg string) bool {
    fmt.Println("success: " + msg)
    return true
}

func onError(err error) bool {
    fmt.Println("error: " + err.Error())
    return false
}
```

A successful chain calls Red and then Blue:

```go
successfulChain := CallServer("Red", false).
    FlatMap(func(redMsg string) results.Result[string] {
        return CallServer("Blue", false)
    })
fmt.Println(results.Fold(successfulChain, onSuccess, onError))

// output:
// server: Red
// server: Blue
// success: server Blue called successfully
```

An unsuccessful chain calls Red, errors, so does not call Blue:

```go
unsuccessfulChain := CallServer("Red", true).
    FlatMap(func(redMsg string) results.Result[string] {
        return CallServer("Blue", false)
    })
fmt.Println(results.Fold(unsuccessfulChain, onSuccess, onError))

// output:
// server: Red
// error: failed to call server Red
// false
```
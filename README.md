# go-generics-play

Just playing around with generics in Go 1.18.

## Contents

- Monads
  - [Option](https://github.com/ryanbrainard/go-generics-play/tree/main/monads/options)
    - some
    - none
  - [Either](https://github.com/ryanbrainard/go-generics-play/tree/main/monads/eithers)
    - left
    - right
  - [Result](https://github.com/ryanbrainard/go-generics-play/tree/main/monads/results)
    - (Ok)
    - (Err)

## Examples

### `Result[R any]`

`Result[R any]` is a monad that is either `R` or an `error`. 
It can  be used to replace to the common `(R, error)` multiple return type in Go.
Instead of having to do the common `if err != nil { ... }` dance for every call in a sequence,
they can be chained together with `FlatMap`, and then be success/error handled one time at the end with `Fold`.

Given a function that returns a `Result[R any]` and success/error handlers that unify to the same type:

```go
parseUserIdResult := parseUserId("2")                          // returns `Result[int64]` (i.e. `int64` or `error`).
loadUserResult := results.FlatMap(parseUserIdResult, loadUser) // returns `Result[User]`  (i.e. `User` or `error`). skips if previous function errored.
return results.Fold(loadUserResult, onSuccess, onError)        // returns `string` for both success and error cases.
```

See [full example](https://github.com/ryanbrainard/go-generics-play/blob/main/monads/results/results_example_full_test.go) for details and additional examples.
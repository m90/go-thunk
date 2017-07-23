# go-thunk

[![Build Status](https://travis-ci.org/m90/go-thunk.svg?branch=master)](https://travis-ci.org/m90/go-thunk)
[![godoc](https://godoc.org/github.com/m90/go-thunk?status.svg)](http://godoc.org/github.com/m90/go-thunk)

> wrap a thunk with defer / recover

Package `thunk` decorates the a passed thunk or an `http.Handler` with a [defer / recover block](https://blog.golang.org/defer-panic-and-recover)

### Installation using go get

```sh
$ go get github.com/m90/go-thunk
```

### Usage

Wrap a thunk using `RunSafely(func())`, discarding any error:

```go
RunSafely(func() {
	result := weirdpackage.DangerousThing()
	// ....
})
```

Pass an error callback a thunk using `RunSafelyWith(func(), func(error))`:

```go
RunSafelyWith(func() {
	result := weirdpackage.DangerousThing()
	// ....
}, func(err error) {
	fmt.Printf("encountered error executing thunk: %v\n", err)
})
```

which (roughly) equals:

```go
defer func() {
	if err := recover(); err != nil {
		fmt.Printf("encountered error executing thunk: %v\n", err)
	}
}()
result := weirdpackage.DangerousThing()
```

### HTTP Middleware

To decorate your `http.Handler`, use `HandleSafelyWith(func(error))` or `HandleSafely()`:

```go
handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	res := shadypkg.DoThings()
	w.Write([]byte(res))
})
middleware := HandleSafelyWith(func(err error) {
	fmt.Printf("Encountered error: %v", err)
})
http.ListenAndServe("0.0.0.0:8080", middleware(handler))
```

In case of an error, the handler will respond with a 500 error code.


### License
MIT © [Frederik Ring](http://www.frederikring.com)

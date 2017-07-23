# go-thunk

[![Build Status](https://travis-ci.org/m90/go-thunk.svg?branch=master)](https://travis-ci.org/m90/go-thunk)
[![godoc](https://godoc.org/github.com/m90/go-thunk?status.svg)](http://godoc.org/github.com/m90/go-thunk)

> wrap a thunk with defer / recover

Package `thunk` decorates the passed thunk with a [defer / recover block](https://blog.golang.org/defer-panic-and-recover)

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

### License
MIT © [Frederik Ring](http://www.frederikring.com)

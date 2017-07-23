package thunk

import (
	"fmt"
)

// RunSafelyWith wraps the passed thunk with defer/recover and calls
// the callback with possible errors
func RunSafelyWith(thunk func(), callback func(error)) {
	defer func() {
		if ex := recover(); ex != nil {
			if err, ok := ex.(error); ok {
				callback(err)
			} else {
				callback(fmt.Errorf("%v", ex))
			}
		}
	}()
	thunk()
}

// RunSafely wraps the passed thunk with defer/recover
func RunSafely(thunk func()) {
	RunSafelyWith(thunk, func(_ error) {})
}

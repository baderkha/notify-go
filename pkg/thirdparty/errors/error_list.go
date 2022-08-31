package errors

import "sync"

// ErrorList is used to chain a list of potential errors and is thread-safe.
type ErrorList struct {
	mux  sync.RWMutex
	errs []error
}

// Error will return the string-form of the errors.
// Implements the error interface.
func (e *ErrorList) Error() string {
	if e == nil {
		return ""
	}

	e.mux.RLock()
	defer e.mux.RUnlock()

	if len(e.errs) == 0 {
		return ""
	}

	if len(e.errs) == 1 {
		return e.errs[0].Error()
	}

	b := []byte("the following errors occured:\n")
	for _, err := range e.errs {
		b = append(b, err.Error()...)
		b = append(b, '\n')
	}

	return string(b)
}

// Err will return an error if the errorlist is not empty.
// If there's only 1 error, it will be directly returned.
// If the errorlist is empty - nil is returned.
func (e *ErrorList) Err() (err error) {
	if e == nil {
		return
	}
	e.mux.RLock()
	switch len(e.errs) {
	case 0: // do nothing
	case 1:
		err = e.errs[0]
	default:
		err = e
	}
	e.mux.RUnlock()
	return
}

// Push will push an error to the errorlist
// If err is a errorlist, it will be merged.
// If the errorlist is nil, it will be created.
func (e *ErrorList) Push(err error) {
	if err == nil {
		return
	}

	e.mux.Lock()
	defer e.mux.Unlock()

	switch v := err.(type) {
	case *ErrorList:
		v.ForEach(func(err error) {
			e.errs = append(e.errs, err)
		})

	default:
		e.errs = append(e.errs, err)
	}
}

// ForEach will iterate through all of the errors within the error list.
func (e *ErrorList) ForEach(fn func(error)) {
	if e == nil {
		return
	}

	e.mux.RLock()
	for _, err := range e.errs {
		fn(err)
	}
	e.mux.RUnlock()
}

// Copy will copy the items from the inbound error list to the source
func (e *ErrorList) Copy(in *ErrorList) {
	if in == nil {
		return
	}

	e.mux.Lock()
	defer e.mux.Unlock()

	in.mux.RLock()
	defer in.mux.RUnlock()

	e.errs = append(e.errs, in.errs...)
}

// Len will return the length of the inner errors list.
func (e *ErrorList) Len() (n int) {
	if e == nil {
		return
	}

	e.mux.RLock()
	n = len(e.errs)
	e.mux.RUnlock()
	return
}

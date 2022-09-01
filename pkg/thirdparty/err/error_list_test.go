package err

import (
	"errors"
	"testing"
)

func TestList(t *testing.T) {
	var errs List
	errs.Push(errors.New("hello world"))
	if errs.Err() == nil {
		t.Fatal("error is nil when it should not be")
	}

	if errs.Len() != 1 {
		t.Fatal("invalid List length")
	}

	errs.ForEach(func(err error) {
		if err.Error() != "hello world" {
			t.Fatal("invalid error")
		}
	})
}

func TestNilList(t *testing.T) {
	var errs List
	errs.Push(errors.New("hello world"))
	if errs.Err() == nil {
		t.Fatal("error is nil when it should not be")
	}

	if errs.Len() != 1 {
		t.Fatal("invalid List length")
	}

	errs.ForEach(func(err error) {
		if err.Error() != "hello world" {
			t.Fatal("invalid error")
		}
	})

	var errs2 *List
	errs.Push(errs2)
}

package serializer

import "io"

// Writer : serializer generic writer
type Writer[T any] interface {
	Write(data *T, w io.Writer) error
}

// Reader : serializer generic reader
type Reader[T any] interface {
	Read(r io.Reader) (*T, error)
}

// Iface : serializer interface
type Iface[T any] interface {
	Reader[T]
	Writer[T]
}

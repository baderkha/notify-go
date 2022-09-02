package serializer

import (
	"encoding/json"
	"io"
	"strings"

	"github.com/pkg/errors"
)

var _ Iface[any] = &JSON[any]{}

const errPrefixJSONSer = "json serializer error "

// Serializes a configuration as json
type JSON[T any] struct {
}

func (j *JSON[T]) Write(data *T, w io.Writer) error {
	b, err := json.Marshal(data)
	if err != nil {

		return errors.Wrap(err, errPrefixJSONSer)
	}
	_, err = w.Write(b)
	if err != nil {
		return errors.Wrap(err, errPrefixJSONSer)
	}
	return nil
}

func (j *JSON[T]) Read(r io.Reader) (*T, error) {
	var res T
	buf := new(strings.Builder)
	_, err := io.Copy(buf, r)
	if err != nil {
		return nil, errors.Wrap(err, errPrefixJSONSer)
	}
	err = json.Unmarshal([]byte(buf.String()), &res)
	if err != nil {
		return nil, errors.Wrap(err, errPrefixJSONSer)
	}
	return &res, nil
}

package notify

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
	"gopkg.in/resty.v1"
)

// WebhookSender : Generic message sender using webhooks
type WebhookSender[T interface{ WithBody(b []byte) T }] struct {
	errPrefix string
	r         *resty.Client
}

// Send send to someone not in your default configs
func (s *WebhookSender[T]) Send(reciever string, bodyContent []byte) error {
	_, err := url.ParseRequestURI(reciever)
	if err != nil {
		return err
	}
	var bdy T
	bdy = bdy.WithBody(bodyContent)
	res, err := s.r.
		R().
		SetBody(bdy).
		Post(reciever)
	return handleRestyResponse(res, err, s.errPrefix)
}

func handleRestyResponse(res *resty.Response, errResty error, wrapErrPrefix string) (err error) {
	if errResty != nil {
		return errors.Wrap(err, wrapErrPrefix)
	} else if res.StatusCode() >= http.StatusBadRequest {
		return fmt.Errorf("%s : %s", wrapErrPrefix, string(res.Body()))
	}
	return nil
}

func NewWebhookSender[T interface{ WithBody(b []byte) T }](errPrefix string) (w *WebhookSender[T]) {
	return &WebhookSender[T]{
		errPrefix: errPrefix,
		r:         resty.New().SetDebug(isDebug),
	}
}

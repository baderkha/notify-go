package notify

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	"gopkg.in/resty.v1"
)

// WebhookConfig : webhook sender configuration
type WebhookConfig struct {
	DefaultChannelWebhook string
}

// WebhookSender : Generic message sender using webhooks
type WebhookSender[T interface{ WithBody(b []byte) T }] struct {
	errPrefix         string
	defaultWebhookURL string
	r                 *resty.Client
}

// SendToDefaultReciever : send to someone in your default configs
func (s *WebhookSender[T]) SendToDefaultReciever(bodyContent []byte) error {
	return s.SendToReciever(s.defaultWebhookURL, bodyContent)
}

func (s *WebhookSender[T]) GetDefaultReciever() string {
	return s.defaultWebhookURL
}

// SendToReciever send to someone not in your default configs
func (s *WebhookSender[T]) SendToReciever(reciever string, bodyContent []byte) error {
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

func NewWebhookSender[T interface{ WithBody(b []byte) T }](cfg *WebhookConfig, errPrefix string) (w *WebhookSender[T]) {
	return &WebhookSender[T]{
		errPrefix:         errPrefix,
		defaultWebhookURL: cfg.DefaultChannelWebhook,
		r:                 resty.New().SetDebug(isDebug),
	}
}

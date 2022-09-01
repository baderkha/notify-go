package notify

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"gopkg.in/resty.v1"
)

type TestBody struct {
}

func (t TestBody) WithBody(b []byte) TestBody {
	return t
}

func Test_Notify_NewWebhookSender(t *testing.T) {
	ws := NewWebhookSender[TestBody](
		&WebhookConfig{
			DefaultChannelWebhook: "some channel",
		},
		"my cool err prefix",
	)

	assert.Equal(t, ws.defaultWebhookURL, "some channel")
	assert.Equal(t, ws.errPrefix, "my cool err prefix")
	assert.NotNil(t, ws.r)
}

func Test_Notify_WebhookSender_SendToReciever(t *testing.T) {
	prefix := "prefix"
	expectedURL := "https://discord.gets.free.stuff"
	restyClient := resty.DefaultClient
	httpmock.ActivateNonDefault(resty.DefaultClient.GetClient())
	//defer httpmock.DeactivateAndReset()
	ws := WebhookSender[TestBody]{
		errPrefix:         prefix,
		defaultWebhookURL: expectedURL,
		r:                 restyClient,
	}

	// for bad status code i expect an error and return the body of the bad response
	{
		for i := http.StatusBadRequest; i <= http.StatusNetworkAuthenticationRequired; i++ {

			httpmock.RegisterResponder("POST", expectedURL,
				httpmock.NewStringResponder(i, `ERROR FROM HTTP`))
			err := ws.SendToReciever(expectedURL, []byte("hello")) // 400 response code == err
			assert.Error(t, err)
			assert.Contains(t, err.Error(), "ERROR FROM HTTP")
		}

	}
	// for good status codes i expect no error return
	{
		for i := http.StatusOK; i < http.StatusBadRequest; i++ {

			httpmock.RegisterResponder("POST", expectedURL,
				httpmock.NewStringResponder(i, `OK`))
			err := ws.SendToReciever(expectedURL, []byte("hello")) // 400 response code == err
			assert.Nil(t, err)
		}
	}

}

func Test_Notify_WebhookSender_SendToDefaultReciever(t *testing.T) {
	prefix := "prefix"
	expectedURL := "https://discord.gets.free.stuff"
	restyClient := resty.DefaultClient
	httpmock.ActivateNonDefault(resty.DefaultClient.GetClient())
	//defer httpmock.DeactivateAndReset()
	ws := WebhookSender[TestBody]{
		errPrefix:         prefix,
		defaultWebhookURL: expectedURL,
		r:                 restyClient,
	}

	// for bad status code i expect an error and return the body of the bad response
	{
		for i := http.StatusBadRequest; i <= http.StatusNetworkAuthenticationRequired; i++ {

			httpmock.RegisterResponder("POST", expectedURL,
				httpmock.NewStringResponder(i, `ERROR FROM HTTP`))
			err := ws.SendToDefaultReciever([]byte("hello")) // 400 response code == err
			assert.Error(t, err)
			assert.Contains(t, err.Error(), "ERROR FROM HTTP")
		}

	}
	// for good status codes i expect no error return
	{
		for i := http.StatusOK; i < http.StatusBadRequest; i++ {

			httpmock.RegisterResponder("POST", expectedURL,
				httpmock.NewStringResponder(i, `OK`))
			err := ws.SendToDefaultReciever([]byte("hello")) // 400 response code == err
			assert.Nil(t, err)
		}
	}

}

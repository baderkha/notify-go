package notify

import "github.com/stretchr/testify/mock"

var _ MessageSender = &SenderMock{}

// no test this is a mock file

// SenderMock : mock implementation you can use in your testing
type SenderMock struct {
	mock.Mock
}

// SendToDefaultReciever : send to someone in your default configs
func (s *SenderMock) SendToDefaultReciever(bodyContent []byte) error {
	args := s.Called(bodyContent)
	return args.Error(0)
}

// SendToReciever send to someone not in your default configs
func (s *SenderMock) SendToReciever(reciever string, bodyContent []byte) error {
	args := s.Called(reciever, bodyContent)
	return args.Error(0)
}

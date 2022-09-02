package notify

import "github.com/stretchr/testify/mock"

var _ MessageSender = &SenderMock{}

// no test this is a mock file

// SenderMock : mock implementation you can use in your testing
type SenderMock struct {
	mock.Mock
}

// Send send to someone not in your default configs
func (s *SenderMock) Send(reciever string, bodyContent []byte) error {
	args := s.Called(reciever, bodyContent)
	return args.Error(0)
}

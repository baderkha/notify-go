package notify

import (
	"errors"
	"testing"

	"github.com/baderkha/notify-go/pkg/thirdparty/err"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_Notify_Manager_AddSender(t *testing.T) {
	mgr := new(Manager)
	mockSender := new(SenderMock)
	mgr.AddSender(TestSenderType, mockSender)
	assert.Len(t, mgr.senders, 1)
	mgr.AddSender(TestSenderType+"2", mockSender)
	assert.Len(t, mgr.senders, 2)
}

func Test_Notify_Manager_SendDefaultAll(t *testing.T) {
	// nil map ? should not panic ? pls
	{
		mgr := new(Manager)
		assert.NotPanics(t, func() { err := mgr.SendDefaultAll([]byte("some message")); assert.Nil(t, err) })
	}
	// not nil map , errors out , should return back 1 err
	{
		expectedErr := errors.New("a really scary err")
		msg := []byte("something")
		mgr := new(Manager)
		mgr.senders = make(map[string]MessageSender)
		mockSender := new(SenderMock)
		mgr.senders[TestSenderType] = mockSender
		mockSender.
			On("SendToDefaultReciever", mock.AnythingOfType("[]uint8")).
			Run(func(args mock.Arguments) {
				assert.Equal(t, string(msg), string(args.Get(0).([]byte)))
			}).
			Return(expectedErr)

		err := mgr.SendDefaultAll(msg)
		assert.ErrorIs(t, expectedErr, err)

	}
	// not nil map , more than 1 , should return back more than 1 err
	{
		expectedErr := errors.New("a really scary err")
		msg := []byte("something")
		mgr := new(Manager)
		mgr.senders = make(map[string]MessageSender)
		mockSender := new(SenderMock)
		mgr.senders[TestSenderType] = mockSender
		mgr.senders[TestSenderType+"2"] = mockSender
		mockSender.
			On("SendToDefaultReciever", mock.AnythingOfType("[]uint8")).
			Run(func(args mock.Arguments) {
				assert.Equal(t, string(msg), string(args.Get(0).([]byte)))
			}).
			Return(expectedErr)

		errGot := mgr.SendDefaultAll(msg)
		var erList err.List
		erList.Push(expectedErr)
		erList.Push(expectedErr)
		assert.Equal(t, erList.Error(), errGot.Error())

	}
	// not nil map , no error , should return nil
	{

		msg := []byte("something")
		mgr := new(Manager)
		mgr.senders = make(map[string]MessageSender)
		mockSender := new(SenderMock)
		mgr.senders[TestSenderType] = mockSender
		mockSender.
			On("SendToDefaultReciever", mock.AnythingOfType("[]uint8")).
			Run(func(args mock.Arguments) {
				assert.Equal(t, string(msg), string(args.Get(0).([]byte)))
			}).
			Return(nil)

		err := mgr.SendDefaultAll(msg)
		assert.Nil(t, err)
	}
}

func Test_Notify_Manager_SendAllToSameReceiver(t *testing.T) {
	// nil map ? should not panic ? pls
	{
		mgr := new(Manager)
		assert.NotPanics(t, func() { err := mgr.SendAllToSameReciever(&RecieverAlias{}, []byte("some message")); assert.Nil(t, err) })
	}
	// not nil map , with a nil alias , should error out with specific err
	{
		msg := []byte("hi")
		mgr := new(Manager)
		mgr.senders = make(map[string]MessageSender)
		mockSender := new(SenderMock)
		mgr.senders[TestSenderType] = mockSender
		err := mgr.SendAllToSameReciever(nil, msg)
		assert.ErrorIs(t, errExpectedAlias, err)
	}
	// not nil map , with non nil alias , but no alias for that type , should error
	{
		var a RecieverAlias
		a.Add(SlackSenderType, "something")
		expectedErr := errors.New("a really scary err")
		msg := []byte("something")
		mgr := new(Manager)
		mgr.senders = make(map[string]MessageSender)
		mockSender := new(SenderMock)
		mgr.senders[TestSenderType] = mockSender

		err := mgr.SendAllToSameReciever(&a, msg)
		assert.Error(t, expectedErr, err)

	}
	// not nil map , with non nil alias , with correct alias for that type, sender throws err , should error
	{
		var a RecieverAlias
		a.Add(TestSenderType, "something")
		expectedErr := errors.New("a really scary err")
		msg := []byte("something")
		mgr := new(Manager)
		mgr.senders = make(map[string]MessageSender)
		mockSender := new(SenderMock)
		mgr.senders[TestSenderType] = mockSender
		mockSender.
			On("SendToReciever", mock.AnythingOfType("string"), mock.AnythingOfType("[]uint8")).
			Run(func(args mock.Arguments) {
				assert.Equal(t, a.Get(TestSenderType), args.Get(0).(string))
				assert.Equal(t, string(msg), string(args.Get(1).([]byte)))
			}).
			Return(expectedErr)
		err := mgr.SendAllToSameReciever(&a, msg)
		assert.Error(t, expectedErr, err)

	}
	// not nil map , with non nil alias , with correct alias for that type,
	// senders (more than 1) throws err , should error
	{
		var a RecieverAlias
		a.Add(TestSenderType, "something")
		a.Add(SlackSenderType, "something")
		expectedErr := errors.New("a really scary err")
		msg := []byte("something")
		mgr := new(Manager)
		mgr.senders = make(map[string]MessageSender)
		mockSender := new(SenderMock)
		mgr.senders[TestSenderType] = mockSender
		mgr.senders[SlackSenderType] = mockSender
		mockSender.
			On("SendToReciever", mock.AnythingOfType("string"), mock.AnythingOfType("[]uint8")).
			Run(func(args mock.Arguments) {
				assert.Equal(t, a.Get(TestSenderType), args.Get(0).(string))
				assert.Equal(t, string(msg), string(args.Get(1).([]byte)))
			}).
			Return(expectedErr)
		errGot := mgr.SendAllToSameReciever(&a, msg)
		var erList err.List
		erList.Push(expectedErr)
		erList.Push(expectedErr)
		assert.Equal(t, erList.Error(), errGot.Error())

	}
	// not nil map , no error , should return nil
	{

		var a RecieverAlias
		a.Add(TestSenderType, "something")
		msg := []byte("something")
		mgr := new(Manager)
		mgr.senders = make(map[string]MessageSender)
		mockSender := new(SenderMock)
		mgr.senders[TestSenderType] = mockSender
		mockSender.
			On("SendToReciever", mock.AnythingOfType("string"), mock.AnythingOfType("[]uint8")).
			Run(func(args mock.Arguments) {
				assert.Equal(t, a.Get(TestSenderType), args.Get(0).(string))
				assert.Equal(t, string(msg), string(args.Get(1).([]byte)))
			}).
			Return(nil)
		err := mgr.SendAllToSameReciever(&a, msg)
		assert.Nil(t, err)
	}
}

func Test_Notify_Manager_SendToSpecificSenderType(t *testing.T) {
	// if nothing setup
	{
		mgr := new(Manager)
		err := mgr.SendToSpecificSenderType("someone", "reciever", []byte("some message"))
		assert.Error(t, err)
	}
	// if something setup and failure , i expect to get failure
	{
		var testErr = errors.New("some err")
		mgr := new(Manager)
		mgr.senders = make(map[string]MessageSender)
		mockSender := new(SenderMock)
		mockSender.
			On("SendToReciever", mock.AnythingOfType("string"), mock.AnythingOfType("[]uint8")).
			Run(func(args mock.Arguments) {
				assert.Equal(t, "reciever", args.Get(0).(string))
				assert.Equal(t, "some message", string(args.Get(1).([]byte)))
			}).
			Return(testErr)
		mgr.senders[TestSenderType] = mockSender

		err := mgr.SendToSpecificSenderType(TestSenderType, "reciever", []byte("some message"))
		assert.Error(t, err)
		assert.ErrorIs(t, testErr, err)
	}
	// if something setup and success , i expect to get success
	{
		mgr := new(Manager)
		mgr.senders = make(map[string]MessageSender)
		mockSender := new(SenderMock)

		mockSender.
			On("SendToReciever", mock.AnythingOfType("string"), mock.AnythingOfType("[]uint8")).
			Run(func(args mock.Arguments) {
				assert.Equal(t, "reciever", args.Get(0).(string))
				assert.Equal(t, "some message", string(args.Get(1).([]byte)))
			}).
			Return(nil)

		mgr.senders[TestSenderType] = mockSender
		err := mgr.SendToSpecificSenderType(TestSenderType, "reciever", []byte("some message"))
		assert.Nil(t, err)
	}
}

func Test_Notify_Manager_SendToSpecificSenderTypeDefault(t *testing.T) {
	// if nothing setup
	{
		mgr := new(Manager)
		err := mgr.SendToSpecificSenderTypeDefault("someone", []byte("some message"))
		assert.Error(t, err)
	}
	// if something setup and failure , i expect to get failure
	{
		var testErr = errors.New("some err")
		mgr := new(Manager)
		mgr.senders = make(map[string]MessageSender)
		mockSender := new(SenderMock)
		mockSender.
			On("SendToDefaultReciever", mock.AnythingOfType("[]uint8")).
			Run(func(args mock.Arguments) {
				assert.Equal(t, "some message", string(args.Get(0).([]byte)))
			}).
			Return(testErr)
		mgr.senders[TestSenderType] = mockSender

		err := mgr.SendToSpecificSenderTypeDefault(TestSenderType, []byte("some message"))
		assert.Error(t, err)
		assert.ErrorIs(t, testErr, err)
	}
	// if something setup and success , i expect to get success
	{

		mgr := new(Manager)
		mgr.senders = make(map[string]MessageSender)
		mockSender := new(SenderMock)
		mockSender.
			On("SendToDefaultReciever", mock.AnythingOfType("[]uint8")).
			Run(func(args mock.Arguments) {
				assert.Equal(t, "some message", string(args.Get(0).([]byte)))
			}).
			Return(nil)
		mgr.senders[TestSenderType] = mockSender

		err := mgr.SendToSpecificSenderTypeDefault(TestSenderType, []byte("some message"))
		assert.Nil(t, err)
	}
}

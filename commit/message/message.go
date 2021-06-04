package message

import "errors"

type MsgSrv struct {}

type Service interface {
	CheckMessage() error
}

func NewMessageService() Service {
	return &MsgSrv{}
}

func (m MsgSrv) CheckMessage() error {
	return errors.New("")
}

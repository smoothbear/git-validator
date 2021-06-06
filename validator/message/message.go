package message

import (
	buisError "git-validator/validator/error"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"time"
)

type MsgSrv struct {
	r *git.Repository
}

type Service interface {
	CheckMessage() error
}

func NewMessageService(repository *git.Repository) Service {
	return &MsgSrv{r: repository}
}

func (m MsgSrv) CheckMessage() error {
	ref, err := m.r.Head()
	if err != nil {
		return buisError.WrapError("Could not get HEAD.")
	}

	logTime := time.Now().Add(-30 * time.Second)
	cIter, err := m.r.Log(&git.LogOptions{From: ref.Hash(), Since: &logTime})

	if err != nil {
		return buisError.WrapError("Could not get log.")
	}

	err = cIter.ForEach(func(commit *object.Commit) error {
		print(commit.Message)
		return nil
	})

	return nil
}

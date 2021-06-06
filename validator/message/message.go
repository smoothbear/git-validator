package message

import (
	bisError "git-validator/validator/error"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/spf13/viper"
	"regexp"
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
		return bisError.WrapError("Could not get HEAD.")
	}

	logTime := time.Now().Add(-30 * time.Second)
	cIter, err := m.r.Log(&git.LogOptions{From: ref.Hash(), Since: &logTime})

	if err != nil {
		return bisError.WrapError("Could not get log.")
	}

	err = cIter.ForEach(func(commit *object.Commit) error {
		result, err := regexp.Match(viper.GetString("regex"), []byte(commit.Message))
		if err != nil {
			return bisError.WrapError("Commit regex is not valid.")
		}

		if result == false {
			return bisError.WrapError("Commit pattern is not matched.")
		}

		return nil
	})

	return err
}

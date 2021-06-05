package main

import (
	"fmt"
	validationError "git-validator/validator/error"
	"git-validator/validator/message"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/spf13/viper"
	"log"
)

func main() {
	// Setting environments from gcmc.yaml
	viper.SetConfigName("gcmc")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileAlreadyExistsError); ok {
			log.Fatalf("Error read configuration file.")
		}
	}

	// If it is empty value, running by default environment
	if viper.GetString("regex") == "" {
		// Example: [ADD] Initial Commit
		viper.Set("regex", "(?i)([+[A-Z])+\\w+]")
	}

	println(viper.GetString("regex"))

	repository, err := git.PlainOpen(".")
	checkIfErr(err)

	ref, err := repository.Head()
	checkIfErr(err)

	cIter, err := repository.Log(&git.LogOptions{From: ref.Hash()})
	checkIfErr(err)

	err = cIter.ForEach(func(c *object.Commit) error {
		fmt.Println(c)
		return nil
	})

	msgSrv := message.NewMessageService()

	switch err := msgSrv.CheckMessage(); err.(type) {
	case validationError.ValidationError:
		log.Fatalf(err.(validationError.ValidationError).GetMessage())
	default:
		log.Print("Git validation is successfully completed!")
	}
}

func checkIfErr(err error) {
	if err != nil {
		log.Fatalf("Something went wrong.")
	}
}

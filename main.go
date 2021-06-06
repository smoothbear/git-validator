package main

import (
	validationError "git-validator/validator/error"
	"git-validator/validator/message"
	"github.com/go-git/go-git/v5"
	"github.com/spf13/viper"
	"log"
)

func main() {
	// Setting environments from gcmc.yaml
	viper.SetConfigName("gcmc")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	_ = viper.ReadInConfig()

	// If it is empty value, running by default environment
	if viper.GetString("regex") == "" {
		// Example: [ADD] Initial Commit
		viper.Set("regex", "([+[A-Z])+\\w+]")
	}

	repository, err := git.PlainOpen(".")
	checkIfErr(err)

	msgSrv := message.NewMessageService(repository)

	switch err := msgSrv.CheckMessage(); err.(type) {
	case validationError.ValidationError:
		log.Fatalf("❌  Failed: %s", err.(validationError.ValidationError).GetMessage())
	case nil:
		log.Print("✅  Git validation is successfully completed!")
	default:
		log.Fatalf("❌  Failed: %v", err)
	}
}

func checkIfErr(err error) {
	if err != nil {
		log.Fatalf("Something went wrong.")
	}
}

package main

import (
	validationError "git-validator/commit/error"
	"git-validator/commit/message"
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
			log.Fatalf("Error read configuration file. Please check gcmc.yaml is exist on your repository root path.")
		} else {
			log.Fatalf("Error: %v", ok)
		}
	}

	// If it is empty value, running by default environment
	if viper.GetString("regex") == "" {
		viper.Set("regex", "(?i)([+[A-Z])+\\w+]")
	}

	msgSrv := message.NewMessageService()

	switch err := msgSrv.CheckMessage(); err.(type) {
	case validationError.ValidationError:
		log.Fatalf(err.(validationError.ValidationError).GetMessage())
	default:
		log.Print("Git validation is successfully completed!")
	}
}

package error

type ValidationError struct {
	error
	message string
}

func (v ValidationError) GetMessage() string {
	return v.message
}

func WrapError(msg string) error {
	return ValidationError{
		message: msg,
	}
}

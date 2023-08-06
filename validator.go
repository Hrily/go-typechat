package typechat

type Validator interface {
	Validate() error
}

type ValidationError struct {
	Message string
}

var _ error = ValidationError{}

func (e ValidationError) Error() string {
	return e.Message
}

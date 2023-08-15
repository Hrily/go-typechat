package typechat

// Validator to perform additional validations on the response
type Validator interface {
	// Validate the response and return `ValidationError` if it fails validation
	Validate() error
}

// ValidationError is an error returned when validation fails
// `Message` will be used in diagnostics to repair the response
type ValidationError struct {
	Message string
}

// Ensure ValidationError implements error
var _ error = ValidationError{}

// Error ...
func (e ValidationError) Error() string {
	return e.Message
}

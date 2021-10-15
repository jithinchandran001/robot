package codecs

type EndpointError struct {
	Code    int
	Message string
}

func NewEndpointError(code int, message string) *EndpointError {
	return &EndpointError{
		Code:    code,
		Message: message,
	}
}

func (e *EndpointError) Error() string {
	return "Endpoint Error: " + e.Message
}

// IsEndpointError()
func IsEE(err error) bool {
	if _, ok := err.(*EndpointError); ok {
		return true
	}
	return false
}

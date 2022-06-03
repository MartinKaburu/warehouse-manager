package errors

type ServerError struct {
	Message string
	Code    string
}

func (err *ServerError) Error() string {
	return err.Code + ": " + err.Message
}

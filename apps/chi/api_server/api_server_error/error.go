package api_server_error

type RecordNotFoundError struct{}

func (e RecordNotFoundError) Error() string {
	return "RecordNotFoundError"
}

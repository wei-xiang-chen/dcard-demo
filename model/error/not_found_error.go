package error

type NotFoundError struct {
	Message string
}

func (e NotFoundError) Error() string {
	return e.Message
}

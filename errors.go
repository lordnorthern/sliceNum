package sliceNum

type GeneralError struct {
	Code int
	msg  string
}

func (c GeneralError) Error() string {
	return c.msg
}

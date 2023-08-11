package errs

type ReadError struct {
	Message string
}

func (err *ReadError) Error() string {
	return err.Message
}

type OpenFileError struct {
	Message string
}

func (err *OpenFileError) Error() string {
	return err.Message
}

type AbsentFileError struct {
	Message string
}

func (err *AbsentFileError) Error() string {
	return err.Message
}

package xerr

import "fmt"

const (
	CodeSuccess       = 0
	CodeInvalidParams = 1001
	CodeNotFound      = 1002
	CodeAlreadExists  = 1003
	CodeInternal      = 9999
)

type CodeError struct {
	Code  int
	Msg   string
	Cause error
}

// Error 返回错误信息
func (e CodeError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%v", e.Cause)
	}
	return e.Msg
}

func (e CodeError) Unwarp() error {
	return e.Cause
}

func New(code int, msg string) *CodeError {
	return &CodeError{
		Code: code,
		Msg:  msg,
	}
}

func Wrap(code int, msg string, cause error) *CodeError {
	return &CodeError{
		Code:  code,
		Msg:   msg,
		Cause: cause,
	}
}

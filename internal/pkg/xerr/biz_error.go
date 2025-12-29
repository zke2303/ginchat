package xerr

import "fmt"

type BizError struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Cause error  `json:"cause"`
}

// 实现Error()方法 == 继承 error
func (bizErr *BizError) Error() string {
	if bizErr.Cause != nil {
		return fmt.Sprintf("%s, %v", bizErr.Msg, bizErr.Cause)
	}
	return bizErr.Msg
}

// 全局常量(业务常用错误)
var (
	// global
	ErrInvalidParam = &BizError{Code: 40001, Msg: "Invalid Params"}
	ErrUnanthorizad = &BizError{Code: 40100, Msg: "Unauthorized"}
	// user module
	ErrUserNotFount     = &BizError{Code: 20001, Msg: "User not exists"}
	ErrUserAlreadyExist = &BizError{Code: 20002, Msg: "User has Already been Exists"}
	ErrUsernameTaken    = &BizError{Code: 20003, Msg: "Username haa Already been Exists"}
	ErrEmailTaken       = &BizError{Code: 20004, Msg: "Email has Already been Exists"}
)

// NewBizError 辅助创建函数
func NewBizError(code int, msg string, cause error) *BizError {
	return &BizError{Code: code, Msg: msg, Cause: cause}
}

// WrapBiz
func WrapBiz(code int, msg string, err error) error {
	if err == nil {
		return nil
	}

	return &BizError{Code: code, Msg: msg, Cause: err}
}

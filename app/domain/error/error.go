package error

import "encoding/json"

type Error struct {
	Err     error  `json:"error"`
	Type    string `json:"type"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Reason  string `json:"reason"`
}

func (e Error) Error() string {
	return e.Err.Error()
}

func (e Error) Format() []byte {
	res, _ := json.Marshal(e)
	return res
}

// 400:クライアントから無効なリクエストが送信されたとき
func NewBadRequestError(err error, reason string) *Error {
	res := new(Error)
	res.Err = err
	res.Type = "InvalidArgument"
	res.Code = 400
	res.Message = "BadRequest"
	res.Reason = reason
	return res
}

// 401:認証エラー
func NewUnauthorizedError(err error, reason string) *Error {
	res := new(Error)
	res.Err = err
	res.Type = "Unauthenticated"
	res.Code = 401
	res.Message = "Unauthorized"
	res.Reason = reason
	return res
}

// 403:認可エラー
func NewForbiddenError(err error, reason string) *Error {
	res := new(Error)
	res.Err = err
	res.Type = "PermissionDenied"
	res.Code = 403
	res.Message = "Forbidden"
	res.Reason = reason
	return res
}

// 404:要素が見つからなかった
func NewNotFoundError(err error, reason string) *Error {
	res := new(Error)
	res.Err = err
	res.Type = "NotFound"
	res.Code = 404
	res.Message = "NotFound"
	res.Reason = reason
	return res
}

// 409:コンフリクトが発生
func NewConflictError(err error, reason string) *Error {
	res := new(Error)
	res.Err = err
	res.Type = "Aborted"
	res.Code = 409
	res.Message = "Conflict"
	res.Reason = reason
	return res
}

// 500:サーバーエラー
func NewInternalServerError(err error, reason string) *Error {
	res := new(Error)
	res.Err = err
	res.Type = "Internal"
	res.Code = 500
	res.Message = "InternalServerError"
	res.Reason = reason
	return res
}

// 504:外部通信タイムアウトエラー
func NewGatewayTimeoutError(err error, reason string) *Error {
	res := new(Error)
	res.Err = err
	res.Type = "DeadlineExceeded"
	res.Code = 504
	res.Message = "GatewayTimeout"
	res.Reason = reason
	return res
}

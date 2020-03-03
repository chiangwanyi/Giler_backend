package serializer

const (
	// OK 正常
	OK = 0
	// BadAuthError 无权限
	BadAuthError = 1
	// BadParameterError 参数或字段不匹配
	BadParameterError = 2
	// InternalServerError 内部错误
	InternalServerError = 3
	// ExistingError 数据已经存在
	ExistingError = 4
)

// Response 基础序列化器
type Response struct {
	Status uint        `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
	Error  interface{} `json:"error"`
}

// BadAuthErrorResponse 返回无权限错误消息
func BadAuthErrorResponse(info string) Response {
	return Response{
		Status: BadAuthError,
		Data:   nil,
		Msg:    "权限错误：" + info,
		Error:  nil,
	}
}

// BadParameterErrorResponse 返回参数或字段匹配错误消息
func BadParameterErrorResponse(err error) Response {
	return Response{
		Status: BadParameterError,
		Data:   nil,
		Msg:    "参数或字段匹配错误：请检查参数",
		Error:  err,
	}
}

// InternalServerErrorResponse 返回内部错误消息
func InternalServerErrorResponse(err error, info string) Response {
	return Response{
		Status: InternalServerError,
		Data:   nil,
		Msg:    "内部错误：" + info,
		Error:  err,
	}
}

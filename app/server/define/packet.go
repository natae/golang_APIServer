package define

type LoginRequest struct {
	Nickname string
}

type BaseResponse struct {
	ResultCode int
	Response   interface{}
}
type LoginResponse struct {
	Id int64
}

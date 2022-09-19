package define

type BaseResponse struct {
	ResultCode int
	Response   interface{}
}

type LoginRequest struct {
	Nickname string
}

type LoginResponse struct {
	Id int64
}

type JoinRequest struct {
	Nickname string
}

type JoinReponse struct {
	Id int64
}

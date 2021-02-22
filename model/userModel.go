package model

import "time"

type ReqLogin struct {
	Nickname string `json:"nickname"`
}
type ResLogin struct {
	Result int16
	Session Session
}
type Session struct {
	AuthToken string
	Nickname string
	ExpireTime time.Time
}

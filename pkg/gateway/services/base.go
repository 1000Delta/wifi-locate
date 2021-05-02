package services

type baseResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg,omitempty"`
}

func Resp(code int, msg string) baseResp {
	return baseResp{code, msg}
}

package robot

type Response struct {
	Code    int    `json:"errcode"`
	Message string `json:"errmsg"`
}

func (e *Response) IsSuccess() bool {
	return e.Code == 0
}

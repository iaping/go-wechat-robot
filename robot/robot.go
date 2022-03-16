package robot

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/iaping/go-wechat-robot/robot/message"
	"github.com/valyala/fasthttp"
)

var (
	defaultClient = &fasthttp.Client{
		Name:                "go-wechat-robot",
		MaxConnsPerHost:     16,
		MaxIdleConnDuration: 10 * time.Second,
		ReadTimeout:         5 * time.Second,
		WriteTimeout:        5 * time.Second,
	}
)

type Robot struct {
	Hook   string
	Client *fasthttp.Client
}

func New(hook string) *Robot {
	return &Robot{
		Hook:   hook,
		Client: defaultClient,
	}
}

func (r *Robot) Send(m message.IMessage) (*Response, error) {
	body, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
	}()

	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")
	req.SetRequestURI(r.Hook)
	req.SetBody(body)

	if err := r.Client.Do(req, resp); err != nil {
		return nil, err
	}

	body = resp.Body()

	if resp.StatusCode() != fasthttp.StatusOK {
		return nil, fmt.Errorf("http status code:%d, desc:%s", resp.StatusCode(), string(body))
	}

	var result Response
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

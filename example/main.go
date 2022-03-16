package main

import (
	"log"

	"github.com/iaping/go-wechat-robot/robot"
	"github.com/iaping/go-wechat-robot/robot/message"
)

func main() {
	wxRobot := robot.New("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=35ce8f27-4c61-4c9e-ac66-xxxxxxxx")

	// text
	text := message.NewTextSimple("hello world!", false)
	resp, err := wxRobot.Send(text)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("success:", resp.IsSuccess(), "code:", resp.Code, "message:", resp.Message)

	// markdown
	markdown := message.NewMarkdownSimple(`实时新增用户反馈<font color=\"warning\">132例</font>，请相关同事注意。\n
	>类型:<font color=\"comment\">用户反馈</font>
	>普通用户反馈:<font color=\"comment\">117例</font>
	>VIP用户反馈:<font color=\"comment\">15例</font>`)
	resp, err = wxRobot.Send(markdown)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("success:", resp.IsSuccess(), "code:", resp.Code, "message:", resp.Message)
}

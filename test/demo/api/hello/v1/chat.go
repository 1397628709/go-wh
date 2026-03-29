package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type ChatReq struct {
	g.Meta   `path:"/chat" method:"get" summary:"对话"`
	Id       string `json:"id" v:"required#ID不能为空"`
	Question string `json:"question" v:"required#问题不能为空"`
}

type ChatRes struct {
	Answer string `json:"answer"`
}

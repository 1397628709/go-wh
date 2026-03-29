package hello

import (
	"context"
	"fmt"

	v1 "demo/api/hello/v1"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) Chat(ctx context.Context, req *v1.ChatReq) (res *v1.ChatRes, err error) {
	answer := fmt.Sprintf("您好!您的ID是: %s, 您的问题是: %s", req.Id, req.Question)
	
	res = &v1.ChatRes{
		Answer: answer,
	}
	
	g.Log().Infof(ctx, "Chat request - ID: %s, Question: %s", req.Id, req.Question)
	
	return res, nil
}

package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	pb "slient.util/generate/proto"
)

var svc = NewServer()
var ctx = context.Background()

func Test(t *testing.T) {
	resp, err := svc.Hello(ctx, &pb.HelloReq{Req: ""})
	assert.Equal(t, nil, err)
	t.Log(resp)
}

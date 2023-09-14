package handler

import (
	"context"

	"github.com/ahakra/servicediscovery/loggerService/writer/internal/controller"
	"github.com/ahakra/servicediscovery/loggerService/writer/internal/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type LogWritter struct {
	proto.UnimplementedLogwriterServer
	Ctrl controller.IMongoCtrl
}

func NewLogReaderHandler(mgctrl controller.IMongoCtrl) *LogWritter {
	return &LogWritter{Ctrl: mgctrl}
}

func (c *LogWritter) SaveLog(ctx context.Context, in *proto.LogPayload) (*emptypb.Empty, error) {

	_, err := c.Ctrl.SaveLog(ctx, in)

	return &emptypb.Empty{}, err
}

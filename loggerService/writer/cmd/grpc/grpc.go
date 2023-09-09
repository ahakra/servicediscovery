package main

import (
	"context"

	"github.com/ahakra/servicediscovery/loggerService/writer/internal/controller"
	"github.com/ahakra/servicediscovery/loggerService/writer/internal/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type LogWritter struct {
	proto.UnimplementedLogwriterServer
	Ctrl *controller.MongoCtrl
}

func (c *LogWritter) SaveLog(ctx context.Context, in *proto.LogData) (*emptypb.Empty, error) {

	_,err := c.Ctrl.SaveLog(ctx, in)
	
	return &emptypb.Empty{},err
}

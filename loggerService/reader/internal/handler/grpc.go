package handler

import (
	"context"

	"github.com/ahakra/servicediscovery/loggerService/reader/internal/controller"
	"github.com/ahakra/servicediscovery/loggerService/reader/internal/proto"
)

type LogReader struct {
	proto.UnimplementedLogReaderServer
	Ctrl controller.IMongoCtrl
}

func NewLogReaderHandler(mgctrl controller.IMongoCtrl) *LogReader {
	return &LogReader{Ctrl: mgctrl}
}

func (c *LogReader) ReadLog(ctx context.Context, in *proto.LogFilter) (*proto.ReturnedData, error) {

	returneddata, err := c.Ctrl.ReadLog(ctx, in)

	if err != nil {
		return &proto.ReturnedData{}, nil
	}

	return returneddata, nil
}

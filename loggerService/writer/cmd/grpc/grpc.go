package main

import "github.com/ahakra/servicediscovery/loggerService/writer/internal/proto"

type LogWritter struct {
	proto.UnimplementedLogWritterServer
}

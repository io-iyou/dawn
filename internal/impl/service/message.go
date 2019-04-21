package service

import (
	"context"
	"fmt"

	"github.com/gogo/protobuf/types"
	pb "github.com/iyou/dawn/service/sdk/go"
	"google.golang.org/grpc/grpclog"
)

var (
	queqeMap = make(map[string]chan *pb.Message)
)

type MessageImpl struct{}

func (s *MessageImpl) Send(ctx context.Context, in *pb.Message) (*types.Empty, error) {
	if queue := queqeMap[in.To]; queue == nil {
		queqeMap[in.To] = make(chan *pb.Message, 100)
	}
	fmt.Println("rec:", *in)
	go func() {
		queqeMap[in.To] <- in
	}()

	return &types.Empty{}, nil
}

func (s *MessageImpl) Receive(in *pb.User, stream pb.Messages_ReceiveServer) error {
	if queue := queqeMap[in.Id]; queue == nil {
		queqeMap[in.Id] = make(chan *pb.Message, 100)
	}

	for msg := range queqeMap[in.Id] {
		fmt.Println("send", msg)
		if err := stream.Send(msg); err != nil {
			grpclog.Errorln(err)
			queqeMap[in.Id] <- msg
			//return err
		}
	}

	return nil
}

package service

import (
	"context"
	"fmt"

	"github.com/gogo/protobuf/types"
	pb "github.com/iyou/dawn/service/sdk/go"
	"google.golang.org/grpc/grpclog"
)

var (
	p2pQueues    = make(map[string]chan *pb.Message)
	pubsubQueues = make(map[string][]chan *pb.Message)
)

type MessageImpl struct{}

func (s *MessageImpl) Send(ctx context.Context, in *pb.Message) (*types.Empty, error) {
	if p2pQueues[in.To] == nil {
		p2pQueues[in.To] = make(chan *pb.Message, 100)
	}
	fmt.Println("rec:", *in)
	go func() {
		p2pQueues[in.To] <- in
	}()

	return &types.Empty{}, nil
}

func (s *MessageImpl) Receive(in *pb.User, stream pb.Messages_ReceiveServer) error {
	if p2pQueues[in.Id] == nil {
		p2pQueues[in.Id] = make(chan *pb.Message, 100)
	}

	for msg := range p2pQueues[in.Id] {
		fmt.Println("send", msg)
		if err := stream.Send(msg); err != nil {
			grpclog.Errorln(err)
			p2pQueues[in.Id] <- msg
			//return err
		}
	}

	return nil
}

func (s *MessageImpl) Publish(ctx context.Context, in *pb.TopicRequest) (*types.Empty, error) {
	if pubsubQueues[in.Topic] == nil {
		pubsubQueues[in.Topic] = []chan *pb.Message{}
	}

	go func() {
		for k := range pubsubQueues[in.Topic] {
			pubsubQueues[in.Topic][k] <- in.Msg
		}
	}()

	return nil, nil
}

func (s *MessageImpl) Subscribe(in *pb.TopicRequest, stream pb.Messages_SubscribeServer) error {
	if pubsubQueues[in.Topic] == nil {
		pubsubQueues[in.Topic] = []chan *pb.Message{}
	}

	queue := make(chan *pb.Message, 100)
	pubsubQueues[in.Topic] = append(pubsubQueues[in.Topic], queue)
	for msg := range queue {
		if err := stream.Send(msg); err != nil {
			grpclog.Errorln(err)
			queue <- msg
		}
	}

	return nil
}

package main

import (
	"log"
	"net"
	"sync"

	"golang.org/x/net/context"

	"google.golang.org/grpc"

	pb "github.com/nerocrux/grpc-example/proto"
)

type awesomeService struct {
	members   []*pb.Member
	m         sync.Mutex
}

func (cs *awesomeService) ListMember(p *pb.RequestType, stream pb.Nogizaka_ListMemberServer) error {
	log.Println("ListMember")
	cs.m.Lock()
	defer cs.m.Unlock()
	for _, p := range cs.members {
		if err := stream.Send(p); err != nil {
			return err
		}
	}
	return nil
}

func (cs *awesomeService) AddMember(c context.Context, p *pb.Member) (*pb.ResponseType, error) {
	log.Println("AddMember: " + p.Name)
	cs.m.Lock()
	defer cs.m.Unlock()
	cs.members = append(cs.members, p)
	return new(pb.ResponseType), nil
}

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()

	pb.RegisterNogizakaServer(server, new(awesomeService))
	log.Println("Listing...")
	server.Serve(lis)
}

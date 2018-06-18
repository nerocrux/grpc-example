package main

import (
	"fmt"
	"io"
	"strconv"

	"github.com/mattn/sc"
	pb "github.com/nerocrux/grpc-example/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func add(id int, name string, birthday string, generation int) error {
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()
	client := pb.NewNogizakaClient(conn)

	member := &pb.Member{
		Id:  int64(id),
		Name: name,
		Birthday: birthday,
		Generation:  int32(generation),
	}
    fmt.Println(member)
	_, err = client.AddMember(context.Background(), member)
	return err
}

func list() error {
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()
	client := pb.NewNogizakaClient(conn)

	stream, err := client.ListMember(context.Background(), new(pb.RequestType))
	if err != nil {
		return err
	}
	for {
		member, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fmt.Println(member)
	}
	return nil
}

func main() {
	(&sc.Cmds{
		{
			Name: "list",
			Desc: "list: listing member",
			Run: func(c *sc.C, args []string) error {
				return list()
			},
		},
		{
			Name: "add",
			Desc: "add [id] [name] [birthday] [generation]: add Nogizaka member",
			Run: func(c *sc.C, args []string) error {
				if len(args) != 4 {
					return sc.UsageError
				}
				id, err := strconv.Atoi(args[0])
				name := args[1]
				birthday := args[2]
				generation, err := strconv.Atoi(args[3])
				if err != nil {
					return err
				}
				return add(id, name, birthday, generation)
			},
		},
	}).Run(&sc.C{})
}

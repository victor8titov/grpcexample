package main

import (
	"fmt"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var port = ":8080"

func AboutToSayIt(ctx context.Context, m MessageServiceClient, text string) (*p.Response, error) {
	request := &Request{
		Text:    text,
		Subtext: "New Message!",
	}

	r, err := m.SayIt(ctx, request)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func main() {
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Dial:", err)
		return
	}

	client := p.NewMessageServiceClient(conn)
	r, err := AboutToSayIt(context.Background(), client, "My Message!")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Response Text:", r.Text)
	fmt.Println("Response SubText:", r.Subtext)
}

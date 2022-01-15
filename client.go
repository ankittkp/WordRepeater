package main

import (
	"context"
	"fmt"
	pb "github.com/jinxankit/WordRepeater/proto/frequency"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Hello client ...")

	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	client := pb.NewWordFrequencyClient(cc)
	request := &pb.InputRequest{Text: "ankit ankit"}

	resp, _ := client.Calculate(context.Background(), request)
	fmt.Printf("Receive response => [%v]", resp.Result)
}
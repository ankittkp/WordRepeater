package main

import (
	"context"
	"fmt"
	pb "github.com/jinxankit/WordRepeater/proto/frequency"
	"google.golang.org/grpc"
	"log"
	"net"
	"sort"
	"strings"
)
type server struct{
	pb.UnimplementedWordFrequencyServer
}
type newMap struct {
	Key string
	Value int64
}
func getTopTenElement(m map[string]int64) map[string]int64 {
	var mapSlice []newMap
	for k, v := range m {
		mapSlice = append(mapSlice, newMap{k, v})
	}
	sort.Slice(mapSlice, func(i, j int) bool {
		return mapSlice[i].Value > mapSlice[j].Value
	})
	newWordCount := make(map[string]int64)
	for _, kv := range mapSlice[:10] {
		newWordCount[kv.Key] = kv.Value
		fmt.Printf("%s, %d\n", kv.Key, kv.Value)
	}
	return newWordCount
}
func (s *server) Calculate(_ context.Context, req *pb.InputRequest) (*pb.OutputResponse, error){
	input := req.Text
	splitWord := strings.Fields(input)
	wordCount := make(map[string]int64)
	for _, value := range splitWord{
		_ , alreadyExist := wordCount[value]
		if alreadyExist{
			wordCount[value] += 1
		} else {
			wordCount[value] = 1
		}
	}
	wordCount = getTopTenElement(wordCount)
	for i,j := range wordCount{
		fmt.Printf("Word is : %v and its count is %v: ", i, j)
		fmt.Println()
	}
	return &pb.OutputResponse{
		Result: wordCount,
	}, nil
}
func main(){
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	s := grpc.NewServer()
	pb.RegisterWordFrequencyServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
	log.Printf("Hosting server on: %s", lis.Addr().String())
}

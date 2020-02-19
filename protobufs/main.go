package main

import (
	"fmt"

	simple2pb "github.com/jmarlow4/learning-go/protobufs/src/simple"
)

func main() {
	newMessage := doSimple()
	newMessage.Name = "Changing the name..."
	fmt.Println(newMessage.GetName())
}

func doSimple() *simple2pb.SimpleMessage {
	sm := simple2pb.SimpleMessage{
		Id:         16,
		IsSimple:   true,
		Name:       "Just a simple message",
		SampleList: []int32{8, 16, 32, 64},
	}
	return &sm
}

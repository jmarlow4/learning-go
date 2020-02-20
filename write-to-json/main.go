package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	simple2pb "github.com/jmarlow4/learning-go/protobufs/src/simple"
)

func main() {
	newMessage := doSimple()

	nmAsString := toJson((newMessage))

	fmt.Println("Converting to JSON", nmAsString)
	nm2 := &simple2pb.SimpleMessage{}
	fromJson(nmAsString, nm2)
	fmt.Println("Converting FROM JSON", nm2)
}

func fromJson(in string, pb proto.Message) {
	err := jsonpb.UnmarshalString(in, pb)
	if err != nil {
		log.Fatalln("Couldn't unmarshal the JSON into the pb struct", err)
	}
}

func toJson(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString((pb))
	if err != nil {
		log.Fatalln("Can't convert to JSON", err)
	}
	return out
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

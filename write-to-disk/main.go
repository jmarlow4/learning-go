package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/proto"
	simple2pb "github.com/jmarlow4/learning-go/protobufs/src/simple"
)

func main() {
	newMessage := doSimple()
	
	readAndWriteDemo(newMessage)
}

func readAndWriteDemo(sm proto.Message) {
	writeToFile("simple.bin", sm)
	sm2 := &simple2pb.SimpleMessage{}
	readFromFile("simple.bin", sm2)
	fmt.Println("Read the content:", sm2)
}

func writeToFile(filename string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialize to bytes", err)
		return err
	}

	if err := ioutil.WriteFile(filename, out, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
		return err
	}
	
	fmt.Println("Data has been written!")
	return nil
}

func readFromFile(filename string, pb proto.Message) error {
	in, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatalln("Something went wrong when reading the file.")
		return err
	}

	err2 := proto.Unmarshal(in, pb)
	if err2 != nil {
		log.Fatalln("Something went wrong when reading the file.")
		return err2
	}

	return nil


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

package main

import (
	"fmt"

	"github.com/fwv/protodec/pb2proto"
	"github.com/fwv/protodec/testingpb"
)

func main() {
	str, err := pb2proto.Analyze(testingpb.File_animal_proto)
	if err != nil {
		fmt.Printf("failed to protodec. %v", err)
	}
	fmt.Printf("%s", str)
}

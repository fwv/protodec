package testingpb

import (
	"fmt"
	"testing"

	"github.com/fwv/protodec/pb2proto"
	_ "github.com/golang/protobuf/protoc-gen-go/descriptor" // required for option

	"github.com/stretchr/testify/assert"
)

func TestEnumeration(t *testing.T) {
	fd := File_enumeration_proto
	str, err := pb2proto.Analyze(fd)
	assert.NoError(t, err)
	fmt.Printf("%v", str)
}

func TestNestedMsg(t *testing.T) {
	fd := File_nested_proto
	str, err := pb2proto.Analyze(fd)
	assert.NoError(t, err)
	fmt.Printf("%v", str)
}

func TestOptions(t *testing.T) {
	fd := File_animal_proto
	str, err := pb2proto.Analyze(fd)
	assert.NoError(t, err)
	fmt.Printf("%v", str)
}

# protodec
a protobuf decompile tool which could convert pb.go to .proto format string in go runtime. (both proto2 or proto3 syntax is supported)

## Usage
Suppose we have a proto file named animal.proto like this:
```protobuf
syntax = "proto3";                 
package animal;                 
option go_package = "protodec/testingpb";
import "myoptions.proto"; 

message Cat{  
  option(myoptions.animal_color) = "white";
  option(myoptions.animal_alias) = "kitty";
  option(myoptions.animal_alias) = "pussy";
  int64 animal_id = 1;  
  string owner_name = 2[(myoptions.owner_id)=666,(myoptions.owner_name)="fwv",(myoptions.owner_is_male)=true]; 
}

```
According to the file above, we can see it contains protobuf customer option feature, but we are not going to show myoptions.proto here, you could find both of these proto files in testingpb package in this repository. 

Then we can get animal.pb.go file through protocol buffer compiler by command like `protoc --proto_path=xxx --go_out=xxx --go_opt=xxx  xxx/animal.proto`, you can find more details in [protobuf go generated code reference](https://developers.google.com/protocol-buffers/docs/reference/go-generated).

Let's dig in this animal.pb.go, we could find a `protorefect.FileDescriptor` variable named `File_animal_proto`, now use this variable to reconstruct proto file format string in Golang runtime like this：
```go
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
```

The reconstructed proto file string looks like this:
```protobuf
syntax = "proto3";
package animal;
option go_package = "protodec/testingpb";
import "myoptions.proto";

message Cat {
  option(myoptions.animal_color) = "white";
  option(myoptions.animal_alias) = "kitty";
  option(myoptions.animal_alias) = "pussy";
  int64 animal_id = 1;
  string owner_name = 2[(myoptions.owner_id) = 666,(myoptions.owner_name) = "fwv",(myoptions.owner_is_male) = true];
}
```

## Features
- proto2 and proto3
- nested element
- enumeration
- options(both built-in and custom)
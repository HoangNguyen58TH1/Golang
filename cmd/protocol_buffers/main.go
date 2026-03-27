package main

import (
	"fmt"
	"log"

	pb "golang_learning/utility/protocol_buffers"

	"google.golang.org/protobuf/proto"
)

// go run cmd/protocol_buffers/main.go
func main() {
	// create protobuf struct
	user := &pb.User{
		Name: "Alice",
		Age:  25,
		Job:  "Engineer",
		Sex:  "Female",
	}

	// Serialize struct -> binary
	data, err := proto.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Serialized bytes:", data)

	// Deserialize binary -> struct
	var decoded pb.User

	err = proto.Unmarshal(data, &decoded)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Decoded data:")
	fmt.Println(decoded.Name, decoded.Age, decoded.Job, decoded.Sex)
}

// Go struct
//    │
// protobuf Marshal
//    │
// TCP / UDP send (binary bytes)
//    │
// network
//    │
// receive
//    │
// protobuf Unmarshal
//    │
// Go struct

// file utility/protocol_buffers/user.proto auto create new file utility/protocol_buffers/user.pb.go

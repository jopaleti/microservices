package serializer

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ProtobufToJSON converts protocol buffer message to JSON string
func ProtobufToJSON(message proto.Message) (string, error) {
	marshaler := protojson.MarshalOptions{
		UseProtoNames:   true, // Use field names from the proto definition
        EmitUnpopulated: true, // Emit fields with zero values
        Indent:          "	", // Indentation for JSON formatting
	}
	jsonBytes, err := marshaler.Marshal(message)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}
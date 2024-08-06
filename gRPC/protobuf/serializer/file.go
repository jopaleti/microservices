package serializer

import (
	"fmt"
	"os"

	"google.golang.org/protobuf/proto"
)

func WriteProtobufToJSONFile(message proto.Message, filename string) error {
	// Convert protobuf message to JSON format
	data, err := ProtobufToJSON(message)
	if err != nil {
		return fmt.Errorf("Cannot marshal proto message to binary: %w", err)
	}
	err = os.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		return fmt.Errorf("Cannot write JSON data to file: %w", err)
	}
	return nil
}

func WriteProtobufToBinaryFile(message proto.Message, filename string) error {
	// Serialize the message to binary first
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("Cannot Marshal proto message to binary: %w", err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("Cannot write binary data to file: %w", err)
	}
	return nil
}

func ReadProtoBufFromBinaryFile(filename string, message proto.Message) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("Cannot read binary data from file: %v",  err)
	}

	err = proto.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("Cannot Unmarshal binary to proto message: %v", err)
	}
	return nil
}
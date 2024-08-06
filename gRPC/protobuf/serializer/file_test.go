package serializer_test

import (
	"testing"

	"github.com/jopaleti/pcbook/pb"
	"github.com/jopaleti/pcbook/sample"
	"github.com/jopaleti/pcbook/serializer"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel() // For paralle running of unit test

	binaryFile := "../tmp/laptop.bin"

	// Creating a new instance of laptop
	laptop1 := sample.NewLaptop()
	err := serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)

	laptop2 := &pb.Laptop{}
	err = serializer.ReadProtoBufFromBinaryFile(binaryFile, laptop2)
	require.NoError(t, err)
	require.True(t, proto.Equal(laptop1, laptop2))
}
package service_test

import (
	"context"
	"testing"

	"github.com/jopaleti/pcbook/pb"
	"github.com/jopaleti/pcbook/sample"
	"github.com/jopaleti/pcbook/service"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// WRITING TABLE-DRIVEN TEST IN GO
func TestServerCreateLaptop(t *testing.T) {
	t.Parallel()

	laptopNoID := sample.NewLaptop()
	laptopNoID.Id = ""

	laptopInvalidID := sample.NewLaptop()
	laptopInvalidID.Id = "invalid-uuid"

	laptopDuplicateID := sample.NewLaptop()
	storeDuplicateID := service.NewInMemoryLaptopStore()
	err := storeDuplicateID.Save(laptopDuplicateID)
	require.Nil(t, err)

	testCases := []struct{
		name string
		laptop *pb.Laptop
		store service.LaptopStore
		code codes.Code
	} {
		{
			name: "success_with_id",
			laptop: sample.NewLaptop(),
			store: service.NewInMemoryLaptopStore(),
			code: codes.OK,
		},
		{
			name: "success_with_no_id",
			laptop: laptopNoID,
			store: service.NewInMemoryLaptopStore(),
			code: codes.OK,
		},
		{
			name: "failure_invalid_id",
			laptop: laptopInvalidID,
			store: service.NewInMemoryLaptopStore(),
			code: codes.InvalidArgument,
		},
		{
			name: "failure_duplicate_id",
			laptop: laptopDuplicateID,
			store: storeDuplicateID,
			code: codes.AlreadyExists,
		},
	}

	for i := range testCases {
		/* Save the current testcase to a local variable, 
			it's important to avoid concurrency issues, 
			as we will run multiple test cases */
		tc := testCases[i]

		t.Run(tc.name, func (t *testing.T) {
			t.Parallel()

			req := &pb.CreateLaptopRequest{
				Laptop: tc.laptop,
			}

			server := service.NewLaptopServer(tc.store)
			res, err := server.CreateLaptop(context.Background(), req)
			if tc.code == codes.OK {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.NotEmpty(t, res.Id)
				
				if len(tc.laptop.Id) > 0 {
					require.Equal(t, tc.laptop.Id, res.Id)
				}
			} else {
				require.Error(t, err)
				require.Nil(t, res)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, tc.code, st.Code())
			}
		})
	}
}
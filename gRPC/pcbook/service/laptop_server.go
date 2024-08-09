package service

import (
	"context"
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/jopaleti/pcbook/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Laptop server is the server that provides laptop services
type LaptopServer struct {
	pb.UnimplementedLaptopServiceServer
	
	laptopStore LaptopStore
	// Store Store
}

func NewLaptopServer(laptopStore LaptopStore) *LaptopServer {
	return &LaptopServer{laptopStore:  laptopStore}	
}

func (server *LaptopServer) CreateLaptop(
	ctx context.Context,
	req *pb.CreateLaptopRequest,
) (*pb.CreateLaptopResponse, error) {
	laptop := req.GetLaptop()
	log.Printf("receive a create-laptop request with id: %s", laptop.Id)

	if len(laptop.Id) > 0 {
		// Check if it's valid UUID
		_, err := uuid.Parse(laptop.Id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "laptop Id is not a valid UUID: %v", err)
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "cannot generate a new laptop ID: %v", err)
		}
		laptop.Id = id.String()
	}
	// Performing some heavy processing to confirm our timeOut setting
	// time.Sleep(5*time.Second)

	/*
		Don't save the laptop data to the In-memory storage if the
		context time is canceled either with context timeout setting using
		context.WithTimeout or user canceled the request before the appointed time
		Example of context timeout canceled::
			Client:> context.WithTimeout(context.Background(), 4*time.Second)
			Server:> time.Sleep(5*time.Second)
			Explanation: Server sleeps for 5 seconds and client request should
			be terminated after 4 minutes, that means before server wake up the 
			client request would have terminated and when such happen, you are to 
			determine how the server should respond to such event either render it 
			useless or still make the data persistent.
	*/
	if ctx.Err() == context.Canceled {
		log.Print("request is canceled")
		return nil, status.Error(codes.Canceled, "request is canceled")
	}
	/*
		Don't save the laptop data to the In-memory storage if the
		context time is exceeded
	*/
	if ctx.Err() == context.DeadlineExceeded {
		log.Println("deadline is exceeded")
		return nil, status.Error(codes.DeadlineExceeded, "deadline is exceeded")
	}
	// Save the laptop to in memory storage
		err := server.laptopStore.Save(laptop)
		if err != nil {
			code := codes.Internal
			if errors.Is(err, ErrAlreadyExists) {
				code = codes.AlreadyExists
			}
			return nil, status.Errorf(code, "cannot save laptop to the store: %v", err)
		}

		log.Printf("saved laptop with id: %s", laptop.Id)

		res := &pb.CreateLaptopResponse{
			Id: laptop.Id,
		}
		return res, nil
}

func (server *LaptopServer) SearchLaptop(
	req *pb.SearchLaptopRequest,
	stream pb.LaptopService_SearchLaptopServer,
	) error {
		filter := req.GetFilter()
		log.Printf("receive a search-laptop request with filter: %v", filter)

		err := server.laptopStore.Search(
			stream.Context(),
			filter,
			func (laptop *pb.Laptop) error {
				res := &pb.SearchLaptopResponse{
					Laptop: laptop,
				}
				log.Print(res)
				err := stream.Send(res)
				if err != nil {
					return err
				}

				log.Printf("sent laptop with id: %s", laptop.GetId())
				return nil
			},
		)

		if err != nil {
			return status.Errorf(codes.Internal, "unexpected error: %v", err)
		}

		return nil
}
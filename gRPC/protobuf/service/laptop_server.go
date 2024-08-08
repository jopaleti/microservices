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
package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/jopaleti/pcbook/pb"
	"github.com/jopaleti/pcbook/sample"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CreateLaptop(laptopClient pb.LaptopServiceClient) {
	laptop := sample.NewLaptop()
	// laptop.Id = "ffcd0403-c975-4209-9862-e88526431047"
	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	// SETTING TIMEOUT IN GO
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	res, err := laptopClient.CreateLaptop(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			// Not a big deal
			log.Print("laptop already exists")
		} else {
			log.Fatal("cannot create laptop: ", err)
		}
		return
	}

	log.Printf("created laptop with id: %s", res.Id)

}

func SearchLaptop(laptopClient pb.LaptopServiceClient, filter *pb.Filter) {
	log.Print("search filter: ", filter)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	req := &pb.SearchLaptopRequest{Filter: filter}
	stream, err := laptopClient.SearchLaptop(ctx, req)
	if err != nil {
		log.Fatal("cannot search laptop: ", err)

	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			return
		}

		if err != nil {
			log.Fatal("cannot receive response: ", err)
		}
		fmt.Print("res")
		laptop := res.GetLaptop()
		log.Println("- found: ", laptop.GetId())
		log.Println("- brand: ", laptop.GetBrand())
		log.Println("- name: ", laptop.GetName())
		log.Println("- cpu cores: ", laptop.GetCpu().GetNumberCores())
		log.Println("- cpu min ghz: ", laptop.GetCpu().GetMinGhz())
		log.Println("- ram: ", laptop.GetRam().GetValue(), laptop.GetRam().GetUnit())
		log.Println("- price: ", laptop.GetPriceUsd(), "usd")
	}
}

func main() {
	serverAddress := flag.String("address", "", "the server address")
	flag.Parse()
	log.Printf("dial server %v", serverAddress)

	conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	laptopClient := pb.NewLaptopServiceClient(conn)

	// Create 10 random laptops with for loop
	for i := 0; i < 100; i++ {
		CreateLaptop(laptopClient)
	}

	filter := &pb.Filter {
		MaxPriceUsd: 2000,
		MinCpuCores: 4,
		MinCpuGhz: 1.5,
		MinRam: &pb.Memory{Value: 20, Unit: pb.Memory_GIGABYTE},
	}

	SearchLaptop(laptopClient, filter)
}
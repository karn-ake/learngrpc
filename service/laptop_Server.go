package service

import (
	"context"
	"errors"
	"grpc/pb"
	"log"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"

	// "google.golang.org/grpc/internal/status"
	"google.golang.org/grpc/status"
)

type LaptopServer struct {
	pb.UnimplementedLaptopServiceServer
	laptopStore LaptopStore
}

func NewLaptopServer(laptopStore LaptopStore) *LaptopServer {
	return &LaptopServer{laptopStore:laptopStore}
}

func (s *LaptopServer) CreateLaptop(ctx context.Context, req *pb.CreateLaptopRequest) (*pb.CreateLaptopResponse, error) {
	laptop := req.GetLaptop()
	log.Printf("get a create laptop request with id: %s", laptop)
	if ctx.Err() == context.Canceled {
		log.Print("request is canceled")
		return nil, status.Error(codes.Canceled, "request is canceled")
	}
	
	if ctx.Err() == context.DeadlineExceeded {
		log.Print("deadline is exceeded")
		return nil, status.Error(codes.DeadlineExceeded, "deadline is exceeded")
	}
	
	if len(laptop.Id) > 0 {
		_, err := uuid.Parse(laptop.Id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "laptop ID is not a valid UUID: %s", err)
			} else {
				id, err := uuid.NewRandom()
				if err != nil {
					return nil, status.Errorf(codes.Internal, "cannot generate new laptop ID: %v", err)
				}
				laptop.Id = id.String()
			}
		}
		
		err := s.laptopStore.Save(laptop)
		if err != nil {
			code := codes.Internal
			if errors.Is(err, ErrAlreadyExists) {
				code = codes.AlreadyExists
			}
			return nil, status.Errorf(code, "cannot save to laptop store: %v", err)
		}
		
		log.Printf("save laptop with ID: %s", laptop.Id)
		
		res := &pb.CreateLaptopResponse{
			Id: laptop.Id,
		}
		log.Printf("create a laptop id: %s", res.Id)
		return res, nil
	}

package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "grpc-user-service/grpc-user-service/proto"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

var users = []*pb.User{
	{Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
	{Id: 2, Fname: "Alice", City: "New York", Phone: 2345678901, Height: 5.5, Married: false},
	{Id: 3, Fname: "Bob", City: "Chicago", Phone: 3456789012, Height: 6.0, Married: true},
	{Id: 4, Fname: "Carol", City: "San Francisco", Phone: 4567890123, Height: 5.6, Married: false},
	{Id: 5, Fname: "Dave", City: "Seattle", Phone: 5678901234, Height: 5.9, Married: true},
	{Id: 6, Fname: "Eve", City: "Austin", Phone: 6789012345, Height: 5.7, Married: false},
	{Id: 7, Fname: "Frank", City: "Boston", Phone: 7890123456, Height: 6.1, Married: true},
	{Id: 8, Fname: "Grace", City: "Miami", Phone: 8901234567, Height: 5.4, Married: false},
	{Id: 9, Fname: "Hank", City: "Denver", Phone: 9012345678, Height: 6.2, Married: true},
	{Id: 10, Fname: "Ivy", City: "Dallas", Phone: 1234509876, Height: 5.3, Married: false},
}

func (s *server) GetUserDetails(ctx context.Context, req *pb.UserIdRequest) (*pb.UserDetailsResponse, error) {
	for _, user := range users {
		if user.Id == req.Id {
			// Create a copy of the user before returning
			return &pb.UserDetailsResponse{User: user}, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "User not found")
}

func (s *server) GetUsersByIds(ctx context.Context, req *pb.UserIdsRequest) (*pb.UsersDetailsResponse, error) {
	var result []*pb.User
	for _, id := range req.Ids {
		for _, user := range users {
			if user.Id == id {
				// Create a copy of the user before appending
				result = append(result, user)
			}
		}
	}
	return &pb.UsersDetailsResponse{Users: result}, nil
}

func (s *server) SearchUsers(ctx context.Context, req *pb.SearchRequest) (*pb.UsersDetailsResponse, error) {
	var result []*pb.User
	for _, user := range users {
		if (req.Fname == "" || user.Fname == req.Fname) &&
			(req.City == "" || user.City == req.City) &&
			(!req.Married || user.Married == req.Married) {
			// Create a copy of the user before appending
			result = append(result, user)
		}
	}
	return &pb.UsersDetailsResponse{Users: result}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

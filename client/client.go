package main

import (
    "context"
    "log"
    "time"

    "google.golang.org/grpc"
    pb "grpc-user-service/grpc-user-service/proto"
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewUserServiceClient(conn)

    // GetUserDetails
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()
    r, err := c.GetUserDetails(ctx, &pb.UserIdRequest{Id: 1})
    if err != nil {
        log.Fatalf("could not get user details: %v", err)
    }
    log.Printf("User details: %v", r.GetUser())

    // GetUsersByIds
    ctx, cancel = context.WithTimeout(context.Background(), time.Second)
    defer cancel()
    r2, err := c.GetUsersByIds(ctx, &pb.UserIdsRequest{Ids: []int32{1, 2}})
    if err != nil {
        log.Fatalf("could not get users by IDs: %v", err)
    }
    log.Printf("Users by IDs: %v", r2.GetUsers())

    // SearchUsers
    ctx, cancel = context.WithTimeout(context.Background(), time.Second)
    defer cancel()
    r3, err := c.SearchUsers(ctx, &pb.SearchRequest{Fname: "Alice"})
    if err != nil {
        log.Fatalf("could not search users: %v", err)
    }
    log.Printf("Search results: %v", r3.GetUsers())
}


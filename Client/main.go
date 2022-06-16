package main

import (
	pb "GitHub.com/mhthrh/JWT/usermgmt"
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address = "localhost:9999"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SignIn(ctx, &pb.Login{
		Username: "mohsen",
		Password: "Qaz@123456",
	})
	if err != nil {
		log.Fatalf("could not create user: %v", err)
	}
	log.Printf("Umtil: %s key: %s", r.ValidTill, r.SignedKey)

}

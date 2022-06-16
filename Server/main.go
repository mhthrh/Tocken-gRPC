package main

import (
	"GitHub.com/mhthrh/JWT/Server/User"
	pb "GitHub.com/mhthrh/JWT/usermgmt"
	"context"
	"crypto/sha256"
	"flag"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net"
)

var (
	ip   string
	port int
	cnn  = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"127.0.0.1", 5432, "postgres", "123456", "MyDb")
)

type UserManagementServer struct {
	pb.UnimplementedUserManagementServer
	db *gorm.DB
	f  func(string) string
}

func Sha256(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}
func (s *UserManagementServer) SignIn(ctx context.Context, in *pb.Login) (*pb.JWT, error) {
	var l User.Request

	l.Username = in.Username
	l.Password = in.Password
	response, err := User.New(s.db, s.f).SignIn(&l)
	if err != nil {
		return nil, err
	}
	res := &pb.JWT{SignedKey: response.ValidTill, ValidTill: response.SignedKey}

	return res, nil
}
func main() {
	flag.StringVar(&ip, "ip", "0.0.0.0", "What is listener IP address?")
	flag.IntVar(&port, "port", 9999, "What is the port?")
	flag.Parse()

	db, err := gorm.Open(postgres.Open(cnn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	err = db.AutoMigrate(User.User{})
	db.Create(&User.User{
		Id:          uuid.UUID{},
		Name:        "mohsen",
		LastName:    "taheri",
		UserName:    "mohsen",
		PassWord:    Sha256("Qaz@123456"),
		Email:       "taheri.mo@outlook.com",
		PhoneNumber: "+447759448882",
		IsActive:    true,
		CreatDate:   0,
	})
	if err != nil {
		fmt.Println(err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserManagementServer(s, &UserManagementServer{
		UnimplementedUserManagementServer: pb.UnimplementedUserManagementServer{},
		db:                                db,
		f:                                 Sha256,
	})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

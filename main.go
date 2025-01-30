package main

import (
	"context"
	"log"
	"net"

	"github.com/izzzicos/grpc-with-golang/ticketing"
	"google.golang.org/grpc"
)

type ticketingServer struct {
	ticketing.UnimplementedTicketServer
}

func (s ticketingServer) Create(context.Context, *ticketing.CreateRequest) (*ticketing.CreateResponse, error) {
	return &ticketing.CreateResponse{
		Doctor: &ticketing.Doctor{
			Name: "John",
			LastName: "Doe",
			Service: "GP",
		},
		OfficeNumber: 104,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err !=nil {
		log.Fatalf("Cannot create listener: %s", err)
	}

	serverRegistrar := grpc.NewServer()
	service := &ticketingServer{}

	ticketing.RegisterTicketServer(serverRegistrar, service)

	err = serverRegistrar.Serve(lis)
	if err !=nil {
		log.Fatalf("Cannot serve: %s", err)
	}
}
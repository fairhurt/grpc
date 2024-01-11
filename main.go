package main

import (
	"context"
	"log"
	"net"

	"github.com/fairhurt/demo-grpc/invoicer"
	"google.golang.org/grpc"
)

type myInvoicerServer struct {
	invoicer.UnimplementedInvoicerServer
}

//	func (s *myInvoicerServer) Create(context.Context, *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
//		return &invoicer.CreateResponse{Pdf: []byte("Test"), Docx: []byte("Test2")}, nil
//	}
func (s myInvoicerServer) Create(ctx context.Context, req *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
	return &invoicer.CreateResponse{
		Pdf:           []byte(req.From),
		Docx:          []byte("test"),
		InvoiceNumber: "124",
	}, nil
}

func main() {
	log.Default().Println("Starting server...")
	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("Failed to listen on port 8089: %s", err)
		panic(err)
	}
	serverRegistrar := grpc.NewServer()
	log.Default().Println("Registering server...")
	service := &myInvoicerServer{}
	log.Default().Println("Registering service...")
	invoicer.RegisterInvoicerServer(serverRegistrar, service)
	log.Default().Println("Serving...")
	err = serverRegistrar.Serve(lis)
	log.Default().Println("Server started")
	if err != nil {
		log.Fatalf("Failed to serve gRPC server over port 8089: %s", err)
	}
}

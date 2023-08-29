package main

import (
	"context"
	"log"
	"net"

	goproject3 "example.com/goproject3/invoicer"
	protoFiles "example.com/goproject3/protoFiles"
	"google.golang.org/grpc"
)

type myInvoicerServer struct {
	goproject3.UnimplementedInvoicerServer
}

type myBookServer struct {
	protoFiles.UnimplementedBook1Server
}
type myAuthorServer struct {
	protoFiles.Auhtor1Server
}

func (server *myInvoicerServer) Create(context.Context, *goproject3.CreateRequest) (*goproject3.CreateResponse, error) {
	return &goproject3.CreateResponse{
		Pdf:  []byte("test1"),
		Docx: []byte("test2"),
		Exe:  []byte("test3"),
	}, nil
}

func (server *myBookServer) Book1CreateService(ctx context.Context, request *protoFiles.BookCreateRequest1) (*protoFiles.BookCreateResponse1, error) {
	return &protoFiles.BookCreateResponse1{
		BookId:    1,
		BookName:  request.BookName,
		BookPrice: request.BookPrice,
	}, nil
}
func (server *myAuthorServer) Author1CreateService(ctx context.Context, request *protoFiles.AuthorCreateRequest1) (*protoFiles.AuthorCreateResponse1, error) {
	return &protoFiles.AuthorCreateResponse1{
		AuthorId:   1,
		AuthorName: request.AuthorName,
		AuthorAge:  request.AuthorAge,
	}, nil
}

func main() {
	var listener, err = net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("Cannot create listener: %s", err)
	}
	serverRegisterer := grpc.NewServer()
	service := &myInvoicerServer{}
	goproject3.RegisterInvoicerServer(serverRegisterer, service)
	service2 := &myBookServer{}
	protoFiles.RegisterBook1Server(serverRegisterer, service2)
	service3 := &myAuthorServer{}
	protoFiles.RegisterAuhtor1Server(serverRegisterer, service3)
	err = serverRegisterer.Serve(listener)
	if err != nil {
		log.Fatalf("Impossible to serve: %s", err)
	}
}

package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/nordicdyno/grpc-listen/api"
	
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func twoPorts(ctx context.Context) {
	grpcListener, err := net.Listen("tcp", "localhost:5554")
	if err != nil {
		panic(err)
	}
	fmt.Printf("grpc addr: %s\n", grpcListener.Addr())
	gwListener, err := net.Listen("tcp", "localhost:5555")
	if err != nil {
		panic(err)
	}
	fmt.Printf("gateway addr: %s\n", grpcListener.Addr())

	// grpcMux := runtime.NewServeMux()
	gwMux := runtime.NewServeMux()

	// mux.HandleFunc("/swagger.json", func(w http.ResponseWriter, req *http.Request) {
	// 	io.Copy(w, strings.NewReader(pb.Swagger))
	// })

	runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: false})
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}
	err = api.RegisterPublisherHandlerFromEndpoint(
		context.Background(), gwMux, grpcListener.Addr().String(), dialOpts)
	if err != nil {
		panic(errors.Wrap(err, "failed register handler from endpoint"))
	}
	gwServer := &http.Server{
		Handler: gwMux,
	}

	// mux.Handle("/", gwmux)

	// add serve swagger here

	// opts := []grpc.ServerOption{}
	// grpcServer := grpc.NewServer(opts...)
	grpcServer := grpc.NewServer()
	api.RegisterPublisherServer(grpcServer, &publisherServer{})
	reflection.Register(grpcServer)

	fin := make(chan struct{})
	// graceful shutdown
	go func() {
		<-ctx.Done()

		log.Println("shutting down gateway server...")
		gwServer.Shutdown(context.Background())

		log.Println("shutting down gRPC server...")
		grpcServer.GracefulStop()

		close(fin)
	}()

	go grpcServer.Serve(grpcListener)
	go gwServer.Serve(gwListener)
	<-fin
}

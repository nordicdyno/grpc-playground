package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/nordicdyno/grpc-listen/api"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// based on https://github.com/philips/grpc-gateway-example/issues/22

func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		contentType := r.Header.Get("Content-Type")
		log.Printf("grpcHandlerFunc: r.ProtoMajor=%v, method=%v, contentType=%v\n",
			r.ProtoMajor, r.Method, contentType)

		if r.ProtoMajor == 2 && strings.Contains(contentType, "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}

func onePort(ctx context.Context) {
	l, err := net.Listen("tcp", "localhost:5556")
	if err != nil {
		panic(err)
	}
	fmt.Printf("common addr: %s\n", l.Addr())

	grpcServer := grpc.NewServer()
	api.RegisterPublisherServer(grpcServer, &publisherServer{})
	reflection.Register(grpcServer)

	mux := http.NewServeMux()
	gwmux := runtime.NewServeMux()

	mux.Handle("/", gwmux)

	dopts := []grpc.DialOption{grpc.WithInsecure()}
	err = api.RegisterPublisherHandlerFromEndpoint(ctx, gwmux, l.Addr().String(), dopts)

	srv := &http.Server{
		Handler: grpcHandlerFunc(grpcServer, mux),
	}

	fin := make(chan struct{})
	go func() {
		<-ctx.Done()
		log.Println("shutting down shared server...")
		srv.Shutdown(ctx)
		close(fin)
	}()
	go func() {
		err = srv.Serve(l)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}()
	<-fin
}

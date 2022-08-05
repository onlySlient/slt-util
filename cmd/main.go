package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/EDDYCJY/go-grpc-example/pkg/gtls"
	"google.golang.org/grpc"

	pb "slient.util/generate/proto"
	"slient.util/service"
)

const (
	port = ":9000"

	certFile = "static/cert/server.crt"
	keyFile  = "static/cert/server.key"
)

func main() {
	tlsServer := gtls.Server{
		CertFile: certFile,
		KeyFile:  keyFile,
	}

	c, err := tlsServer.GetTLSCredentials()
	if err != nil {
		panic(err)
	}

	grpcS := grpc.NewServer(grpc.Creds(c))
	pb.RegisterSltUtilServiceServer(grpcS, service.NewServer())

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello: grpc http"))
	})

	log.Println("server start ", port)

	if err = http.ListenAndServeTLS(port, certFile, keyFile, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcS.ServeHTTP(w, r)
		} else {
			mux.ServeHTTP(w, r)
		}
	})); err != nil {
		panic(err)
	}
}

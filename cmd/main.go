package main

import (
	"net/http"
	"strings"

	"github.com/EDDYCJY/go-grpc-example/pkg/gtls"
	"google.golang.org/grpc"

	pb "slient.util/generate/proto"
	"slient.util/service"
)

var (
	certFile = ""
	keyFile  = ""
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

	// TODO create http server by mux

	if err = http.ListenAndServeTLS(":8000", certFile, keyFile, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcS.ServeHTTP(w, r)
		} else {
			// TODO
			// mux.ServeHTTP(w, r)
		}

	})); err != nil {
		panic(err)
	}
}

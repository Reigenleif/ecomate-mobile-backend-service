package main

import (
	"context"
	"flag"
	"log"
	"net"

	go_grpc_gauth_pg "github.com/Reigenleif/go-grpc-gauth-pg/api"
	"github.com/Reigenleif/go-grpc-gauth-pg/internal/db"
	"github.com/Reigenleif/go-grpc-gauth-pg/service"
	_ "github.com/jackc/pgx/v4/stdlib"
	"google.golang.org/grpc"
)

var (
	// tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	// certFile   = flag.String("cert_file", "", "The TLS cert file")
	// keyFile    = flag.String("key_file", "", "The TLS key file")
	// jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features")
	port       = flag.Int("port", 50051, "The server port")
)


func main() {
	// Initiate database connection
	_, err := db.ConnectSQL()
	if err != nil {
		panic(err)
	}

	// testping
	r, err := db.GetDB().Query(context.Background(), "SELECT * FROM public.\"Emojis\"")
	if err != nil {
		panic(err)
	}

	defer r.Close()

	// Initiate gRPC server
	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}

	serverOpts := []grpc.ServerOption{
		
	}

	serverRegistrar := grpc.NewServer(serverOpts...)
	emojiService := &service.EmojiService{}

	
	go_grpc_gauth_pg.RegisterEmojiServer(serverRegistrar, emojiService)
	log.Print("Server started on port 8089")
	err = serverRegistrar.Serve(lis)
	if err != nil {
		log.Fatalf("impossible to serve: %s", err)
	}
	log.Print("Server started on port 8089")
}


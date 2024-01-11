package main

import (
	"flag"
	"log"
	"net"

	proto"github.com/Reigenleif/ecomate-mobile-backend-service/proto"
	"github.com/Reigenleif/ecomate-mobile-backend-service/internal/db"
	"github.com/Reigenleif/ecomate-mobile-backend-service/service"
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

	// Initiate gRPC server
	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}

	serverOpts := []grpc.ServerOption{
		
	}

	serverRegistrar := grpc.NewServer(serverOpts...)
	proto.RegisterNewsServiceServer(serverRegistrar, &service.NewsService{})
	proto.RegisterAuthServiceServer(serverRegistrar, &service.AuthService{})
	proto.RegisterMarketplaceServer(serverRegistrar, &service.MarketplaceService{})

	log.Print("Server started on port 8089")
	err = serverRegistrar.Serve(lis)
	if err != nil {
		log.Fatalf("impossible to serve: %s", err)
	}
	log.Print("Server started on port 8089")
}


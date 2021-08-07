package controllers

import (
	"context"
	"fmt"
	"log"
	"net"
	// "time"
	"os"

	"google.golang.org/grpc"
	"github.com/jackc/pgx/v4/pgxpool"

	pb "github.com/bartmika/growlog-server/proto"
	"github.com/bartmika/growlog-server/internal/repositories"
)

type Controller struct {
	port               int
	databaseUrl        string
	dbpool             *pgxpool.Pool
	grpcServer         *grpc.Server
}

func New(port int, databaseUrl string) *Controller {
	dbpool, err := pgxpool.Connect(context.Background(), databaseUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return &Controller{
		port:               port,
		databaseUrl:        databaseUrl,
		dbpool:             dbpool,
		grpcServer:         nil,
	}
}

// Function will consume the main runtime loop and run the business logic
// of the application.
func (s *Controller) RunMainRuntimeLoop() {
	// Open a TCP server to the specified localhost and environment variable
	// specified port number.
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", s.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Initialize our gRPC server using our TCP server.
	grpcServer := grpc.NewServer()

	tenantRepo := repositories.NewTenantRepo(s.dbpool)

	// Save reference to our application state.
	s.grpcServer = grpcServer

	// For debugging purposes only.
	log.Printf("gRPC server is running on port %v", s.port)

	// Block the main runtime loop for accepting and processing gRPC requests.
	pb.RegisterGrowLogServer(grpcServer, &ControllerImpl{
		// DEVELOPERS NOTE:
		// We want to attach to every gRPC call the following variables...
		tenantRepo: tenantRepo,
	})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// Function will tell the application to stop the main runtime loop when
// the process has been finished.
func (s *Controller) StopMainRuntimeLoop() {
	log.Printf("Starting graceful shutdown now...")

	// Finish our database operations running.
	// s.storage.Close()

	defer s.dbpool.Close()

	// Finish any RPC communication taking place at the moment before
	// shutting down the gRPC server.
	s.grpcServer.GracefulStop()
}

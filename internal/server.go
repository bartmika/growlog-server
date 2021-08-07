package internal

import (
	"fmt"
	"log"
	"net"
	// "time"

	"google.golang.org/grpc"

	pb "github.com/bartmika/growlog-server/proto"
)

type GrowLogServer struct {
	port               int
	grpcServer         *grpc.Server
}

func New(port int) *GrowLogServer {

	return &GrowLogServer{
		port:               port,
		grpcServer:         nil,
	}
}

// Function will consume the main runtime loop and run the business logic
// of the application.
func (s *GrowLogServer) RunMainRuntimeLoop() {
	// Open a TCP server to the specified localhost and environment variable
	// specified port number.
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", s.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Initialize our gRPC server using our TCP server.
	grpcServer := grpc.NewServer()



	// Save reference to our application state.
	s.grpcServer = grpcServer

	// For debugging purposes only.
	log.Printf("gRPC server is running on port %v", s.port)

	// Block the main runtime loop for accepting and processing gRPC requests.
	pb.RegisterGrowLogServer(grpcServer, &GrowLogServerImpl{
		// DEVELOPERS NOTE:
		// We want to attach to every gRPC call the following variables...
		// storage: s.storage,
	})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// Function will tell the application to stop the main runtime loop when
// the process has been finished.
func (s *GrowLogServer) StopMainRuntimeLoop() {
	log.Printf("Starting graceful shutdown now...")

	// Finish our database operations running.
	// s.storage.Close()

	// Finish any RPC communication taking place at the moment before
	// shutting down the gRPC server.
	s.grpcServer.GracefulStop()
}

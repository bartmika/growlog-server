package cmd

import (
	// "log"
	"os"
	"os/signal"
	"syscall"
	// "time"

	"github.com/spf13/cobra"

	"github.com/bartmika/growlog-server/internal/controllers"
	// "github.com/bartmika/growlog-server/utils"
)

var (
	port                     int
	databaseUrl              string
)

func init() {
	// The following are optional and will have defaults placed when missing.
	serveCmd.Flags().IntVarP(&port, "port", "p", 50051, "The port to run this server on")
	serveCmd.Flags().StringVarP(&databaseUrl, "databaseUrl", "d", os.Getenv("GROWLOG_DATABASE_URL"), "The databaseUrl to run this server on")

	// Make this sub-command part of our application.
	rootCmd.AddCommand(serveCmd)
}

func doServe() {
	// Convert the user inputted integer value to be a `time.Duration` type.

	// Setup our server.
	server := controllers.New(port, databaseUrl)

	// DEVELOPERS CODE:
	// The following code will create an anonymous goroutine which will have a
	// blocking chan `sigs`. This blocking chan will only unblock when the
	// golang app receives a termination command; therfore the anyomous
	// goroutine will run and terminate our running application.
	//
	// Special Thanks:
	// (1) https://gobyexample.com/signals
	// (2) https://guzalexander.com/2017/05/31/gracefully-exit-server-in-go.html
	//
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs // Block execution until signal from terminal gets triggered here.
		server.StopMainRuntimeLoop()
	}()
	server.RunMainRuntimeLoop()
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run the gRPC server",
	Long:  `Run the gRPC server to allow other services to access the application`,
	Run: func(cmd *cobra.Command, args []string) {
		// Defensive code. ...
		// Do nothing for now...

		// Execute our command with our validated inputs.
		doServe()
	},
}

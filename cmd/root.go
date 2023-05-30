package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"proto-service/cmd/generate"
	"syscall"
)

func Start() {
	rootCmd := &cobra.Command{}

	_, cancel := context.WithCancel(context.Background())

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-quit
		cancel()
	}()

	generateCmd := &cobra.Command{
		Use:   "generate:grpc",
		Short: "generate grpc proto file",
		Run: func(cmd *cobra.Command, args []string) {
			generate.GenerateProto(args)
		},
	}

	generateCmd.Flags().BoolP("version", "", false, "print version")
	generateCmd.Flags().BoolP("verbose", "", false, "enable verbose mode")
	generateCmd.Flags().BoolP("guide", "", false, "print help")

	cmd := []*cobra.Command{
		//{
		//	Use:   "generate:grpc",
		//	Short: "Run HTTP Server",
		//	Run: func(cmd *cobra.Command, args []string) {
		//		fmt.Println("grpc server running")
		//	},
		//},
		generateCmd,
	}

	rootCmd.AddCommand(cmd...)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while execute yout cli %s", err)
		os.Exit(1)
	}
}

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "heart_beat",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root")
	},
}

func init() {
	RootCmd.AddCommand(ServerCmd)
	RootCmd.AddCommand(ClientCmd)
}

func Execute() error {
	return RootCmd.Execute()
}

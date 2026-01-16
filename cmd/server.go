package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Heart beat server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server")
	},
}

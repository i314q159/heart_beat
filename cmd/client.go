package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ClientCmd = &cobra.Command{
	Use:   "client",
	Short: "Heart beat client",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cient")
	},
}

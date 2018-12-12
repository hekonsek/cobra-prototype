package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(provisionCmd)
}

var provisionCmd = &cobra.Command{
	Use: "foo",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is sub-command!")
	},
}

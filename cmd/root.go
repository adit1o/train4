package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// [ ] root command
var rootCmd = &cobra.Command{
	Use:   "javit",
	Short: "Matches JAV videos and normalizes metadata.",

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "testcmd",
	Args: cobra.NoArgs,
	RunE: mainCommand,
}

func init() {
}

func main() {
	_ = rootCmd.Execute()
}

func mainCommand(cmd *cobra.Command, args []string) error {
	var homeDir string
	{
		var err error
		homeDir, err = os.UserHomeDir()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			os.Exit(1)
		}
	}

	fmt.Fprintf(os.Stdout, "Home dir: %s", homeDir)
	return nil
}

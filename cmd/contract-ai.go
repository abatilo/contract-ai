package main

import (
	"github.com/abatilo/contract-ai/cmd/api"
	"github.com/spf13/cobra"
)

func main() {
	var (
		rootCmd = &cobra.Command{
			Use:   "contract-ai",
			Short: "Read a document and ask questions about it",
		}
	)

	rootCmd.AddCommand(api.Cmd)
	rootCmd.Execute()
}

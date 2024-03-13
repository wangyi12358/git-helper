package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func Execute() {
	var rootCmd = &cobra.Command{
		Use:   "git-helper",
		Short: "A helper for git",
	}
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

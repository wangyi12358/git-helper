package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wangyi12358/git-helper/cmd/commit"
	"github.com/wangyi12358/git-helper/cmd/help"
	"os"
)

func Execute() {
	var rootCmd = &cobra.Command{
		Use:   "git-helper",
		Short: "A helper for git",
	}
	rootCmd.AddCommand(help.Cmd)
	rootCmd.AddCommand(commit.Cmd)
	var commitType string
	commit.Cmd.PersistentFlags().StringVarP(&commitType, "type", "t", "", "commit type")
	fmt.Printf("commit type: %s\n", commitType)
	rootCmd.Flags()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

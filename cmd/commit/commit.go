package commit

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "commit",
	Short: "commit",
	Run:   CommitCmd,
}

func CommitCmd(cmd *cobra.Command, _ []string) {
	//cmd.PersistentFlags().StringVarP(&commitType, "type", "t", "", "commit type")
	//fmt.Printf("commit type: %s\n", commitType)
}

package help

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "help",
	Short: "Help about any command",
}

func HelpCmd(_ *cobra.Command, _ []string) {

}

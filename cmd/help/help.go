package help

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"os"
)

var Cmd = &cobra.Command{
	Use:   "help",
	Short: "Help about any command",
	Run:   HelpCmd,
}

func HelpCmd(_ *cobra.Command, _ []string) {
	var title string
	err := survey.AskOne(&survey.Select{
		Message: "请选择命令",
		Options: Titles,
	}, &title)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(0)
	}
	for _, command := range Commands {
		if command.Title == title {
			fmt.Println(command.Short)
			break
		}
	}
}

package cmdmodelset

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var diffLong = `show the current diff for a modelset`

func DiffRun(name string) (err error) {

	return err
}

var DiffCmd = &cobra.Command{

	Use: "diff",

	Short: "show the current diff for a modelset",

	Long: diffLong,

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		if 0 >= len(args) {
			fmt.Println("missing required argument: 'Name'")
			cmd.Usage()
			os.Exit(1)
		}

		var name string

		if 0 < len(args) {

			name = args[0]

		}

		err = DiffRun(name)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
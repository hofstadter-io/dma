package cmdstore

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var connLong = `connect to the local datastore`

func ConnRun(name string) (err error) {

	return err
}

var ConnCmd = &cobra.Command{

	Use: "conn",

	Short: "connect to the local datastore",

	Long: connLong,

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

		err = ConnRun(name)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
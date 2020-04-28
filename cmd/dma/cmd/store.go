package cmd

import (
	"github.com/spf13/cobra"

	"github.com/hofstadter-io/dma/cmd/dma/cmd/store"
)

var storeLong = `create, checkpoint, and migrate your storage engines`

var StoreCmd = &cobra.Command{

	Use: "store",

	Aliases: []string{
		"s",
	},

	Short: "create, checkpoint, and migrate your storage engines",

	Long: storeLong,
}

func init() {
	StoreCmd.AddCommand(cmdstore.RunCmd)
	StoreCmd.AddCommand(cmdstore.ConnCmd)
}
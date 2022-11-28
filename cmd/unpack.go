package cmd

import "github.com/spf13/cobra"

var unpackCmd = &cobra.Command{
	Use:   "unpack",
	Short: "Unpack file using",
}

func init() {
	rootCmd.AddCommand(unpackCmd)
}

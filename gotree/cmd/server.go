package cmd

import "github.com/spf13/cobra"

var server *cobra.Command

func init() {
	server = &cobra.Command{
		Use: "server",
	}

	root.AddCommand(server)
}

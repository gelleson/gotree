package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var root *cobra.Command

func init() {
	root = &cobra.Command{
		Use: "root",
	}
}

func Execute() {
	err := root.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	loginCmd = &cobra.Command{
		Use:   "login [token]",
		Short: "Logs in to Thingiverse, or saves one to config if provided",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("Login flow not yet implemented.")
				fmt.Println("Until then:")
				fmt.Println("  Create a Desktop app from https://www.thingiverse.com/apps/create and then")
				fmt.Println("  Copy the read-only app token")
				fmt.Println("then run: thing <command> --token <token>")
				fmt.Println("or to save it: thing login <token>")
			} else {
				fmt.Println("[TODO] Token saved. The --token flag is no longer required.")
			}
		},
	}
)

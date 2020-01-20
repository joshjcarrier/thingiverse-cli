package cmd

import (
	"fmt"
	"os"

	"github.com/joshjcarrier/thingiverse-cli/pkg/render"

	api "github.com/joshjcarrier/thingiverse-cli/pkg/api/thingiverse"
	"github.com/spf13/cobra"
)

var (
	searchCmd = &cobra.Command{
		Use:   "search term ...",
		Short: "Search for things by a term",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			response, _ := api.SearchThingsByTerm(args, options.API)
			err := render.Execute(os.Stdout, response, *options.Render)
			if err != nil {
				fmt.Println(err)
			}
		},
	}
)

func init() {
	api.InstallAPICommandFlags(searchCmd, options.API)
	render.InstallCommandFlags(searchCmd, options.Render)
}

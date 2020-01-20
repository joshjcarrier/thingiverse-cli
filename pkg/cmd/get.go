package cmd

import (
	"fmt"
	"os"

	api "github.com/joshjcarrier/thingiverse-cli/pkg/api/thingiverse"
	"github.com/joshjcarrier/thingiverse-cli/pkg/render"

	"github.com/spf13/cobra"
)

var (
	getCmd = &cobra.Command{
		Use:   "get id",
		Short: "Gets a thing by its id",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			response, _ := api.GetThingByID(args[0], options.API)
			err := render.Execute(os.Stdout, response, *options.Render)
			if err != nil {
				fmt.Println(err)
			}
		},
	}
)

func init() {
	api.InstallAPICommandFlags(getCmd, options.API)
	render.InstallCommandFlags(getCmd, options.Render)
}

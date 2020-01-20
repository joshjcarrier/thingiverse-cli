package cmd

import (
	"fmt"
	"os"

	"github.com/joshjcarrier/thingiverse-cli/pkg/render"

	Thingiverse "github.com/joshjcarrier/thingiverse-cli/pkg/api/thingiverse"
	api "github.com/joshjcarrier/thingiverse-cli/pkg/api/thingiverse"
	"github.com/spf13/cobra"
)

var (
	userCmd = &cobra.Command{
		Use:   "user username",
		Short: "Gets Things for a given user",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			response, _ := Thingiverse.ListThingsByUsername(args[0], options.API)
			err := render.Execute(os.Stdout, response, *options.Render)
			if err != nil {
				fmt.Println(err)
			}
		},
	}
)

func init() {
	api.InstallAPICommandFlags(userCmd, options.API)
	render.InstallCommandFlags(userCmd, options.Render)
}

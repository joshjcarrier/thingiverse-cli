package cmd

import (
	"fmt"

	"github.com/joshjcarrier/thingiverse-cli/pkg/render"
	"github.com/joshjcarrier/thingiverse-cli/pkg/writer"

	api "github.com/joshjcarrier/thingiverse-cli/pkg/api/thingiverse"
	"github.com/spf13/cobra"
)

var (
	userCmd = &cobra.Command{
		Use:   "user username",
		Short: "Gets Things for a given user",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			t, _ := api.ListThingsByUsername(args[0], options.API)
			w, _ := writer.NewWriter(t, options.Render.Format, options.Writer)
			err := render.Execute(w, t, options.Render)
			if err != nil {
				fmt.Println(err)
			}
		},
	}
)

func init() {
	api.InstallAPICommandFlags(userCmd, options.API)
	render.InstallCommandFlags(userCmd, options.Render)
	writer.InstallWriterCommandFlags(userCmd, options.Writer)
}

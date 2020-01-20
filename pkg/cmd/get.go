package cmd

import (
	"fmt"

	api "github.com/joshjcarrier/thingiverse-cli/pkg/api/thingiverse"
	"github.com/joshjcarrier/thingiverse-cli/pkg/render"
	"github.com/joshjcarrier/thingiverse-cli/pkg/writer"

	"github.com/spf13/cobra"
)

var (
	getCmd = &cobra.Command{
		Use:   "get id",
		Short: "Gets a thing by its id",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			t, _ := api.GetThingByID(args[0], options.API)
			w, _ := writer.NewWriter(t, options.Render.Format, options.Writer)
			err := render.Execute(w, t, options.Render)
			if err != nil {
				fmt.Println(err)
			}
		},
	}
)

func init() {
	api.InstallAPICommandFlags(getCmd, options.API)
	render.InstallCommandFlags(getCmd, options.Render)
	writer.InstallWriterCommandFlags(getCmd, options.Writer)
}

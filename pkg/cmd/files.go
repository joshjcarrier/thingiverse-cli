package cmd

import (
	"fmt"

	api "github.com/joshjcarrier/thingiverse-cli/pkg/api/thingiverse"
	"github.com/joshjcarrier/thingiverse-cli/pkg/render"
	"github.com/joshjcarrier/thingiverse-cli/pkg/writer"

	"github.com/spf13/cobra"
)

var (
	filesCmd = &cobra.Command{
		Use:   "files thing-id",
		Short: "Gets a thing's files by its thing id",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			t, _ := api.GetThingFilesByThingID(args[0], options.API)
			w, _ := writer.NewWriter(t, options.Render.Format, options.Writer)
			err := render.Execute(w, t, options.Render)
			if err != nil {
				fmt.Println(err)
			}
		},
	}
)

func init() {
	api.InstallAPICommandFlags(filesCmd, options.API)
	render.InstallCommandFlags(filesCmd, options.Render)
	writer.InstallWriterCommandFlags(filesCmd, options.Writer)
}

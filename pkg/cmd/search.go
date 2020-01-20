package cmd

import (
	"fmt"

	"github.com/joshjcarrier/thingiverse-cli/pkg/render"
	"github.com/joshjcarrier/thingiverse-cli/pkg/writer"

	api "github.com/joshjcarrier/thingiverse-cli/pkg/api/thingiverse"
	"github.com/spf13/cobra"
)

var (
	searchCmd = &cobra.Command{
		Use:   "search term ...",
		Short: "Search for things by a term",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			t, _ := api.SearchThingsByTerm(args, options.API)
			w, _ := writer.NewWriter(t, options.Render.Format, options.Writer)
			err := render.Execute(w, t, options.Render)
			if err != nil {
				fmt.Println(err)
			}
		},
	}
)

func init() {
	api.InstallAPICommandFlags(searchCmd, options.API)
	render.InstallCommandFlags(searchCmd, options.Render)
	writer.InstallWriterCommandFlags(searchCmd, options.Writer)
}

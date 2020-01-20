package render

import (
	"fmt"
	"io"
	textTemplate "text/template"

	api "github.com/joshjcarrier/thingiverse-cli/pkg/api/thingiverse"
	"github.com/spf13/cobra"
)

var (
	fmtMap = map[string]*format{
		"md": &format{
			item: item_mdTemplate(),
			list: list_mdTemplate(),
			funcMap: map[string]interface{}{
				"test": MD,
			},
		},
		"txt": &format{
			item: item_txtTemplate(),
			list: list_txtTemplate(),
			funcMap: map[string]interface{}{
				"test": MD,
			},
		},
	}
	funcMap = map[string]interface{}{
		"customSections": func(t api.Thing) interface{} {
			customSections := []api.DetailsSectionDataItem{}
			for _, v := range *t.DetailsParts {
				if v.Type == "custom" && v.Data != nil && len(*v.Data) > 0 {
					fmt.Println("YES")
					customSections = append(customSections, (*v.Data)[0])
				}
			}
			return customSections
		},
		"printSettings": func(t api.Thing) interface{} {
			for _, v := range *t.DetailsParts {
				if v.Type == "settings" && v.Data != nil && len(*v.Data) > 0 {
					return (*v.Data)[0]
				}
			}
			return nil
		},
	}
)

// Options for different templates
type Options struct {
	Format     string
	LongFormat bool
}

// Execute renders data to the writer using a specified format.
func Execute(w io.Writer, data interface{}, opts Options) error {
	f := *fmtMap[opts.Format]

	funcs := make(map[string]interface{}, len(funcMap))
	for k, v := range funcMap {
		funcs[k] = v
	}
	for k, v := range f.funcMap {
		funcs[k] = v
	}

	t, _ := textTemplate.New("").Funcs(funcs).Parse(f.template(data))
	return t.Execute(w, struct {
		Data    interface{}
		Options interface{}
	}{Data: data, Options: opts})
}

// InstallCommandFlags enables a command to read and write config options for rendering.
func InstallCommandFlags(c *cobra.Command, opts *Options) {
	c.Flags().StringVarP(&opts.Format, "format", "f", opts.Format, "Output format: \"md\", \"txt\".")
	c.Flags().BoolVarP(&opts.LongFormat, "long-format", "l", opts.LongFormat, "List in long format, if applicable.")
}

type format struct {
	item    string
	list    string
	funcMap map[string]interface{}
}

func (t format) template(v interface{}) string {
	switch v.(type) {
	case api.Thing:
		return t.item
	case api.Things:
		return t.list
	default:
		return "NOT SUPPORTED"
	}
}

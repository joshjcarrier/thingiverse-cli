package writer

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	api "github.com/joshjcarrier/thingiverse-cli/pkg/api/thingiverse"
	"github.com/spf13/cobra"
)

const (
	// StdoutOutputDirectory is a special directory path that denotes using STDOUT
	StdoutOutputDirectory = "-"
)

// Options for writing rendering.
type Options struct {
	KeepFiles       bool
	OutputDirectory string
}

// InstallWriterCommandFlags installs writer-specific command flags.
func InstallWriterCommandFlags(c *cobra.Command, opts *Options) {
	c.Flags().BoolVarP(&opts.KeepFiles, "keep-files", "k", false, "download file contents to output directory")

	c.Flags().StringVarP(&opts.OutputDirectory, "output-dir", "o", opts.OutputDirectory, fmt.Sprintf("Output directory, or \"%s\" for STDOUT.", StdoutOutputDirectory))
}

// NewWriter runs the writer.
func NewWriter(data interface{}, extension string, opts *Options) (io.Writer, error) {
	if opts.OutputDirectory == StdoutOutputDirectory {
		return os.Stdout, nil
	}
	return os.Create(filepath.Join(opts.OutputDirectory, fileName(data, extension)))
}

func fileName(data interface{}, extension string) string {
	switch v := data.(type) {
	case api.Thing:
		return fmt.Sprintf("%v.%s", v.Id, extension)
	default:
		return fmt.Sprintf("things.%s", extension)
	}
}

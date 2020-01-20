package cmd

import (
	"fmt"
	"os"

	api "github.com/joshjcarrier/thingiverse-cli/pkg/api/thingiverse"
	"github.com/joshjcarrier/thingiverse-cli/pkg/render"
	"github.com/joshjcarrier/thingiverse-cli/pkg/writer"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	options = &Options{
		API: &api.Options{},
		Render: &render.Options{
			Format: "txt",
		},
		Writer: &writer.Options{
			OutputDirectory: writer.StdoutOutputDirectory,
		},
	}
	rootCmd = &cobra.Command{
		Use:   "thing",
		Short: "Manage your Things from Thingiverse!",
	}
)

// Options for command execution
type Options struct {
	API    *api.Options
	Render *render.Options
	Writer *writer.Options
}

// Execute executes the root command.
func Execute(version string) {
	rootCmd.Version = version
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(searchCmd)
	rootCmd.AddCommand(userCmd)
	rootCmd.AddCommand(loginCmd)
}

func initConfig() {
	viper.SetConfigFile(".thingrc")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME")
	viper.AutomaticEnv()
	viper.ReadInConfig()

	options.API.Token = viper.GetString("token")
}

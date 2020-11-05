package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/birchwood-langham/bootstrap/pkg/config"
)

var Version = &cobra.Command{
	Use:   "version",
	Short: "Current version",
	Long:  "Get the current version number of the service",
	Run: func(*cobra.Command, []string) {
		fmt.Printf("Version: %s\n", viper.GetString(config.VersionKey))
	},
}

func init() {
	AddCommand(Version)
}

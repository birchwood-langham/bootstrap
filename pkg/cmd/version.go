package cmd

import (
	"fmt"

	"github.com/birchwood-langham/web-service-bootstrap/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

package cmd

import (
	"fmt"
	"math"
	"os"
	"os/signal"
	"syscall"

	"github.com/birchwood-langham/go-toolkit/io/strings"
	"github.com/birchwood-langham/web-service-bootstrap/service"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/bl-go/service-bootstrap/pkg/logger"
	"go.uber.org/zap"
)

var cfgFile string
var application service.Application
var log *zap.Logger

var Root = &cobra.Command{
	Run: startService,
}

func MaxPort() int {
	return math.MaxUint16
}

func startService(cmd *cobra.Command, args []string) {
	if err := application.Init(); err != nil {
		log.Fatal("could not initialize the application", zap.Error(err))
	}

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	incoming := <-signalCh

	log.Warn("Caught signal, terminating", zap.String("signal", incoming.String()))

	if err := application.Cleanup(); err != nil {
		log.Fatal("could not execute cleanup", zap.Error(err))
	}
}

func init() {
	log = logger.ConsoleLogger()
	cobra.OnInitialize(initConfig)

	Root.PersistentFlags().StringVar(&cfgFile, "config", "", "configuration file to use for the service")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			log.Fatal("could not get user home directory", zap.Error(err))
		}

		executable, err := os.Executable()

		if err != nil {
			log.Fatal("could not get the current executable name", zap.Error(err))
		}

		executablePath := strings.SplitAndTrimSpace(executable, string(os.PathSeparator))

		appName := executablePath[len(executablePath)-1]

		viper.AddConfigPath(".")
		viper.AddConfigPath("./config")
		viper.AddConfigPath(fmt.Sprintf("%s/.config/%s", home, appName))
		viper.SetConfigName("configuration")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("could not read application configuration file", zap.Error(err))
	}

	setupLogger()

	log.Info("using configuration", zap.String("file-path", viper.ConfigFileUsed()))
}

func setupLogger() {
	log = logger.New(logger.ApplicationLogLevel(), logger.ConfiguredLumberjackLogger())
	zap.ReplaceGlobals(log)
}

func AddCommand(commands ...*cobra.Command) {
	Root.AddCommand(commands...)
}

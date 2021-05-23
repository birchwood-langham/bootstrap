package cmd

import (
	"context"
	"fmt"
	"math"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/birchwood-langham/bootstrap/pkg/io/strings"
	"github.com/birchwood-langham/bootstrap/pkg/logger"
	"github.com/birchwood-langham/bootstrap/pkg/service"
)

var cfgFile string
var log *zap.Logger
var ctx context.Context
var app service.Application
var state service.StateStore

var RootCmd = &cobra.Command{
	Run: startService,
}

func MaxPort() int {
	return math.MaxUint16
}

func startService(cmd *cobra.Command, args []string) {
	if err := app.Init(ctx, state); err != nil {
		log.Fatal("could not initialize the application", zap.Error(err))
	}

	if app.RunFunction() != nil {
		if err := app.RunFunction()(ctx, state); err != nil {
			log.Error("Command failed", zap.Error(err))
		}

		if err := app.Cleanup(state); err != nil {
			log.Fatal("could not execute cleanup", zap.Error(err))
		}

		return
	}

	log.Info("Starting service. Ctrl-C to terminate")

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	incoming := <-signalCh

	log.Warn("Caught signal, terminating", zap.String("signal", incoming.String()))

	if err := app.Cleanup(state); err != nil {
		log.Fatal("could not execute cleanup", zap.Error(err))
	}
}

func init() {
	log = logger.ConsoleLogger()
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "configuration file to use for the service")
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

	log.Debug("using configuration", zap.String("file-path", viper.ConfigFileUsed()))
}

func setupLogger() {
	log = logger.Get(logger.ApplicationLogLevel(), logger.ConfiguredLumberjackLogger())
	zap.ReplaceGlobals(log)
}

func AddCommand(commands ...*cobra.Command) {
	RootCmd.AddCommand(commands...)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(c context.Context, a service.Application, s service.StateStore) {
	ctx = c
	app = a
	state = s

	RootCmd.Use = service.Usage
	RootCmd.Short = service.ShortDescription
	RootCmd.Long = service.LongDescription

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

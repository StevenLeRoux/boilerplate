package cmd

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var projectName = "boilerplate"

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().Int32P("log-level", "l", 4, "set logging level between 0 and 5")
	RootCmd.PersistentFlags().String("config", "", "set config file")

	if err := viper.BindPFlags(RootCmd.PersistentFlags()); err != nil {
		log.Fatal(err)
	}

	if err := viper.BindPFlags(RootCmd.Flags()); err != nil {
		log.Fatal(err)
	}
}

func initConfig() {
	// Set viper defaults
	viper.SetDefault("metrics.host", "127.0.0.1")
	viper.SetDefault("metrics.port", 9100)

	// Environment variables management
	viper.SetEnvPrefix(projectName)
	viper.SetEnvKeyReplacer(strings.NewReplacer("_", "."))
	viper.AutomaticEnv()

	// Set config search path
	viper.AddConfigPath("/etc/" + projectName + "/")
	viper.AddConfigPath("$HOME/." + projectName)
	viper.AddConfigPath(".")

	// Load config
	viper.SetConfigName("config")
	if err := viper.MergeInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Debug("Config error: no config file found")
		} else {
			log.Fatalf("Config error: fatal error in config file: %v \n", err)
		}
	}

	// Load user defined config
	config := viper.GetString("config")
	if len(config) > 0 {
		viper.SetConfigFile(config)
		err := viper.ReadInConfig()
		if err != nil {
			log.Panicf("Config error: fatal error in config file: %v \n", err)
		}
	}

	level := viper.GetInt("log-level")
	log.SetLevel(log.InfoLevel) // level == 4
	if level >= 0 && level < len(log.AllLevels) {
		log.SetLevel(log.AllLevels[level])
	}
}

// RootCmd is the cli root command
var RootCmd = &cobra.Command{
	Use: projectName,
	Run: func(cmd *cobra.Command, arguments []string) {

	},
}

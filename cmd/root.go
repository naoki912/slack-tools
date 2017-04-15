package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
)

var RootCmd = &cobra.Command{
	Use:   "slack-tools",
	Short: "slack cli tools",
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringP("token", "", "", "slack token")

	viper.BindPFlag("slack.token", RootCmd.PersistentFlags().Lookup("token"))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".slack-tools")
	viper.AddConfigPath("$HOME")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	viper.BindEnv("slack.token", "ST_SLACK_TOKEN")
}

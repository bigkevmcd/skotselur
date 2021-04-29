package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const debugFlag = "debug"

func init() {
	cobra.OnInitialize(initConfig)
}

func logIfError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func initConfig() {
	viper.AutomaticEnv()
}

func makeRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "kapp",
		Short:         "application manager",
		SilenceErrors: true,
	}
	cmd.AddCommand(makeInstallCmd())
	return cmd
}

// Execute is the main entry point into this component.
func Execute() {
	if err := makeRootCmd().Execute(); err != nil {
		log.Fatal(err)
	}
}

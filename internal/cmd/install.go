package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/bigkevmcd/skotuselur/pkg/profiles"
	"github.com/go-logr/zapr"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

const (
	profileURLFlag    = "profile-url"
	profileBranchFlag = "branch"
)

func makeInstallCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "install",
		Short: "install profile",
		Run: func(cmd *cobra.Command, args []string) {
			logger := zapr.NewLogger(makeLogger(viper.GetBool(debugFlag)))
			// TODO: cmd flag for path
			if err := profiles.Install(context.Background(), logger, ".",
				profiles.ProfileRef{
					RepoURL: viper.GetString(profileURLFlag),
					Branch:  viper.GetString(profileBranchFlag),
				}); err != nil {
				fmt.Println("failed to install profile: %s\n", err)
				os.Exit(1)
			}
		},
	}

	cmd.Flags().String(
		profileURLFlag,
		"",
		"profile URL to install",
	)
	cmd.Flags().String(
		profileBranchFlag,
		"",
		"branch within profile repository to install from",
	)
	cmd.Flags().Bool(
		debugFlag,
		false,
		"enable debug logging",
	)
	logIfError(cmd.MarkFlagRequired(profileURLFlag))
	logIfError(viper.BindPFlags(cmd.Flags()))
	return cmd
}

func makeLogger(debug bool) *zap.Logger {
	var zapLog *zap.Logger
	var err error
	if debug {
		zapLog, err = zap.NewDevelopment()
	} else {
		zapLog, err = zap.NewProduction()
	}
	logIfError(err)
	return zapLog
}

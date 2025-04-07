package cmd

import (
	"blog/internal/app"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	StartCmd = &cobra.Command{
		Use:   "start",
		Short: "Start the server",
		Long:  `Start the server`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := app.RunServer(); err != nil {
				log.Fatal().Err(err).Msg("Failed to start server")
			}
		},
	}
)

func init() {
	RootCmd.AddCommand(StartCmd)
}

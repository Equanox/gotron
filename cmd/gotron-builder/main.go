package main

import (
	"gotron/cmd/gotron-builder/internal/application"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.PersistentFlags().StringP("example-string", "", "", "description")
	rootCmd.PersistentFlags().IntP("example-int", "p", 1, "description")
	rootCmd.PersistentFlags().BoolP("example-bool", "", false, "description")
}

func Run(cmd *cobra.Command, args []string) {
	zerolog.TimeFieldFormat = ""
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Caller().Logger()

	// Parse flags
	// s := cmd.Flag("example-string").Value.String()
	// i, _ := strconv.ParseInt(cmd.Flag("example-int").Value.String(), 10, 0)
	// b, _ := strconv.ParseBool(cmd.Flag("example-bool").Value.String())

	app := application.New()

	if err := app.Run(); err != nil {
		log.Fatal().Msg(err.Error())
	}
}

var rootCmd = &cobra.Command{
	Use:   "gotron-builder",
	Short: "building gotron",
	Long:  `TODO`,
	Run:   Run,
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Error().Msg(err.Error())
		os.Exit(1)
	}
}

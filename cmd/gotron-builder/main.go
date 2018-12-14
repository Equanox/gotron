package main

import (
	"fmt"
	"github.com/Equanox/gotron/cmd/gotron-builder/internal/application"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.PersistentFlags().StringP("example-string", "", "", "description")
	rootCmd.PersistentFlags().StringP("go-entrypoint", "g", ".", "description")
	rootCmd.PersistentFlags().StringP("app-directory", "a", ".gotron/assets/", "description")
	rootCmd.PersistentFlags().IntP("example-int", "p", 1, "description")
	rootCmd.PersistentFlags().BoolP("example-bool", "", false, "description")
}

func Run(cmd *cobra.Command, args []string) {
	zerolog.TimeFieldFormat = ""
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Caller().Logger()

	// Parse flags
	s := cmd.Flag("example-string").Value.String()
	appDir := cmd.Flag("app-directory").Value.String()
	goDir := cmd.Flag("go-entrypoint").Value.String()
	// i, _ := strconv.ParseInt(cmd.Flag("example-int").Value.String(), 10, 0)
	// b, _ := strconv.ParseBool(cmd.Flag("example-bool").Value.String())

	fmt.Println(s)
	fmt.Println(appDir)
	fmt.Println(goDir)

	app := application.New()

	app.GoEntryPoint = goDir
	app.AppDir = appDir

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

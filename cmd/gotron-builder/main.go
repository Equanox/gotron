package main

import (
	"github.com/Benchkram/errz"
	"fmt"
	"github.com/Equanox/gotron/cmd/gotron-builder/internal/application"
	"os"
	"path/filepath"
	"runtime"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.PersistentFlags().StringP("go", "g", ".",
		"Go entrypoint, must point to a directory containing a main.go")
	rootCmd.PersistentFlags().StringP("app", "a", ".gotron/assets/",
		"Application directory, must point to a directory containing a webapp starting at index.html")

	rootCmd.PersistentFlags().StringP("target", "t", runtime.GOOS,
		"target system, defaults to your os")

	//rootCmd.PersistentFlags().StringP("example-string", "", "", "description")
	//rootCmd.PersistentFlags().IntP("example-int", "p", 1, "description")
	//rootCmd.PersistentFlags().BoolP("example-bool", "", false, "description")
}

func Run(cmd *cobra.Command, args []string) {
	zerolog.TimeFieldFormat = ""
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Caller().Logger()

	// Parse flags
	goDir := cmd.Flag("go").Value.String()
	appDir := cmd.Flag("app").Value.String()
	target := cmd.Flag("target").Value.String()
	// s := cmd.Flag("example-string").Value.String()
	// i, _ := strconv.ParseInt(cmd.Flag("example-int").Value.String(), 10, 0)
	// b, _ := strconv.ParseBool(cmd.Flag("example-bool").Value.String())

	if (goDir != ".") && ((appDir == ".gotron/assets/") || (appDir == ".gotron/assets")) {
		appDir = filepath.Join(goDir, appDir)
	}

	// make paths absolute
	appDir, err := filepath.Abs(appDir)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	goDir, err = filepath.Abs(goDir)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	fmt.Println(appDir)
	fmt.Println(goDir)
	fmt.Println(target)

	app := application.New()

	app.GoEntryPoint = goDir
	app.AppDir = appDir
	err = app.SetTarget(target)
	errz.Log(err)

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

// TODO
//
// - gotron-builder deletes asset dir from .gotron

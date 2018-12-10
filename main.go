package main

import (
	"os"

	gotron "github.com/Benchkram/gotron-browser-window"

	"github.com/fatih/color"
	"github.com/spf13/viper"

	"github.com/Benchkram/errz"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

//Globals
var mainLogger zerolog.Logger // = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

func main() {

	//Logging
	mainLogger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	subLogger := mainLogger.With().Caller().Logger()

	//Register Logger in libraries
	gotron.UseLogger(subLogger)
	errz.UseZeroLog(subLogger)

	config := loadConfig()

	window, err := gotron.New(config.name, config.pathsToJSFiles, config.pathsToCSSFiles, config.appFolder)
	errz.Log(err)

	window.WindowOptions.Width = 1200
	window.WindowOptions.Height = 600

	done, err := window.Start()
	errz.Log(err)

	window.OpenDevTools()

	<-done
}

//Backend Configuration returned by loadConfig
type configuration struct {
	name          		string //Application Name
	pathsToJSFiles 		[]string //Application Frontend
	pathsToCSSFiles     []string //Application Frontend Styling
	appFolder     		string //Application Frontend Path
}

// Loads configuration from file
// or inits values with default values
func loadConfig() configuration {
	viper.SetConfigName("config")

	// Paths to search for a config file
	viper.AddConfigPath("./")
	viper.SetConfigType("toml")

	err := viper.ReadInConfig()
	if err != nil {
		color.Set(color.FgRed)
		mainLogger.Warn().Msg("No configuration file loaded - using defaults")
		color.Unset()

		// default values
		viper.SetDefault("name", "")
		viper.SetDefault("pathsToJSFiles", []string{"ui/react/build/bundle.js"})
		viper.SetDefault("pathsToCSSFiles", []string{})
		viper.SetDefault("appFolder", "gotron/")
	}

	return configuration{
		name:          viper.GetString("name"),
		pathsToJSFiles: viper.GetStringSlice("pathsToJSFiles"),
		pathsToCSSFiles:     viper.GetStringSlice("pathsToCSSFiles"),
		appFolder:     viper.GetString("appFolder")}
}

package application

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/Benchkram/errz"

	gotron "github.com/Benchkram/gotron-browser-window"
)

// Globals constants
const (
	tmpDir = ".gotron-builder"
)

type App struct {
	GoEntryPoint string
	AppDir       string
	BuildOS      string
}

type GoBuildOptions struct {
	GoEnv        map[string]string
	buildOptions map[string]string
}

func (app *App) Run() (err error) {
	fmt.Print("All your bases are belong to us!\n")

	gbw, err := gotron.New(app.AppDir)
	gbw.AppDirectory = filepath.Join(app.GoEntryPoint, ".gotron/")

	err = gbw.CreateAppStructure(false)

	err = app.makeTempDir()
	err = app.installDependencies()

	err = app.buildElectron()

	err = app.buildGoCode()

	return
}

func New() *App {
	return &App{}
}

func (app *App) makeTempDir() (err error) {
	err = os.Mkdir(tmpDir, 0777)
	// err = os.Mkdir(filepath.Join(tmpDir, ".gotron/"), 0777)

	return
}

func runCmd(runDir, command string, args ...string) (err error) {
	defer errz.Recover(&err)

	fmt.Println(runDir)
	fmt.Println(command)
	fmt.Println(args)

	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = runDir
	err = cmd.Start()

	errz.Fatal(err)

	err = cmd.Wait()
	errz.Fatal(err)

	return
}

func (app *App) installDependencies() (err error) {

	args := []string{"install", "electron-builder", "--save-dev"}

	err = runCmd(tmpDir, "npm", args...)

	return
}

func (app *App) buildElectron() (err error) {

	projDir, err := filepath.Abs(filepath.Join(app.GoEntryPoint, ".gotron/"))

	buildOS := "-l"

	switch app.BuildOS {
	case "windows":
		buildOS = "-w"
	case "linux":
		buildOS = "-l"
	case "darwin":
		buildOS = "-m"
	}

	args := []string{buildOS, "--x64", "--dir", "--projectDir=" + projDir}

	runDir := tmpDir
	command := filepath.Join("node_modules/.bin/", "electron-builder")

	err = runCmd(runDir, command, args...)

	return
}

func (app *App) buildGoCode() (err error) {

	args := []string{"build", "-tags", "gotronbrowserwindowprod"}
	runDir := app.GoEntryPoint
	command := "go"

	fName := filepath.Base(runDir)

	if app.BuildOS == "windows" {
		fName = fName + ".exe"
	}

	err = runCmd(runDir, command, args...)

	buildOS := "linux"

	switch app.BuildOS {
	case "windows":
		buildOS = "win"
	case "linux":
		buildOS = "linux"
	case "darwin":
		buildOS = "mac"
	}

	err = os.Rename(filepath.Join(runDir, fName), filepath.Join(app.GoEntryPoint, ".gotron/dist/"+buildOS+"-unpacked/", fName))

	return
}

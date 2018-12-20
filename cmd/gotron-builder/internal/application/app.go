package application

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/Equanox/gotron/cmd/gotron-builder/internal/file"

	"github.com/Benchkram/errz"
	gotron "github.com/Equanox/gotron/pkg/browser-window"
)

// Globals constants
const (
	gotronBuilderDirectory = ".gotron-builder"
)

type App struct {
	GoEntryPoint string
	AppDir       string
	Target       string
}

type GoBuildOptions struct {
	GoEnv        map[string]string
	buildOptions map[string]string
}

func (app *App) Run() (err error) {
	defer errz.Recover(&err)

	// Use gotron-browser-window to copy webapp
	// to .gotron dir. Let it handle the necessary logic
	// to validate webapp.
	gbw, err := gotron.New(app.AppDir)
	err = gbw.CreateAppStructure()
	errz.Fatal(err)

	err = app.makeTempDir()
	errz.Fatal(err)

	err = app.installDependencies()
	errz.Fatal(err)

	err = app.buildElectron()
	errz.Fatal(err)

	err = app.buildGoCode()
	errz.Fatal(err)

	return err
}

func New() *App {
	var target string
	switch runtime.GOOS {
	case "windows":
		target = "win"
	case "linux":
		target = "linux"
	case "darwin":
		target = "mac"
	default:
		target = runtime.GOOS
	}

	return &App{
		Target: target,
	}
}

func (app *App) makeTempDir() (err error) {
	os.RemoveAll(gotronBuilderDirectory)
	return os.Mkdir(gotronBuilderDirectory, os.ModePerm)
}

func runCmd(runDir, command string, args ...string) (err error) {
	defer errz.Recover(&err)

	// fmt.Println(runDir)
	// fmt.Println(command)
	// fmt.Println(args)

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

	return runCmd(gotronBuilderDirectory, "npm", args...)
}

// buildElectron
//
// (1) Uses AppDir
// (2)
func (app *App) buildElectron() (err error) {
	if !file.Exists(app.AppDir) {
		return errors.New(
			fmt.Sprintf("Given application directory [%s] does not exist",
				app.AppDir,
			))
	}
	// contains

	projDir, err := filepath.Abs(filepath.Join(app.GoEntryPoint, ".gotron/"))

	args := []string{app.Target, "--x64", "--dir", "--projectDir=" + projDir}

	runDir := gotronBuilderDirectory
	command := filepath.Join("node_modules/.bin/", "electron-builder")

	return runCmd(runDir, command, args...)
}

func (app *App) buildGoCode() (err error) {
	defer errz.Recover(&err)
	args := []string{"build", "-tags", "gotronbrowserwindowprod"}
	runDir := app.GoEntryPoint
	command := "go"

	fName := filepath.Base(runDir)

	if app.Target == "win" {
		fName = fName + ".exe"
	}

	err = runCmd(runDir, command, args...)
	errz.Fatal(err)

	from := filepath.Join(runDir, fName)
	to := filepath.Join(app.GoEntryPoint, ".gotron/dist", app.Target+"-unpacked", fName)
	return os.Rename(from, to)
}

// Will copy everythin from .gotron/dist to .dist
func (app *App) syncDistDirs() (err error) {
	defer errz.Recover(&err)

	err = os.MkdirAll(".dist", os.ModePerm)
	errz.Fatal(err)

	wkdirFiles, err := ioutil.ReadDir(".dist")
	errz.Fatal(err)

	files, err := ioutil.ReadDir(".gotron/dist")
	errz.Fatal(err)

	for _, f := range files {
		for _, wf := range wkdirFiles {
			if f.Name() == wf.Name() {
				p := filepath.Join(".dist", wf.Name())
				err = os.RemoveAll(p)
				errz.Fatal(err)
			}
		}
	}

	// TODO copy tree ....

	return nil
}

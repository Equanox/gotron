// +build !gotronpack

package gotron

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/Equanox/gotron/internal/file"
	"github.com/pkg/errors"

	"github.com/puengel/copy"

	"github.com/Benchkram/errz"
)

const (
	templateApplicationDir = "templates/app"
)

// Start starts an Instance of gotronbrowserwindow
func (gbw *BrowserWindow) Start(forceInstall ...bool) (isdone chan bool, err error) {
	defer errz.Recover(&err)

	var _forceInstall bool
	for _, v := range forceInstall {
		_forceInstall = v
		break
	}

	isdone = done

	// build up structure
	err = gbw.CreateAppStructure(_forceInstall)
	errz.Fatal(err)

	// run sockets and electron
	err = gbw.runApplication()
	errz.Fatal(err)

	return
}

// CreatAppStructure -
// Get electron and web files. Put them into gbw.AppFolder (default ".gotron")
func (gbw *BrowserWindow) CreateAppStructure(forceInstall ...bool) (err error) {
	var _forceInstall bool
	for _, v := range forceInstall {
		_forceInstall = v
	}
	defer errz.Recover(&err)

	err = os.MkdirAll(gbw.AppDirectory, 0777)
	errz.Fatal(err)

	// Copy Electron Files
	err = gbw.copyElectronApplication(_forceInstall)
	errz.Fatal(err)

	// Run npm install
	err = gbw.runNPM(_forceInstall)
	errz.Fatal(err)

	return nil
}

// runApplication starts websockets and runs the electron application
func (gbw *BrowserWindow) runApplication() (err error) {
	//run websocket
	gbw.runWebsocket()

	//get electron start parameters
	electronPath, args, err := gbw.createStartParameters()
	errz.Fatal(err)

	//run electron
	electron := exec.Command(electronPath, args...)

	electron.Stdout = os.Stdout
	electron.Stderr = os.Stderr

	err = electron.Start()
	errz.Fatal(err)

	gbw.Running = true

	return
}

// copyElectronApplication from library package to defined app directory.
// copy app files (.js .css) to app directory
//
// forceInstall forces a reinstallation of electron
// and resets AppDirectory/assets if no UIFolder was set.
//
// On the first run we copy a default application
// into AppDirectory and install electronjs locally.
// When a ui directory was set we use the contents of those
// and copy it into AppDirectory/assets
func (gbw *BrowserWindow) copyElectronApplication(forceInstall bool) (err error) {
	defer errz.Recover(&err)

	// Copy app Directory
	mainJS := filepath.Join(gbw.AppDirectory, "main.js")
	firstRun := !file.Exists(mainJS)

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return errors.New("No caller information")
	}
	gbwDirectory := filepath.Dir(filename)

	if firstRun || forceInstall {
		templateDir := filepath.Join(gbwDirectory, templateApplicationDir)
		err = copy.Perm(templateDir, gbw.AppDirectory, 0777, 0644)
		errz.Fatal(err)
	}

	// If no UI folder is set use default ui files
	if gbw.UIFolder == "" {
		return
	}

	// UIFolder must contain a index.htm(l)
	html := filepath.Join(gbw.UIFolder, "index.html")
	htm := filepath.Join(gbw.UIFolder, "index.htm")
	if !(file.Exists(html) || file.Exists(htm)) {
		return fmt.Errorf("index.htm(l) missing in %s", gbw.UIFolder)
	}

	// No need to copy web application files
	// when no ui folder is set.
	// Also check for ".gotron/assets". This is the
	// default directory when called from gotron-builder,
	// avoids deleting asset dir by accident.
	src, err := filepath.Abs(gbw.UIFolder)
	errz.Fatal(err)
	dst, err := filepath.Abs(filepath.Join(gbw.AppDirectory, "assets"))
	errz.Fatal(err)

	if src != dst {
		err = os.RemoveAll(filepath.Join(gbw.AppDirectory, "assets"))
		errz.Fatal(err)

		err = copy.Copy(gbw.UIFolder, filepath.Join(gbw.AppDirectory, "assets"))
		errz.Fatal(err)
	}

	return nil
}

//runNPM - run npm install if not done already or foced.
func (gbw *BrowserWindow) runNPM(forceinstall bool) (err error) {
	defer errz.Recover(&err)

	nodeModules := filepath.Join(gbw.AppDirectory, "node_modules/")
	forceinstall = !file.Exists(nodeModules)

	if forceinstall {
		logger.Debug().Msgf("Installing npm packages...")

		cmd := exec.Command("npm", "install")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Dir = gbw.AppDirectory
		err = cmd.Start()

		errz.Fatal(err)

		logger.Debug().Msgf("Waiting for batch")

		err = cmd.Wait()
		errz.Fatal(err)
		logger.Debug().Msgf("Batch done")
	}
	return err
}

//runWebsocket with defined port or look for free port if taken
func (gbw *BrowserWindow) runWebsocket() {
	var err error
	errz.Recover(&err)
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(gbw.Port))
	errz.Fatal(err)

	logger.Debug().Msgf("Using port: %d", listener.Addr().(*net.TCPAddr).Port)
	gbw.Port = listener.Addr().(*net.TCPAddr).Port

	// Endpoint for Electron startup/teardown
	// + browser window events to nodejs
	http.HandleFunc("/browser/window/events", gbw.mainEventSocket)
	// Endpoint for ipc like messages
	// send from user web application
	http.HandleFunc("/web/app/events", gbw.onSocket)
	go http.Serve(listener, nil) // Start websockets in goroutine

}

//createStartParameters returns absolute electron path and list of arguments to be passed on electron run call.
func (gbw *BrowserWindow) createStartParameters() (electronPath string, arguments []string, err error) {
	defer errz.Recover(&err)

	electronPath, err = filepath.Abs(filepath.Join(gbw.AppDirectory + "/node_modules/.bin/electron"))
	errz.Fatal(err)
	appPath, err := filepath.Abs(gbw.AppDirectory + "main.js")
	errz.Fatal(err)
	logger.Debug().Msgf(appPath)

	configString, err := json.Marshal(gbw.WindowOptions)
	errz.Fatal(err)

	arguments = []string{appPath, strconv.Itoa(gbw.Port), string(configString)}

	return
}

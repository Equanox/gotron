// +build gotronpack

package gotron

import (
	"encoding/json"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/Benchkram/errz"
)

// Start starts an Instance of gotronbrowserwindow
func (gbw *BrowserWindow) Start(forceInstall ...bool) (isdone chan bool, err error) {
	defer errz.Recover(&err)

	isdone = done

	// run sockets and electron
	err = gbw.runApplication()
	errz.Fatal(err)

	return
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

	ex, err := os.Executable()
	errz.Fatal(err)

	exPath := filepath.Dir(ex)

	if runtime.GOOS == "darwin" {
		electronPath, err = filepath.Abs(filepath.Join(exPath, "electronjs", "gotron-browser-window.app/Contents/MacOS/gotron-browser-window"))
	} else {
		electronPath, err = filepath.Abs(filepath.Join(exPath, "electronjs", "gotron-browser-window"))
	}
	errz.Fatal(err)
	appPath := ""
	logger.Debug().Msgf(appPath)

	configString, err := json.Marshal(gbw.WindowOptions)
	errz.Fatal(err)

	//TODO: test if this call works properly
	arguments = []string{appPath, strconv.Itoa(gbw.Port), string(configString)}

	return
}

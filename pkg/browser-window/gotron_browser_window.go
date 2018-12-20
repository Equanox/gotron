// Package gotron :
//
// Rampup electron from golang using only a go api.
package gotron

import (
	"encoding/json"
	"github.com/Equanox/gotron/pkg/browser-window/internal/runner"
	"net/http"
	"sync"

	"github.com/Benchkram/errz"
	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
)

// SocketEvent event
type SocketEvent struct {
	Event string
	Data  interface{}
	ID    uuid.UUID
}

// Configuration Backend Configuration returned by loadConfig
type Configuration struct {
	UIFolder     string
	AppDirectory string // Directory to search for a electron application
	Port         int
}

// optionsQueueElement
// Data: event to be sent
// Waiter: received event
type optionsQueueElement struct {
	Waiter chan SocketEvent
	Data   SocketEvent
}

// BrowserWindow Instance for a gotronbrowserwindow
type BrowserWindow struct {
	Configuration
	UseZerolog       bool
	Running          bool
	handledMessages  map[string]func(interface{})
	WindowOptions    WindowOptions
	optionsQueue     chan optionsQueueElement
	optionsReturnMap sync.Map
}

// New creates a new gotronbrowserwindow,
// parameter uiFolder must point to a folder containing either an index.htm or an index.html file
// if empty a default aplication is used
func New(uiFolders ...string) (gbw *BrowserWindow, err error) {
	err = nil

	uiFolder := ""
	for _, v := range uiFolders {
		uiFolder = v
		break
	}

	gbw = &BrowserWindow{
		Configuration: Configuration{
			AppDirectory: ".gotron/",
			UIFolder:     uiFolder,
		},
		UseZerolog:      false,
		Running:         false,
		handledMessages: make(map[string]func(interface{})),
		optionsQueue:    make(chan optionsQueueElement, 100),
		//Set default WindowOption bools
		WindowOptions: WindowOptions{
			Width:          800,
			Height:         800,
			Resizable:      true,
			Movable:        true,
			Minimizable:    true,
			Maximizable:    true,
			Closable:       true,
			Focusable:      true,
			Fullscreenable: true,
			Show:           true,
			Frame:          true,
			HasShadow:      true,
			ThickFrame:     true,
			WebPreferences: WebPreferences{
				DevTools:              true,
				NodeIntegration:       true,
				EnableRemoteModule:    true,
				Javascript:            true,
				WebSecurity:           true,
				Images:                true,
				TextAreasAreResizable: true,
				Webgl:                 true,
				Webaudio:              true,
				BackgroundThrottling:  true,
				WebviewTag:            true,
			},
		},
	}

	return gbw, nil
}

//Handles msgs to communicate with nodejs electron for rampup & shutdown
func (gbw *BrowserWindow) mainEventSocket(w http.ResponseWriter, r *http.Request) {
	var err error
	defer errz.Recover(&err)

	c, err := upgrader.Upgrade(w, r, nil)
	errz.Fatal(err)
	defer c.Close()

	//Writer
	writerTask := runner.Go(func(stop runner.StopChan, finish runner.Finish) {
		for {
			var request optionsQueueElement
			select {
			case request = <-gbw.optionsQueue:
				logger.Debug().Msgf("Has Option Request")
				u, err := uuid.NewV4()
				errz.Fatal(err)

				logger.Debug().Msgf("Sending Request")
				logger.Debug().Msgf("%+v\n", u)
				request.Data.ID = u

				gbw.optionsReturnMap.Store(u, request.Waiter)

				logger.Debug().Msgf("%+v\n", request.Data)

				err = c.WriteJSON(request.Data)
				errz.Log(err)
			case _, ok := <-stop:
				if !ok {
					finish()
					return
				}

			}
		}
	})

	//Reader
	for {

		// ReadMessages
		var event SocketEvent
		_, message, err := c.ReadMessage()
		errz.Log(err, "ElectronSocket: [err]")

		//Handle Message
		err = json.Unmarshal(message, &event)
		errz.Fatal(err, "Unmashal: ")
		logger.Debug().Msgf("ElectronSocket: [received] %+v", event)

		//Shutdown Event
		if event.Event == "shutdown" {
			switch t := event.Data.(type) {
			case bool:
				if t {
					writerTask.Stop()
					writerTask.Wait()
					gbw.Running = false
					done <- true
				}
			}
		} else {
			ch, ok := gbw.optionsReturnMap.Load(event.ID)
			if !ok {
				logger.Debug().Msgf("Event not in return map")
				break
			}
			ch.(chan SocketEvent) <- event
			gbw.optionsReturnMap.Delete(event.ID)
		}
	}
}

//Handles msgs to communicate with nodejs electron for rampup & shutdown
func (gbw *BrowserWindow) onSocket(w http.ResponseWriter, r *http.Request) {
	var err error
	defer errz.Recover(&err)

	c, err := upgrader.Upgrade(w, r, nil)
	errz.Fatal(err)
	defer c.Close()

	for {
		var event SocketEvent
		_, message, err := c.ReadMessage()
		errz.Fatal(err, "ElectronSocket: [err]")

		//Handle Message
		err = json.Unmarshal(message, &event)
		errz.Fatal(err, "Unmashal: ")
		logger.Debug().Msgf("ElectronSocket: [received] %+v", event)

		//Execute event function if exists
		if f, ok := gbw.handledMessages[event.Event]; ok {
			f(event.Data)
		} else {
			logger.Debug().Msgf("Event not registered: %s", event.Event)
		}
	}
}

//Globals
var done = make(chan bool, 1)      //Wait for Shutdown signal over websocket
var upgrader = websocket.Upgrader{ //Upgrader for websockets
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	Subprotocols:    []string{"p0", "p1"},
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//On register handler for messages incoming from js frontend
func (gbw *BrowserWindow) On(message string, handler func(interface{})) {
	logger.Debug().Msgf("Adding handler for message: " + message)
	gbw.handledMessages[message] = handler
}

//Send send message (with data) to js frontend
func (gbw *BrowserWindow) Send(message string, data interface{}) {
	logger.Debug().Msgf("Sending message: " + message)
	logger.Debug().Msgf("Content: ")
	logger.Debug().Msgf("%+v\n", data)
}

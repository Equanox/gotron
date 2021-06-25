[![Build Status](https://travis-ci.org/Equanox/gotron.svg?branch=master)](https://travis-ci.org/Equanox/gotron)

# Gotron
Go Api for Electron.

### ⚠️  **[This project is no longer maintained.](https://github.com/pxgrid/aigis/issues/121)** ⚠️
Feel free to fork and make your own changes if needed.

## Example Projects
A list of boilerplate projects using gotron.

* https://github.com/Marlon-Monroy/gotron-react-boilerplate

## Prerequisites
**go1.11** with modules enabled, **nodejs** and **npm** must be available on your system.

## Quick Start
On the first run it will download Electron and stores it in *.gotron* in your working directory.

```go
package main

import (
    "github.com/Equanox/gotron"
)

func main() {
    // Create a new browser window instance
    window, err := gotron.New()
    if err != nil {
        panic(err)
    }

    // Alter default window size and window title.
    window.WindowOptions.Width = 1200
    window.WindowOptions.Height = 980
    window.WindowOptions.Title = "Gotron"

    // Start the browser window.
    // This will establish a golang <=> nodejs bridge using websockets,
    // to control ElectronBrowserWindow with our window object.
    done, err := window.Start()
    if err != nil {
        panic(err)
    }
    
    // Open dev tools must be used after window.Start 
    // window.OpenDevTools()
    
    // Wait for the application to close
    <-done
}
```     

When everything worked you should see this

<img src="doc/hello_gotron.png" alt="Hello Gotron" width="400px"/>


### Use Your Own WebUI
gotron expects a folder containing your HTML/JS/CSS code and passes it to electronJS. Make sure it contains at least an index.html as an entrypoint.

Pass a path to your webUI on gotron's `New(uiFolder ...string)` function.
```go
window, err := gotron.New("path/to/your/webui")
if err != nil {
    panic(err)
}
```

### Communicate between backend and frontend

Frontend to backend communication is realized through javascript like event driven approach.

#### Backend

Handle incoming events
```go
window.On(&gotron.Event{Event: "event-name"}, func(bin []byte) {
	//Handle event here
}
```

Send event to frontend
```go
// Create a custom event struct that has a pointer to gotron.Event
type CustomEvent struct {
    *gotron.Event
    CustomAttribute string 'json:"AtrNameInFrontend"'
}

window.Send(&CustomEvent{
    Event: &gotron.Event{Event: "event-name"},
    CustomAttribute: "Hello World!",
    })
```

#### Frontend

In frontend a websocket needs to be created. Address is always localhost and port can be taken from global variable `global.backendPort`
```javascript
let ws = new WebSocket("ws://localhost:" + global.backendPort + "/web/app/events");
```

Handle incoming events
```javascript
// This is being called for all incoming messages
ws.onmessage = (message) => {
    let obj = JSON.parse(message.data);
    
    // event name
    console.log(obj.event);

    // event data
    console.log(obj.AtrNameInFrontend);
}
```

Send event to backend

```javascript
ws.send(JSON.stringify({
    "event": "event-name",
    "AtrNameInFrontend": "Hello World!",
}))
```

## Distribution/Packaging
To package a go application together with electornjs use `gotron-builder`.    

#### Install gotron-builder
We provide executables for Linux, MacOS, and Windows.    
Download the newest release from [https://github.com/Equanox/gotron/releases](https://github.com/Equanox/gotron/releases) and add it to your $PATH.

#### Using  gotron-builder
It expects...
* a directory containing a golang main package 
* and a directory with a webUI containing at least an index.html

By default it will implicitly use...
* golang main package from the current directory
* webUI from *.gotron/assets*

To pack the code from **Quick Start** use
```sh
gotron-builder
```
in the root of your repo.

Pass your go code and webUI explicitly.
```sh
gotron-builder --go=your/go/dir --app=your/webapp/dir
```

For cross compilation you can use the same flags as electron-builder would expect them
```
gotron-builder --win 
```
Read about the requirements for cross-compilation in [electron-builders](https://github.com/electron-userland/electron-builder) documentation.

## Tasks
- [x] Basic js + webpack example
- [x] React example
- [x] Typescript-React example
- [x] Vue.js example
- [ ] Elm example
- [ ] Flutter Hummingbird example
- [X] Hide nodejs/Electron behind go api
- [X] Create executables for Win, Linux, MacOS
- [X] Hide nodejs/Electron behind go api
- [ ] Msgs between golang and Electron renderer process,
      abstracted in a javascript/typescript package
- [ ] Implement complete BrowserWindow api see => [BrowserWindow.md](BrowserWindow.md)
- [ ] Implement complete electron-builder api in gotron-builder

## Sponsors
<a href="https://benchkram.de"><img src="https://raw.githubusercontent.com/equanox/gotron/master/doc/benchkram.png" alt="benchkram-logo" width="50"/></a>

# License
MIT  

Except Roboto (ui/js/src/Roboto-Light.ttf , ui/react/src/Roboto-Light.ttf) which is licensed under Apache 2.0   
https://github.com/google/roboto

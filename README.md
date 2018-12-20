[![Build Status](https://travis-ci.org/Equanox/gotron.svg?branch=master)](https://travis-ci.org/Equanox/gotron)

# Gotron
A go api for electronjs.

**IMPORTANT NOTICE:**     
This repository has undergone a complete rewrite. It is no longer a boilerplate application, it rather is a full electronjs api in go containing a golang <=> nodejs bridge and a separate executable to help distribute your application, it's named `gotron-builder`. You can `go get` this package and import it by your go application. 
You can still acces the old repo using the gotron-boilerplate branch. Be aware that it wont't be maintained.
**IMPORTANT NOTICE:**

## Prerequisites
**go**, **nodejs** and **npm** should be available on your system.  

## Quick start
```
package main

import (
	gotron "github.com/Equanox/gotron"
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

Run 


![Hello Gotron](https://raw.githubusercontent.com/equanox/gotron/master/doc/hello_gotron.png)


## Tasks
- [x] Basic js + webpack example
- [x] React example
- [x] Typescript-React example
- [ ] Elm example
- [x] Vue.js example
- [ ] Communication between go and electron renderer process
- <del>[ ] Electron appearance on OS</del>
- [X] Create executables for Win, Linux
- [ ] Create executables for MacOS
- [ ] Config for go-nodejs socket
- [X] Hide all nodejs/electron behind go api

## Frontend Development Workflow
Take a look into [ui/js](https://github.com/Equanox/gotron/tree/master/ui/js), [ui/react](https://github.com/Equanox/gotron/tree/master/ui/react),
[ui/typescript-react](https://github.com/Equanox/gotron/tree/master/ui/typescript) or [ui/vue](https://github.com/Equanox/gotron/tree/master/ui/vue) for details.

For plain javascript (default) use

    cd ui && npm run build  

For other frontend use

    cd ui && npm run build:${frontend}

where ${frontend} is one out of (js|react|typescript|vue).

Then type

    go run .

to bring up go backend and electron frontend.

Reload updated index.js using 'r' key.

## Distribution/Packaging

Since all of electron and nodejs is now behind gotron-browser-window api Packaging is disabled for the time being.

# License
MIT  

Except Roboto (ui/js/src/Roboto-Light.ttf , ui/react/src/Roboto-Light.ttf) which is licensed under Apache 2.0   
https://github.com/google/roboto

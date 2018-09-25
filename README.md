[![Build Status](https://travis-ci.com/Equanox/gotron.svg?branch=master)](https://travis-ci.com/Equanox/gotron)

# Gotron
A boilerplate for cross-platform desktop applications using Golang and Electron.

## Run
**go**, **nodejs** and **npm (v2.0 or later)** should be available on your system.  

Clone to your go workspace (e.g. go/src)

    git clone https://github.com/equanox/gotron

Use npm install script and start the application
```
cd gotron
npm run install
npm run build
go build
./gotron
```
Now you should see this

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

## Frontend Development Workflow
Take a look into [ui/js](https://github.com/Equanox/gotron/tree/master/ui/js), [ui/react](https://github.com/Equanox/gotron/tree/master/ui/react),
[ui/typescript-react](https://github.com/Equanox/gotron/tree/master/ui/typescript) or [ui/vue](https://github.com/Equanox/gotron/tree/master/ui/vue) for details.

For plain javascript (default) use

    npm run build  

For other frontend use

    npm run build:${frontend}

where ${frontend} is one out of (js|react|typescript|vue).

Then type

    go build

to create an executable gotron or gotron.exe (windows).

Type

    ./gotron
    
or

    gotron.exe

to bring up go backend and electron frontend.

Reload updated index.js using 'r' key.

## Distribution/Packaging

Build the required frontend first.

The electron application will be built with **electron-builder**.
For more information about build configuration visit [electron.build](https://www.electron.build/)

For required distribution type

    npm run pack:${os}

where ${os} is one out of (linux|mac|win).

Apllication will be created in ./dist/${os}/ with an executable named **gotron** inside this directory.

Run this executable to start the application.

### Cross Platfrom Compilation

Cross Platform Compilation is supported for following cases.

- Linux to:
    - Linux
    - Windows (Wine version 1.8 or later is required)
    - Mac (Compiles but not tested)
- Windows to:
    - Windows
    - Linux
- Mac to: (No tests for compilation on mac have been performed)

# License
MIT  

Except Roboto (ui/js/src/Roboto-Light.ttf , ui/react/src/Roboto-Light.ttf) which is licensed under Apache 2.0   
https://github.com/google/roboto

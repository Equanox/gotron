# Gotron
A boilerplate for cross-platform desktop applications using Golang and Electron.

## Run
**go**, **nodejs** and **npm (v2.0 or later)** should be available on your system.  

Install Electron globally

    npm install -g electron

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
- [ ] Create executables for Win, MacOS, Linux
- [ ] Config for go-nodejs socket

## Frontend Development Workflow
Take a look into [ui/js](https://github.com/Equanox/gotron/tree/master/ui/js), [ui/react](https://github.com/Equanox/gotron/tree/master/ui/react),
[ui/typescript-react](https://github.com/Equanox/gotron/tree/master/ui/typescript) or [ui/vue](https://github.com/Equanox/gotron/tree/master/ui/vue) for details.

For plain Javascript use

    npm run build  

For the react frontend use

    npm run build:react

For the typescript-react frontend use

    npm run build:typescript

For the vueJS frontend use

    npm run build:vue    

then type

    go build

to create an executable gotron or gorton.exe (windows).

Type

    ./gotron

to bring up go backend and electron frontend.

Reload updated index.js using 'r' key.

## Distribution/Packaging

Build the required frontend first.

For windows distribution type

    npm run dist:win

For linux distribution type

    npm run dist:linux

For mac distribution type

    npm run dist:mac

Distributables will be created in ./dist/\<OS\>-unpacked/

Execute 

    ./gotron

or (windows)

    gotron.exe

to run the application.

### Cross Platfrom Compilation

Cross Platform Compilation is supported for following cases.

- Linux:
    - Linux
    - Windows (Wine version 1.8 or later is required)
    - Mac (Compiles but not tested)
- Windows:
    - Windows
    - Linux
- Mac: (No tests for compilation on mac have been performed)

# License
MIT  

Except Roboto (ui/js/src/Roboto-Light.ttf , ui/react/src/Roboto-Light.ttf) which is licensed under Apache 2.0   
https://github.com/google/roboto

# Gotron
A boilerplate for cross-platform desktop applications using Golang and Electron.

## Run
**go**, **nodejs** and **npm** should be available on your system.  

Install Electron globally

    npm install -g electron

Clone to your go workspace (e.g. go/src)

    git clone https://github.com/equanox/gotron

Use npm install script and start the application
```
cd gotron
npm run install
go run main.go
```
Now you should see this

![Hello Gotron](https://raw.githubusercontent.com/equanox/gotron/master/doc/hello_gotron.png)


## Tasks
- [x] Basic js + webpack example
- [ ] Typscript example
- [ ] Elm example
- [x] React example
- [ ] Vue.js example
- [ ] Electron appearance on OS
- [ ] Create executables for Win, MacOS, Linux
- [ ] Config for go-nodejs socket

## Frontend Development Workflow
Take a look into ui/js or ui/react for details.

For plain Javascript use

    npm run build  

For the react frontend use

    npm run build-react

then type

    go run main.go

to bring up go backend and electron frontend.

Reload updated index.js using 'r' key.
# License
MIT  

Except Roboto (ui/js/src/Roboto-Light.ttf , ui/react/src/Roboto-Light.ttf) which is licensed under Apache 2.0   
https://github.com/google/roboto

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
npm run build
go run main.go
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

    npm run build-react

For the typescript-react frontend use

    npm run build-typescript

For the vueJS frontend use

    npm run build-vue    

then type

    go run main.go

to bring up go backend and electron frontend.

Reload updated index.js using 'r' key.
# License
MIT  

Except Roboto (ui/js/src/Roboto-Light.ttf , ui/react/src/Roboto-Light.ttf) which is licensed under Apache 2.0   
https://github.com/google/roboto

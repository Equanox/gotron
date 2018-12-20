# Gotron Example
An example application showing how to use 

## Run
**go**, **nodejs** and **npm** should be available on your system.  

Clone to your go workspace (e.g. go/src)

    git clone https://github.com/equanox/gotron

Use npm install script and start the application
```
cd gotron
cd ui && npm install && npm run build
go run .
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

// preload.js

const { ipcRenderer } = require('electron')
global.backendPort = (ipcRenderer.sendSync('backend-port-request', ''));
// ipcRenderer.on('backend-port-reply', (event, arg) => global.backendPort =arg);

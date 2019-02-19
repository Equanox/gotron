const { app, BrowserWindow } = require('electron')
var WebSocketClient = require('websocket').client;
const path = require('path');
var client = new WebSocketClient();
var socket = null;

const browserWindowEvents = "/browser/window/events"


// Keep a global reference of the window object, if you don't, the window will
// be closed automatically when the JavaScript object is garbage collected.
let win

function close() {
  if (socket != null) {
    socket.sendUTF(JSON.stringify({ event: "shutdown", data: true }));
    socket.close()
  }
}

function createWindow() {

  // Get BrowserWindow options from process arguments
  let opts = process.argv[3]
  opts = JSON.parse(opts)

  // Replace Preload Script
  opts.webPreferences.preload = path.resolve(`${__dirname}/preload.js`);

  // Create the browser window.
  win = new BrowserWindow(opts)

  // and load the index.html of the app.
  win.loadURL(`file://${__dirname}/assets/index.html`)

  // Open the DevTools.
  // win.webContents.openDevTools()

  // Emitted when the window is closed.
  win.on('closed', () => {
    // Dereference the window object, usually you would store windows
    // in an array if your app supports multi windows, this is the time
    // when you should delete the corresponding element.
    win = null
    close()
  })
}

// This method will be called when Electron has finished
// initialization and is ready to create browser windows.
// Some APIs can only be used after this event occurs.
app.on('ready', createWindow)

// Quit when all windows are closed.
app.on('window-all-closed', () => {
  // On macOS it is common for applications and their menu bar
  // to stay active until the user quits explicitly with Cmd + Q
  close()
  if (process.platform !== 'darwin') {
    app.quit()
  }
})

app.on('activate', () => {
  // On macOS it's common to re-create a window in the app when the
  // dock icon is clicked and there are no other windows open.
  if (win === null) {
    createWindow()
  }
})

// In this file you can include the rest of your app's specific main process
// code. You can also put them in separate files and require them here.
const ipc = require('electron').ipcMain
ipc.on('asynchronous-message', function (event, arg) {
  // event.sender.send('asynchronous-reply', 'pong')
})

//Websocket
client.on('connectFailed', function (error) {
  console.log('Connect Error: ' + error.toString());
});

client.on('connect', function (connection) {
  console.log('WebSocket Client Connected');

  socket = connection; //copy to global scope

  connection.on('error', function (error) {
    console.log("Connection Error: " + error.toString());
  });

  connection.on('close', function () {
    console.log('Websockt Connection Closed');
    app.quit()
  });

  connection.on('message', function (message) {
    // if (message.type === 'utf8') {
    //   console.log("Received: '" + message.utf8Data + "'");
    // }
    console.log("Received message:");
    //console.log(message);
    event = JSON.parse(message.utf8Data);
    console.log(event);
    // if (event.Event === "devTools"){
    //   win.webContents.openDevTools();
    // }
    let result = handleEvent(event);
    if (result !== undefined) {
      socket.sendUTF(JSON.stringify(result));
    } else {
      event.Data = "ok"
      socket.sendUTF(JSON.stringify(event));
    }
  });
});

// Port from process arguments
let port = process.argv[2]

console.log("Main.js")
console.log(port)

client.connect('ws://127.0.0.1:' + port + browserWindowEvents, []);

ipc.on('backend-port-request', (event, arg) => {
  event.returnValue = port;
});


// This array represents the map of functions which can be called by go Event messages.
// To add new functions simply extend this array.
let eventMapArray = [
  ['devTools', () => win.webContents.openDevTools()],
  ['destroy', () => win.destroy()],
  ['close', () => win.close()],
  ['focus', () => win.focus()],
  ['blur', () => win.blur()],
  ['isFocused', () => console.log(win.isFocused())],
  ['isDestroyed', () => console.log(win.isDestroyed())],
  ['show', () => win.show()],
  ['showInactive', () => win.showInactive()],
  ['hide', () => win.hide()],
  //
  ['isVisible', () => win.isVisible()],
  //
  ['isModal', () => win.isModal()],
  ['maximize', () => win.maximize()],
  ['unmaximize', () => win.unmaximize()],
  //
  ['isMaximized', () => win.isMaximized()],
  ['minimize', () => win.minimize()],
  ['restore', () => win.restore()],
  //
  ['isMinimized', () => win.isMinimized()],
  ['setFullScreen', (event) => win.setFullScreen(event.Data.Flag)],
  //
  ['isFullScreen', () => win.isFullScreen()],
  ['setSimpleFullScreen', (event) => win.setSimpleFullScreen(event.Data.Flag)],
  //
  ['isSimpleFullScreen', () => win.isSimpleFullScreen()],
  //
  ['isNormal', () => win.isNormal()],
  ['setAspectRatio', (event) => win.setAspectRatio(event.Data.AspectRatio)],
  ['setBackgroundColor', (event) => win.setBackgroundColor(event.Data.BackgroundColor)],
  ['previewFile', () => win.previewFile()],
  ['closeFilePreview', () => win.closeFilePreview()],
  ['setBounds', (event) => win.setBounds(event.Data.Bounds)],
  //
  ['getBounds', () => win.getBounds()],
  ['setContentBounds', (event) => win.setContentBounds(event.Data.Bounds)],
  //
  ['getContentBounds', () => win.getContentBounds()],
  //
  ['getNormalBounds', () => win.getNormalBounds()],
  ['setEnabled', (event) => win.setEnabled(event.Data.Enable)],
  ['setSize', (event) => {
    let height = event.Data.Height;
    let width = event.Data.Width;
    //TODO: animate?
    win.setSize(width, height);
    event.Data = "success"
    return event
  }],
  //
  ['getSize', () => win.getSize()],
  ['setContentSize', (event) => {
    let width = event.Data.Width;
    let height = event.Data.Height;
    win.setContentSize(width, height);
  }],
  //
  ['getContentSize', () => win.getContentSize()],
  ['setMinimumSize', (event) => {
    let width = event.Data.Width;
    let height = event.Data.Height;
    win.setMinimumSize(width, height);
  }],
  //
  ['getMinimumSize', () => win.getMinimumSize()],
  ['setMaximumSize', (event) => {
    let width = event.Data.Width;
    let height = event.Data.Height;
    win.setMaximumSize(width, height);
  }],
  //
  ['getMaximumSize', () => win.getMaximumSize()],
  ['setResizable', (event) => win.setResizable(event.Data.Resizable)],
  //
  ['isResizable', () => win.isResizable()],
  ['setMovable', (event) => win.setMovable(event.Data.Movable)],
  //
  ['isMovable', () => win.isMovable()],
  ['setMinimizable', (event) => win.setMinimizable(event.Data.Minimizable)],
  //
  ['isMinimizable', () => win.isMinimizable()],
  ['setMaximizable', (event) => win.setMaximizable(event.Data.Maximizable)],
  //
  ['isMaximizable', () => win.isMaximizable()],
  ['setFullScreenable', (event) => win.setFullScreenable(event.Data.FullScreenable)],
  //
  ['isFullScreenable', () => win.isFullScreenable()],
  ['setClosable', (event) => win.setClosable(event.Data.Closable)],
  //
  ['isClosable', () => win.isClosable()],
  ['setAlwaysOnTop', (event) => win.setAlwaysOnTop(event.Data.Flag)],
  //
  ['isAlwaysOnTop', () => win.isAlwaysOnTop()],
  ['moveTop', () => win.moveTop()],
  ['center', () => win.center()],
  ['setPosition', (event) => {
    let x = event.Data.X;
    let y = event.Data.Y;
    win.setPosition(x, y)
  }],
  //
  ['getPosition', () => win.getPosition()],
  ['setTitle', (event) => win.setTitle(event.Data.Title)],
  //
  ['getTitle', () => win.getTitle()],
  ['setSheetOffset', (event) => {
    let offsetY = event.Data.OffsetY;
    let offsetX = event.Data.offsetX;
    win.setSheetOffset(offsetY, offsetX);
  }],
  ['flashFrame', (event) => win.flashFrame(event.Data.Flag)],
  ['setSkipTaskbar', (event) => win.setSkipTaskbar(event.Data.Skip)],
  ['setKiosk', (event) => win.setKiosk(event.Data.Flag)],
  //
  ['isKiosk', () => win.isKiosk()],
  //
  ['getNativeWindowHandle', () => win.getNativeWindowHandle()],
  //TODO: How to do callbacks through websocket?
  ['hookWindowMessage', (event) => win.hookWindowMessage(event.Data.Message, event.Data.Callback)],
  //
  ['isWindowMessageHooked', (event) => win.isWindowMessageHooked(event.Data.Message)],
  ['unhookWindowMessage', (event) => win.unhookWindowMessage(event.Data.Message)],
  ['unhookAllWindowMessages', () => win.unhookAllWindowMessages()],
  ['setRepresentedFilename', (event) => win.setRepresentedFilename(event.Data.Filename)],
  //
  ['getRepresentedFilename', () => win.getRepresentedFilename()],
  ['setDocumentEdited', (event) => win.setDocumentEdited(event.Data.Edited)],
  //
  ['isDocumentEdited', () => win.isDocumentEdited()],
  ['focusOnWebView', () => win.focusOnWebView()],
  ['blurWebView', () => win.blurWebView()],
  ['capturePage', (event) => {
    let Rect = event.Data.Rect;
    win.capturePage(Rect, event.Data.Callback);
  }],
  ['loadURL', (event) => {
    let url = event.Data.Url;
    let options = event.Data.Options;
    win.loadURL(url, options);
  }],
  ['loadFile', (event) => {
    let filePath = event.Data.FilePath;
    let options = event.Data.Options;
    win.loadFile(filePath, options);
  }],
  ['reload', () => win.reload()],
  ['setMenu', (event) => win.setMenu(event.Data.Menu)],
  ['setProgressBar', (event) => {
    let progress = event.Data.Progress;
    let options = event.Data.Options;
    win.setProgressBar(progress, options);
  }],
  ['setOverlayIcon', (event) => {
    let overlay = event.Data.Overlay;
    let description = event.Data.Description;
    win.setOverlayIcon(overlay, description);
  }],
  ['setHasShadow', (event) => win.setHasShadow(event.Data.HasShadow)],
  //
  ['hasShadow', () => win.hasShadow()],
  ['setOpacity', (event) => win.setOpacity(event.Data.Opacity)],
  //
  ['getOpacity', () => win.getOpacity()],
  ['setShape', (event) => win.setShape(event.Data.Rects)],
  //
  ['setThumbarButtons', (event) => win.setThumbarButtons(event.Data.Buttons)],
  ['setThumbnailClip', (event) => win.setThumbnailClip(event.Data.Region)],
  ['setThumbnailToolTip', (event) => win.setThumbnailToolTip(event.Data.ToolTip)],
  ['setAppDetails', (event) => win.setAppDetails(event.Data.Options)],
  ['showDefinitionForSelection', () => win.shshowDefinitionForSelectionow()],
  ['setIcon', (event) => win.setIcon(event.Data.Icon)],
  ['setWindowButtonVisibility', (event) => win.setWindowButtonVisibility(event.Data.Visible)],
  ['setAutoHideMenuBar', (event) => win.setAutoHideMenuBar(event.Data.Hide)],
  //
  ['isMenuBarAutoHide', () => win.isMenuBarAutoHide()],
  ['setMenuBarVisibility', (event) => win.setMenuBarVisibility(event.Data.Visible)],
  //
  ['isMenuBarVisible', () => win.isMenuBarVisible()],
  ['setVisibleOnAllWorkspaces', (event) => {
    let visible = event.Data.Visible;
    let options = event.Data.Options;
    win.setVisibleOnAllWorkspaces(visible, options);
  }],
  //
  ['isVisibleOnAllWorkspaces', () => win.isVisibleOnAllWorkspaces()],
  ['isVisibleOnAllWorkspaces', (event) => {
    let ignore = event.Data.Ignore;
    let options = event.Data.Options;
    win.isVisibleOnAllWorkspaces(ignore, options);
  }],
  ['setContentProtection', (event) => win.setContentProtection(event.Data.Enable)],
  ['setFocusable', (event) => win.setFocusable(event.Data.Focusable)],
  ['setParentWindow', (event) => win.setParentWindow(event.Data.Parent)],
  //
  ['getParentWindow', () => win.getParentWindow()],
  //
  ['getChildWindows', () => win.getChildWindows()],
  ['setAutoHideCursor', (event) => win.setAutoHideCursor(event.Data.AutoHide)],
  ['selectPreviousTab', () => win.selectPreviousTab()],
  ['selectNextTab', () => win.selectNextTab()],
  ['mergeAllWindows', () => win.mergeAllWindows()],
  ['moveTabToNewWindow', () => win.moveTabToNewWindow()],
  ['toggleTabBar', () => win.toggleTabBar()],
  ['addTabbedWindow', (event) => win.addTabbedWindow(event.Data.BrowserWindow)],
  ['setVibrancy', (event) => win.setVibrancy(event.Data.Type)],
  ['setTouchBar', (event) => win.setTouchBar(event.Data.TouchBar)],
  ['setBrowserView', (event) => win.setBrowserView(event.Data.BrowserView)],
  //
  ['getBrowserView', () => win.getBrowserView()],
  ['other', () => console.log("other event")],
];

let eventMap = new Map(eventMapArray)

handleEvent = function (event) {
  let handler = eventMap.get(event.Event);
  let result;
  if (handler !== undefined) {
    result = handler(event);
  } else {
    console.log("unknown event " + event.Event)
  }
  return result;
}
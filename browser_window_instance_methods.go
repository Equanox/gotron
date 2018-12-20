package gotron

// Structs used to create JSON messages

type flagStruct struct {
	Flag bool
}

type aspectRatioStruct struct {
	AspectRatio float64
	ExtraSize   string
}

type backgroundColorStruct struct {
	BackgroundColor string
}

type previewFileStruct struct {
	Path        string
	DisplayName []string
}

type setBoundsStruct struct {
	Bounds  Rectangle
	Animate []bool
}

type setSizeStruct struct {
	Width   int
	Height  int
	Animate []bool
}

type setEnabledStruct struct {
	Enable bool
}

type setResizableStruct struct {
	Resizable bool
}

type setMovableStruct struct {
	Movable bool
}

type setMinimizableStruct struct {
	Minimizable bool
}

type setMaximizableStruct struct {
	Maximizable bool
}

type setFullScreenableStruct struct {
	FullScreenable bool
}

type setClosableStruct struct {
	Closable bool
}

type setAlwaysOnTopStruct struct {
	Flag          bool
	Level         string
	RelativeLevel []int
}

type setPositionStruct struct {
	X       int
	Y       int
	Animate []bool
}

type setTitleStruct struct {
	Title string
}

type setSheetOffsetStruct struct {
	OffsetY float64
	OffsetX []float64
}

type setSkipStruct struct {
	Skip bool
}

type hookWindowMessageStruct struct {
	Message  int
	Callback func(param []interface{}) []interface{}
}

type windowMessageStruct struct {
	Message int
}

//OpenDevTools open the devtools view in this BrowserWindow
func (gbw *BrowserWindow) OpenDevTools() {
	res, err := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "devTools", Data: ""})
	if err != nil {
		logger.Debug().Msgf("devTools fail")
	} else {
		logger.Debug().Msgf("devTools success")
		logger.Debug().Msgf("%+v\n", res)
	}
}

//SetHeight sets the height of the electron window to requested size
func (gbw *BrowserWindow) SetHeight(height int) {
	gbw.WindowOptions.Height = height
}

//SetWidth sets the width of the electron window to requested size
func (gbw *BrowserWindow) SetWidth(width int) {
	gbw.WindowOptions.Width = width
}

//Destroy - Force closing the window, the unload and beforeunload event won't be emitted for the web page, and close event will also not be emitted for this window, but it guarantees the closed event will be emitted.
func (gbw *BrowserWindow) Destroy() {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "destroy"})
	logger.Debug().Msgf("%+v\n", res)
}

//Close - Try to close the window. This has the same effect as a user manually clicking the close button of the window. The web page may cancel the close though. See the close event.
func (gbw *BrowserWindow) Close() {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "close"})
	logger.Debug().Msgf("%+v\n", res)
}

//Focus on the window.
func (gbw *BrowserWindow) Focus() {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "focus"})
	logger.Debug().Msgf("%+v\n", res)
}

//Blur - Removes focus from the window.
func (gbw *BrowserWindow) Blur() {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "blur"})
	logger.Debug().Msgf("%+v\n", res)
}

//IsFocused - Returns Boolean - Whether the window is focused.
func (gbw *BrowserWindow) IsFocused() bool {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "isFocused"})
	logger.Debug().Msgf("%+v\n", res)
	return false
}

//IsDestroyed - Returns Boolean - Whether the window is destroyed.
func (gbw *BrowserWindow) IsDestroyed() bool {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "isDestroyed"})
	logger.Debug().Msgf("%+v\n", res)
	return false
}

//Show - Shows and gives focus to the window.
func (gbw *BrowserWindow) Show() {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "show"})
	logger.Debug().Msgf("%+v\n", res)
}

//ShowInactive - Shows the window but doesn't focus on it.
func (gbw *BrowserWindow) ShowInactive() {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "showInactive"})
	logger.Debug().Msgf("%+v\n", res)
}

//Hide - Hides the window.
func (gbw *BrowserWindow) Hide() {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "hide"})
	logger.Debug().Msgf("%+v\n", res)
}

//IsVisible - Returns Boolean - Whether the window is visible to the user.
func (gbw *BrowserWindow) IsVisible() bool {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "isVisible"})
	logger.Debug().Msgf("%+v\n", res)
	return false
}

//IsModal - Returns Boolean - Whether current window is a modal window.
func (gbw *BrowserWindow) IsModal() bool {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "isModal"})
	logger.Debug().Msgf("%+v\n", res)
	return false
}

//Maximize - Maximizes the window. This will also show (but not focus) the window if it isn't being displayed already.
func (gbw *BrowserWindow) Maximize() {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "maximize"})
	logger.Debug().Msgf("%+v\n", res)
}

//UnMaximize - Unmaximizes the window.
func (gbw *BrowserWindow) UnMaximize() {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "unmaximize"})
	logger.Debug().Msgf("%+v\n", res)
}

//IsMaximised - Returns Boolean - Whether the window is maximized.
func (gbw *BrowserWindow) IsMaximised() bool {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "isMaximized"})
	logger.Debug().Msgf("%+v\n", res)
	return false
}

//Minimize - Minimizes the window. On some platforms the minimized window will be shown in the Dock.
func (gbw *BrowserWindow) Minimize() {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "minimize"})
	logger.Debug().Msgf("%+v\n", res)
}

//IsMinimized - Returns Boolean - Whether the window is minimized.
func (gbw *BrowserWindow) IsMinimized() bool {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "isMinimized"})
	logger.Debug().Msgf("%+v\n", res)
	return false
}

//Restore - Restores the window from minimized state to its previous state.
func (gbw *BrowserWindow) Restore() {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "restore"})
	logger.Debug().Msgf("%+v\n", res)
}

//SetFullScreen - Sets whether the window should be in fullscreen mode.
func (gbw *BrowserWindow) SetFullScreen(flag bool) {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "setFullScreen", Data: flagStruct{Flag: flag}})
	logger.Debug().Msgf("%+v\n", res)
}

//IsFullScreen - Returns Boolean - Whether the window is in fullscreen mode.
func (gbw *BrowserWindow) IsFullScreen() bool {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "isFullScreen"})
	logger.Debug().Msgf("%+v\n", res)
	return false
}

//SetSimpleFullScreen - Enters or leaves simple fullscreen mode.
//Simple fullscreen mode emulates the native fullscreen behavior found in versions of Mac OS X prior to Lion (10.7).
func (gbw *BrowserWindow) SetSimpleFullScreen(flag bool) {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "setSimpleFullScreen", Data: flagStruct{Flag: flag}})
	logger.Debug().Msgf("%+v\n", res)
}

//IsSimpleFullScreen - Returns Boolean - Whether the window is in simple (pre-Lion) fullscreen mode.
func (gbw *BrowserWindow) IsSimpleFullScreen() bool {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "isSimpleFullScreen"})
	logger.Debug().Msgf("%+v\n", res)
	return false
}

//IsNormal - Returns Boolean - Whether the window is in normal state (not maximized, not minimized, not in fullscreen mode).
func (gbw *BrowserWindow) IsNormal() bool {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "isNormal"})
	logger.Debug().Msgf("%+v\n", res)
	return false
}

//TODO: find out what the extraSize object is

//SetAspectRatio :
//
// aspectRatio Float - The aspect ratio to maintain for some portion of the content view.
//
// extraSize Size - The extra size not to be included while maintaining the aspect ratio.
//
// This will make a window maintain an aspect ratio. The extra size allows a developer to have space, specified in pixels, not included within the aspect ratio calculations. This API already takes into account the difference between a window's size and its content size.
//
// Consider a normal window with an HD video player and associated controls. Perhaps there are 15 pixels of controls on the left edge, 25 pixels of controls on the right edge and 50 pixels of controls below the player. In order to maintain a 16:9 aspect ratio (standard aspect ratio for HD @1920x1080) within the player itself we would call this function with arguments of 16/9 and [ 40, 50 ]. The second argument doesn't care where the extra width and height are within the content view--only that they exist. Sum any extra width and height areas you have within the overall content view.
//
// Calling this function with a value of 0 will remove any previously set aspect ratios.
func (gbw *BrowserWindow) SetAspectRatio(aspectRatio float64, extraSize string) {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "setAspectRatio", Data: aspectRatioStruct{AspectRatio: aspectRatio, ExtraSize: extraSize}})
	logger.Debug().Msgf("%+v\n", res)
}

//SetBackgroundColor :
//
//backgroundColor String - Window's background color as a hexadecimal value, like #66CD00 or #FFF or #80FFFFFF (alpha is supported if transparent is true). Default is #FFF (white).
//
//Sets the background color of the window.
func (gbw *BrowserWindow) SetBackgroundColor(backgroundColor string) {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "setBackgroundColor", Data: backgroundColorStruct{BackgroundColor: backgroundColor}})
	logger.Debug().Msgf("%+v\n", res)
}

//PreviewFile :
//
//path String - The absolute path to the file to preview with QuickLook. This is important as Quick Look uses the file name and file extension on the path to determine the content type of the file to open.
//
//displayName String (optional) - The name of the file to display on the Quick Look modal view. This is purely visual and does not affect the content type of the file. Defaults to path.
//
//Uses Quick Look to preview a file at a given path.
func (gbw *BrowserWindow) PreviewFile(path string, displayName ...string) {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "isNormal", Data: previewFileStruct{Path: path, DisplayName: displayName}})
	logger.Debug().Msgf("%+v\n", res)
}

//CloseFilePreview - Closes the currently open Quick Look panel.
func (gbw *BrowserWindow) CloseFilePreview() {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "closeFilePreview"})
	logger.Debug().Msgf("%+v\n", res)
}

//SetBounds :
//
//bounds Rectangle
//
//animate Boolean (optional) macOS
//
//Resizes and moves the window to the supplied bounds. Any properties that are not supplied will default to their current values.
func (gbw *BrowserWindow) SetBounds(bounds Rectangle, animate ...bool) {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "setBounds", Data: setBoundsStruct{Bounds: bounds, Animate: animate}})
	logger.Debug().Msgf("%+v\n", res)
}

//GetBounds - Returns Rectangle
func (gbw *BrowserWindow) GetBounds() Rectangle {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "getBounds"})
	logger.Debug().Msgf("%+v\n", res)
	return Rectangle{}
}

//SetContentBounds - Returns Rectangle - Contains the window bounds of the normal state
//
//Note: whatever the current state of the window : maximized, minimized or in fullscreen, this function always returns the position and size of the window in normal state. In normal state, getBounds and getNormalBounds returns the same Rectangle.
func (gbw *BrowserWindow) SetContentBounds(bounds Rectangle, animate ...bool) {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "setContentBounds", Data: setBoundsStruct{Bounds: bounds, Animate: animate}})
	logger.Debug().Msgf("%+v\n", res)
}

//GetContentBounds - Returns Rectangle
func (gbw *BrowserWindow) GetContentBounds() Rectangle {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "getContentBounds"})
	logger.Debug().Msgf("%+v\n", res)
	return Rectangle{}
}

//GetNormalBounds - Returns Rectangle
//
//Note: whatever the current state of the window : maximized, minimized or in fullscreen, this function always returns the position and size of the window in normal state. In normal state, getBounds and getNormalBounds returns the same Rectangle.
func (gbw *BrowserWindow) GetNormalBounds() Rectangle {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "getNormalBounds"})
	logger.Debug().Msgf("%+v\n", res)
	return Rectangle{}
}

//SetEnabled - Disable or enable the window.
func (gbw *BrowserWindow) SetEnabled(enable bool) {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "setEnabled", Data: setEnabledStruct{Enable: enable}})
	logger.Debug().Msgf("%+v\n", res)
}

//SetSize :
//
//Resizes the window to width and height. If width or height are below any set minimum size constraints the window will snap to its minimum size.
func (gbw *BrowserWindow) SetSize(width int, height int, animate ...bool) {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "setSize", Data: setSizeStruct{Width: width, Height: height, Animate: animate}})
	logger.Debug().Msgf("%+v\n", res)
}

//GetSize - Returns Integer[] - Contains the window's width and height.
func (gbw *BrowserWindow) GetSize() []int {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "getSize"})
	logger.Debug().Msgf("%+v\n", res)
	return nil
}

//SetContentSize - Resizes the window's client area (e.g. the web page) to width and height.
func (gbw *BrowserWindow) SetContentSize(width int, height int, animate ...bool) {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "setContentSize", Data: setSizeStruct{Width: width, Height: height, Animate: animate}})
	logger.Debug().Msgf("%+v\n", res)
}

//GetContentSize - Returns Integer[] - Contains the window's client area's width and height.
func (gbw *BrowserWindow) GetContentSize() []int {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "getContentSize"})
	logger.Debug().Msgf("%+v\n", res)
	return nil
}

//SetMinimumSize - Sets the minimum size of window to width and height.
func (gbw *BrowserWindow) SetMinimumSize(width, height int) {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "setMinimumSize", Data: setSizeStruct{Width: width, Height: height}})
	logger.Debug().Msgf("%+v\n", res)
}

//GetMinimumSize - Returns Integer[] - Contains the window's minimum width and height.
func (gbw *BrowserWindow) GetMinimumSize() []int {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "getMinimumSize"})
	logger.Debug().Msgf("%+v\n", res)
	return nil
}

//SetMaximumSize - Sets the maximum size of window to width and height.
func (gbw *BrowserWindow) SetMaximumSize(width, height int) {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "setMaximumSize", Data: setSizeStruct{Width: width, Height: height}})
	logger.Debug().Msgf("%+v\n", res)
}

//GetMaximumSize - Returns Integer[] - Contains the window's maximum width and height.
func (gbw *BrowserWindow) GetMaximumSize() []int {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "getMaximumSize"})
	logger.Debug().Msgf("%+v\n", res)
	return nil
}

//SetResizable - Sets whether the window can be manually resized by user.
func (gbw *BrowserWindow) SetResizable(resizable bool) {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "setResizable", Data: setResizableStruct{Resizable: resizable}})
	logger.Debug().Msgf("%+v\n", res)
}

//IsResizable - Returns Boolean - Whether the window can be manually resized by user.
func (gbw *BrowserWindow) IsResizable() bool {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "isResizable"})
	logger.Debug().Msgf("%+v\n", res)
	return false
}

//SetMovable - Sets whether the window can be moved by user. On Linux does nothing.
func (gbw *BrowserWindow) SetMovable(movable bool) {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "setMovable", Data: setMovableStruct{Movable: movable}})
	logger.Debug().Msgf("%+v\n", res)
}

//IsMovable :
//
// Returns Boolean - Whether the window can be moved by user.
//
//On Linux always returns true.
func (gbw *BrowserWindow) IsMovable() bool {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "isMovable"})
	logger.Debug().Msgf("%+v\n", res)
	return false
}

//SetMinimizable - Sets whether the window can be manually minimized by user. On Linux does nothing.
func (gbw *BrowserWindow) SetMinimizable(minimizable bool) {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "setMinimizable", Data: setMinimizableStruct{Minimizable: minimizable}})
	logger.Debug().Msgf("%+v\n", res)
}

//IsMinimizable :
//
//Returns Boolean - Whether the window can be manually minimized by user
//
// On Linux always returns true.
func (gbw *BrowserWindow) IsMinimizable() bool {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "isMinimizable"})
	logger.Debug().Msgf("%+v\n", res)
	return false
}

//SetMaximizable - Sets whether the window can be manually maximized by user. On Linux does nothing.
func (gbw *BrowserWindow) SetMaximizable(maximizable bool) {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "setMaximizable", Data: setMaximizableStruct{Maximizable: maximizable}})
	logger.Debug().Msgf("%+v\n", res)
}

//IsMaximizable :
//
// Returns Boolean - Whether the window can be manually maximized by user.
//
// On Linux always returns true.
func (gbw *BrowserWindow) IsMaximizable() bool {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "isMaximizable"})
	logger.Debug().Msgf("%+v\n", res)
	return false
}

//SetFullScreenable - Sets whether the maximize/zoom window button toggles fullscreen mode or maximizes the window.
func (gbw *BrowserWindow) SetFullScreenable(fullscreenable bool) {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "setFullScreenable", Data: setFullScreenableStruct{FullScreenable: fullscreenable}})
	logger.Debug().Msgf("%+v\n", res)
}

//IsFullScreenable :
//
//Returns Boolean - Whether the maximize/zoom window button toggles fullscreen mode or maximizes the window.
func (gbw *BrowserWindow) IsFullScreenable() bool {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "isFullScreenable"})
	logger.Debug().Msgf("%+v\n", res)
	return false
}

//SetClosable - Sets whether the window can be manually closed by user. On Linux does nothing.
func (gbw *BrowserWindow) SetClosable(closable bool) {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "setClosable", Data: setClosableStruct{Closable: closable}})
	logger.Debug().Msgf("%+v\n", res)
}

//IsClosable :
//
//Returns Boolean - Whether the window can be manually closed by user.
//
// On Linux always returns true.
func (gbw *BrowserWindow) IsClosable() bool {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "isClosable"})
	logger.Debug().Msgf("%+v\n", res)
	return false
}

//SetAlwaysOnTop :
//
//flag Boolean
//
// level String (optional) macOS - Values include normal, floating, torn-off-menu, modal-panel, main-menu, status, pop-up-menu, screen-saver, and dock (Deprecated). The default is floating. See the macOS docs for more details.
//
// relativeLevel Integer (optional) macOS - The number of layers higher to set this window relative to the given level. The default is 0. Note that Apple discourages setting levels higher than 1 above screen-saver.
//
// Sets whether the window should show always on top of other windows. After setting this, the window is still a normal window, not a toolbox window which can not be focused on.
func (gbw *BrowserWindow) SetAlwaysOnTop(flag bool, level string, relativeLevel ...int) {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "setAlwaysOnTop", Data: setAlwaysOnTopStruct{Flag: flag, Level: level, RelativeLevel: relativeLevel}})
	logger.Debug().Msgf("%+v\n", res)
}

//IsAlwaysOnTop - Returns Boolean - Whether the window is always on top of other windows.
func (gbw *BrowserWindow) IsAlwaysOnTop() bool {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "isAlwaysOnTop"})
	logger.Debug().Msgf("%+v\n", res)
	return false
}

//MoveTop - Moves window to top(z-order) regardless of focus
func (gbw *BrowserWindow) MoveTop() {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "moveTop"})
	logger.Debug().Msgf("%+v\n", res)
}

//Center - Moves window to the center of the screen.
func (gbw *BrowserWindow) Center() {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "center"})
	logger.Debug().Msgf("%+v\n", res)
}

//SetPosition - Moves window to x and y.
func (gbw *BrowserWindow) SetPosition(x, y int, animate ...bool) {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "setPosition", Data: setPositionStruct{X: x, Y: y, Animate: animate}})
	logger.Debug().Msgf("%+v\n", res)
}

//GetPosition - Returns Integer[] - Contains the window's current position.
func (gbw *BrowserWindow) GetPosition() []int {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "getPosition"})
	logger.Debug().Msgf("%+v\n", res)
	return nil
}

//SetTitle - Changes the title of native window to title.
func (gbw *BrowserWindow) SetTitle(title string) {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "setTitle", Data: setTitleStruct{Title: title}})
	logger.Debug().Msgf("%+v\n", res)
}

//GetTitle - Returns String - The title of the native window.
//
// Note: The title of web page can be different from the title of the native window.
func (gbw *BrowserWindow) GetTitle() string {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "getTitle"})
	logger.Debug().Msgf("%+v\n", res)
	return ""
}

//SetSheetOffset :
//
//offsetY Float
//
// offsetX Float (optional)
//
// Changes the attachment point for sheets on macOS. By default, sheets are attached just below the window frame, but you may want to display them beneath a HTML-rendered toolbar. For example:
func (gbw *BrowserWindow) SetSheetOffset(offsetY float64, offsetX ...float64) {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "setSheetOffset", Data: setSheetOffsetStruct{OffsetY: offsetY, OffsetX: offsetX}})
	logger.Debug().Msgf("%+v\n", res)
}

//FlashFrame - Starts or stops flashing the window to attract user's attention.
func (gbw *BrowserWindow) FlashFrame(flag bool) {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "flashFrame", Data: flagStruct{Flag: flag}})
	logger.Debug().Msgf("%+v\n", res)
}

//SetSkipTaskbar - Makes the window not show in the taskbar.
func (gbw *BrowserWindow) SetSkipTaskbar(skip bool) {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "setSkipTaskbar", Data: setSkipStruct{Skip: skip}})
	logger.Debug().Msgf("%+v\n", res)
}

//SetKiosk - Enters or leaves the kiosk mode.
func (gbw *BrowserWindow) SetKiosk(flag bool) {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "setKiosk", Data: flagStruct{Flag: flag}})
	logger.Debug().Msgf("%+v\n", res)
}

//IsKiosk - Returns Boolean - Whether the window is in kiosk mode.
func (gbw *BrowserWindow) IsKiosk() bool {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "isKiosk"})
	logger.Debug().Msgf("%+v\n", res)
	return false
}

//TODO returns OS native handle

//GetNativeWindowHandle :
//
//Returns Buffer - The platform-specific handle of the window.
//
// The native type of the handle is HWND on Windows, NSView* on macOS, and Window (unsigned long) on Linux.
func (gbw *BrowserWindow) GetNativeWindowHandle() {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "getNativeWindowHandle"})
	logger.Debug().Msgf("%+v\n", res)
}

//TODO how to register go callbacks to electron

//HookWindowMessage :
//
//Hooks a windows message. The callback is called when the message is received in the WndProc.
func (gbw *BrowserWindow) HookWindowMessage(message int, callback func([]interface{}) []interface{}) {
	res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "hookWindowMessage", Data: hookWindowMessageStruct{Message: message, Callback: callback}})
	logger.Debug().Msgf("%+v\n", res)
}

//IsWindowMessageHooked - Returns Boolean - true or false depending on whether the message is hooked.
func (gbw *BrowserWindow) IsWindowMessageHooked(message int) bool {
	//res, _ := gbw.sendAndReceiveSocketEvent(SocketEvent{Event: "isWindowMessageHooked", Data: windo}
	panic("Not implemented")
	return false
}

//UnhookWindowMessage - Unhook the window message.
func (gbw *BrowserWindow) UnhookWindowMessage(message int) {
	panic("Not implemented")
}

//UnhookAllWindowMessages - Unhooks all of the window messages.
func (gbw *BrowserWindow) UnhookAllWindowMessages() {
	panic("Not implemented")
}

//SetRepresentedFilename - Sets the pathname of the file the window represents, and the icon of the file will show in window's title bar.
func (gbw *BrowserWindow) SetRepresentedFilename(filename string) {
	panic("Not implemented")
}

//GetRepresentedFilename - Returns String - The pathname of the file the window represents.
func (gbw *BrowserWindow) GetRepresentedFilename() string {
	panic("Not implemented")
	return ""
}

//SetDocumentEdited - Specifies whether the window’s document has been edited, and the icon in title bar will become gray when set to true.
func (gbw *BrowserWindow) SetDocumentEdited(edited bool) {
	panic("Not implemented")
}

//IsDocumentEdited - Returns Boolean - Whether the window's document has been edited.
func (gbw *BrowserWindow) IsDocumentEdited() {
	panic("Not implemented")
}

//FocusOnWebView - Sets focus on webView.
func (gbw *BrowserWindow) FocusOnWebView() {
	panic("Not implemented")
}

//BlurWebView - Unsets focus from webView.
func (gbw *BrowserWindow) BlurWebView() {
	panic("Not implemented")
}

//CapturePage :
//
//rect Rectangle (optional) - The bounds to capture
//
// callback Function
//
// image NativeImage
//
// Same as webContents.capturePage([rect, ]callback).
func (gbw *BrowserWindow) CapturePage(rect []Rectangle, handler func(image string)) {
	panic("Not implemented")
}

//TODO: find out how to handle options object

// LoadURL :
//
// url String
//
// options Object (optional){
//
// httpReferrer (String | Referrer) (optional)  An HTTP Referrer url.
//
// userAgent String (optional)  A user agent originating the request.
//
// extraHeaders String (optional)  Extra headers separated by "\n"
//
// postData (UploadRawData[] | UploadFile[] | UploadBlob[]) (optional)
//
// baseURLForDataURL String (optional)  Base url (with trailing path separator) for files to be loaded by the data url. This is needed only if the specified url is a data url and needs to load other files.}
//
// Same as webContents.loadURL(url[, options]).
//
// The url can be a remote address (e.g. http://) or a path to a local HTML file using the file:// protocol.
//
// To ensure that file URLs are properly formatted, it is recommended to use Node's url.format method:
//
//   let url = require('url').format({
//     protocol: 'file',
//     slashes: true,
//     pathname: require('path').join(__dirname, 'index.html')
//   })
//
//  win.loadURL(url)
// You can load a URL using a POST request with URL-encoded data by doing the following:
//
//  win.loadURL('http://localhost:8000/post', {
//    postData: [{
//      type: 'rawData',
//      bytes: Buffer.from('hello=world')
//    }],
//    extraHeaders: 'Content-Type: application/x-www-form-urlencoded'
//  })
func (gbw *BrowserWindow) LoadURL(url, options interface{}) {
	panic("Not implemented")
}

//TODO: find out how to handle options object

//LoadFile :
//
//filePath String
//
// options Object (optional)
//
// query Object (optional) - Passed to url.format().
//
// search String (optional) - Passed to url.format().
//
// hash String (optional) - Passed to url.format().
//
// Same as webContents.loadFile, filePath should be a path to an HTML file relative to the root of your application. See the webContents docs for more information.
func (gbw *BrowserWindow) LoadFile(filePath string, options interface{}) {
	panic("Not implemented")
}

//Reload - Same as webContents.reload.
func (gbw *BrowserWindow) Reload() {
	panic("Not implemented")
}

//SetMenu - Sets the menu as the window's menu bar, setting it to null will remove the menu bar.
func (gbw *BrowserWindow) SetMenu(menu string) {
	panic("Not implemented")
}

//TODO: find out how to handle options object

//SetProgressBar :
//
//Sets progress value in progress bar. Valid range is [0, 1.0].
//
// Remove progress bar when progress < 0; Change to indeterminate mode when progress > 1.
//
// On Linux platform, only supports Unity desktop environment, you need to specify the *.desktop file name to desktopName field in package.json. By default, it will assume app.getName().desktop.
//
// On Windows, a mode can be passed. Accepted values are none, normal, indeterminate, error, and paused. If you call setProgressBar without a mode set (but with a value within the valid range), normal will be assumed.
func (gbw *BrowserWindow) SetProgressBar(progress float64, options ...interface{}) {
	panic("Not implemented")
}

//SetOverlayIcon - Sets a 16 x 16 pixel overlay onto the current taskbar icon, usually used to convey some sort of application status or to passively notify the user.
func (gbw *BrowserWindow) SetOverlayIcon(overlay string, description string) {
	panic("Not implemented")
}

//SetHasShadow - Sets whether the window should have a shadow. On Windows and Linux does nothing.
func (gbw *BrowserWindow) SetHasShadow(hsShadow bool) {
	panic("Not implemented")
}

//HasShadow - Returns Boolean - Whether the window has a shadow.
//
// On Windows and Linux always returns true.
func (gbw *BrowserWindow) HasShadow() bool {
	panic("Not implemented")
	return false
}

//SetOpacity - Sets the opacity of the window. On Linux does nothing.
func (gbw *BrowserWindow) SetOpacity(opacity float64) {
	panic("Not implemented")
}

//GetOpacity - Returns Number - between 0.0 (fully transparent) and 1.0 (fully opaque)
func (gbw *BrowserWindow) GetOpacity() float64 {
	panic("Not implemented")
	return 0
}

//SetShape :
//
//Setting a window shape determines the area within the window where the system permits drawing and user interaction. Outside of the given region, no pixels will be drawn and no mouse events will be registered. Mouse events outside of the region will not be received by that window, but will fall through to whatever is behind the window.
func (gbw *BrowserWindow) SetShape(rects []Rectangle) {
	panic("Not implemented")
}

//TODO: find out how to handle buttons object

//SetThumbarButtons :
//
// Returns Boolean - Whether the buttons were added successfully
//
// Add a thumbnail toolbar with a specified set of buttons to the thumbnail image of a window in a taskbar button layout. Returns a Boolean object indicates whether the thumbnail has been added successfully.
//
// The number of buttons in thumbnail toolbar should be no greater than 7 due to the limited room. Once you setup the thumbnail toolbar, the toolbar cannot be removed due to the platform's limitation. But you can call the API with an empty array to clean the buttons.
//
// The buttons is an array of Button objects:
//
// Button Object
//
// icon NativeImage - The icon showing in thumbnail toolbar.
//
// click Function
//
// tooltip String (optional) - The text of the button's tooltip.
//
// flags String[] (optional) - Control specific states and behaviors of the button. By default, it is ['enabled'].
//
// The flags is an array that can include following Strings:
//
//
// enabled - The button is active and available to the user.
//
// disabled - The button is disabled. It is present, but has a visual state indicating it will not respond to user action.
//
// dismissonclick - When the button is clicked, the thumbnail window closes immediately.
//
// nobackground - Do not draw a button border, use only the image.
//
// hidden - The button is not shown to the user.
//
// noninteractive - The button is enabled but not interactive; no pressed button state is drawn. This value is intended for instances where the button is used in a notification.
func (gbw *BrowserWindow) SetThumbarButtons(buttons interface{}) bool {
	panic("Not implemented")
	return false
}

//SetThumbnailClip :
//
// Sets the region of the window to show as the thumbnail image displayed when hovering over the window in the taskbar. You can reset the thumbnail to be the entire window by specifying an empty region: { x: 0, y: 0, width: 0, height: 0 }.
func (gbw *BrowserWindow) SetThumbnailClip(region Rectangle) {
	panic("Not implemented")
}

//SetThumbnailToolTip - Sets the toolTip that is displayed when hovering over the window thumbnail in the taskbar.
func (gbw *BrowserWindow) SetThumbnailToolTip(toolTip string) {
	panic("Not implemented")
}

//TODO: find out how to handle options object

//SetAppDetails :
//
// options Object
//
// appId String (optional) - Window's App User Model ID. It has to be set, otherwise the other options will have no effect.
//
// appIconPath String (optional) - Window's Relaunch Icon.
//
// appIconIndex Integer (optional) - Index of the icon in appIconPath. Ignored when appIconPath is not set. Default is 0.
//
// relaunchCommand String (optional) - Window's Relaunch Command.
//
// relaunchDisplayName String (optional) - Window's Relaunch Display Name.
//
// Sets the properties for the window's taskbar button.
//
// Note: relaunchCommand and relaunchDisplayName must always be set together. If one of those properties is not set, then neither will be used.
func (gbw *BrowserWindow) SetAppDetails(options interface{}) {
	panic("Not implemented")
}

//ShowDefinitionForSelection - Same as webContents.showDefinitionForSelection().
func (gbw *BrowserWindow) ShowDefinitionForSelection() {
	panic("Not implemented")
}

//SetIcon - Changes window icon.
func (gbw *BrowserWindow) SetIcon(icon string) {
	panic("Not implemented")
}

//SetWindowButtonVisibility - Sets whether the window traffic light buttons should be visible.
//
// This cannot be called when titleBarStyle is set to customButtonsOnHover.
func (gbw *BrowserWindow) SetWindowButtonVisibility(visible bool) {
	panic("Not implemented")
}

//SetAutoHideMenuBar - Sets whether the window menu bar should hide itself automatically. Once set the menu bar will only show when users press the single Alt key.
//
// If the menu bar is already visible, calling setAutoHideMenuBar(true) won't hide it immediately.
func (gbw *BrowserWindow) SetAutoHideMenuBar(hide bool) {
	panic("Not implemented")
}

//IsMenuBarAutoHide - Returns Boolean - Whether menu bar automatically hides itself.
func (gbw *BrowserWindow) IsMenuBarAutoHide() bool {
	panic("Not implemented")
	return false
}

//SetMenuBarVisibility - Sets whether the menu bar should be visible. If the menu bar is auto-hide, users can still bring up the menu bar by pressing the single Alt key.
func (gbw *BrowserWindow) SetMenuBarVisibility(visible bool) {
	panic("Not implemented")
}

//IsMenuBarVisible - Returns Boolean - Whether the menu bar is visible.
func (gbw *BrowserWindow) IsMenuBarVisible() bool {
	panic("Not implemented")
	return false
}

//TODO: find out how to handle options object

//SetVisibleOnAllWorkspaces - Sets whether the window should be visible on all workspaces.
//
// Note: This API does nothing on Windows.
func (gbw *BrowserWindow) SetVisibleOnAllWorkspaces(visible bool, options ...interface{}) {
	panic("Not implemented")
}

//IsVisibleOnAllWorkspaces - Returns Boolean - Whether the window is visible on all workspaces.
func (gbw *BrowserWindow) IsVisibleOnAllWorkspaces() bool {
	panic("Not implemented")
	return false
}

//TODO: find out how to handle options object

//SetIgnoreMouseEvents - Makes the window ignore all mouse events.
//
// All mouse events happened in this window will be passed to the window below this window, but if this window has focus, it will still receive keyboard events.
func (gbw *BrowserWindow) SetIgnoreMouseEvents(ignore bool, options ...interface{}) {
	panic("Not implemented")
}

//SetContentProtection - Prevents the window contents from being captured by other apps.
//
// On macOS it sets the NSWindow's sharingType to NSWindowSharingNone. On Windows it calls SetWindowDisplayAffinity with WDA_MONITOR.
func (gbw *BrowserWindow) SetContentProtection(enable bool) {
	panic("Not implemented")
}

//SetFocusable - Changes whether the window can be focused.
func (gbw *BrowserWindow) SetFocusable(focusable bool) {
	panic("Not implemented")
}

//TODO: find out how to set parent browserwindow from go

//SetParentWindow - Sets parent as current window's parent window, passing null will turn current window into a top-level window.
func (gbw *BrowserWindow) SetParentWindow(parent *BrowserWindow) {
	panic("Not implemented")
}

//GetParentWindow - Returns BrowserWindow - The parent window.
func (gbw *BrowserWindow) GetParentWindow() BrowserWindow {
	panic("Not implemented")
	return BrowserWindow{}
}

//GetChildWindows - Returns BrowserWindow[] - All child windows.
func (gbw *BrowserWindow) GetChildWindows() []BrowserWindow {
	panic("Not implemented")
	return nil
}

//SetAutoHideCursor - Controls whether to hide cursor when typing.
func (gbw *BrowserWindow) SetAutoHideCursor(autoHide bool) {
	panic("Not implemented")
}

//SelectPreviousTab - Selects the previous tab when native tabs are enabled and there are other tabs in the window.
func (gbw *BrowserWindow) SelectPreviousTab() {
	panic("Not implemented")
}

//SelectNextTab - Selects the next tab when native tabs are enabled and there are other tabs in the window.
func (gbw *BrowserWindow) SelectNextTab() {
	panic("Not implemented")
}

//MergeAllWindows - Merges all windows into one window with multiple tabs when native tabs are enabled and there is more than one open window.
func (gbw *BrowserWindow) MergeAllWindows() {
	panic("Not implemented")
}

//MoveTabToNewWindow - Moves the current tab into a new window if native tabs are enabled and there is more than one tab in the current window.
func (gbw *BrowserWindow) MoveTabToNewWindow() {
	panic("Not implemented")
}

//ToggleTabBar - Toggles the visibility of the tab bar if native tabs are enabled and there is only one tab in the current window.
func (gbw *BrowserWindow) ToggleTabBar() {
	panic("Not implemented")
}

//TODO: find out how to pass this to electron

//AddTabbedWindow - Adds a window as a tab on this window, after the tab for the window instance.
func (gbw *BrowserWindow) AddTabbedWindow(browserwindow *BrowserWindow) {
	panic("Not implemented")
}

//SetVibrancy - Adds a vibrancy effect to the browser window. Passing null or an empty string will remove the vibrancy effect on the window.
func (gbw *BrowserWindow) SetVibrancy(vibrancy string) {
	panic("Not implemented")
}

//TODO: find out how to handle touchbar object

//SetTouchBar - Sets the touchBar layout for the current window. Specifying null or undefined clears the touch bar. This method only has an effect if the machine has a touch bar and is running on macOS 10.12.1+.
//
// Note: The TouchBar API is currently experimental and may change or be removed in future Electron releases.
func (gbw *BrowserWindow) SetTouchBar(touchBar interface{}) {
	panic("Not implemented")
}

//TODO: find out how to handle browserView object

//SetBrowserView - experimental
func (gbw *BrowserWindow) SetBrowserView(browserView interface{}) {
	panic("Not implemented")
}

//TODO: find out how to handle browserView object

//GetBrowserView - experimental
func (gbw *BrowserWindow) GetBrowserView() interface{} {
	panic("Not implemented")
	return nil
}

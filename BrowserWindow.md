# BrowserWindow

> Create and control browser windows.

## Instance Events

To be implemented.

## Static Methods

To be implemented.

## Instance Properties

To be implemented

## Instance Methods

Objects created with `win := gotron.New()` have the following instance methods. These Methods will only have effect after running `win.Start()`. This is just a list to show implemented and unimplemented methods from electron's BrowserWindow. See [browser-window](https://github.com/electron/electron/blob/master/docs/api/browser-window.md) for full documentation of electron BrowserWindow.

**Note:** Some methods are only available on specific operating systems and are
labeled as such.

### Implemented
* `win.OpenDevTools()`
* `win.SetSize(width int, height int, animate ...bool)`
* `win.Close()`
* `win.Show()`
* `win.Hide()`
* `win.Minimize()`
* `win.Restore()`
* `win.Maximize()`
* `win.UnMaximize()`
* `win.SetFullScreen(flag bool)`

### Not (/partly) implemented

* `win.Destroy()`
* `win.Focus()`
* `win.Blur()`
* `win.IsFocused() bool`
* `win.IsDestroyed() bool`
* `win.ShowInactive()`
* `win.IsVisible() bool`
* `win.IsModal() bool`
* `win.IsMaximised() bool`
* `win.IsMinimized() bool`
* `win.IsFullScreen() bool`
* `win.SetSimpleFullScreen(flag bool)`
* `win.IsSimpleFullScreen() bool`
* `win.IsNormal() bool`
* `win.SetAspectRatio(aspectRatio float64, extraSize string)`
* `win.SetBackgroundColor(backgroundColor string)`
* `win.PreviewFile(path string, displayName ...string)`
* `win.CloseFilePreview()`
* `win.SetBounds(bounds Rectangle, animate ...bool)`
* `win.GetBounds() Rectangle`
* `win.SetContentBounds(bounds Rectangle, animate ...bool)`
* `win.GetContentBounds() Rectangle`
* `win.GetNormalBounds() Rectangle`
* `win.SetEnabled(enable bool)`
* `win.GetSize() []int`
* `win.SetContentSize(width int, height int, animate ...bool)`
* `win.GetContentSize() []int`
* `win.SetMinimumSize(width, height int)`
* `win.GetMinimumSize() []int`
* `win.SetMaximumSize(width, height int)`
* `win.GetMaximumSize() []int`
* `win.SetResizable(resizable bool)`
* `win.IsResizable() bool`
* `win.SetMovable(movable bool)`
* `win.IsMovable() bool`
* `win.SetMinimizable(minimizable bool)`
* `win.IsMinimizable() bool`
* `win.SetMaximizable(maximizable bool)`
* `win.IsMaximizable() bool`
* `win.SetFullScreenable(fullscreenable bool)`
* `win.IsFullScreenable() bool`
* `win.SetClosable(closable bool)`
* `win.IsClosable() bool`
* `win.SetAlwaysOnTop(flag bool, level string, relativeLevel ...int)`
* `win.IsAlwaysOnTop() bool`
* `win.MoveTop()`
* `win.Center()`
* `win.SetPosition(x, y int, animate ...bool)`
* `win.GetPosition() []int`
* `win.SetTitle(title string)`
* `win.GetTitle() string`
* `win.SetSheetOffset(offsetY float64, offsetX ...float64)`
* `win.FlashFrame(flag bool)`
* `win.SetSkipTaskbar(skip bool)`
* `win.SetKiosk(flag bool)`
* `win.IsKiosk() bool`
* `win.GetNativeWindowHandle()`
* `win.HookWindowMessage(message int, callback func([]interface{}) []interface{})`
* `win.IsWindowMessageHooked(message int) bool`
* `win.UnhookWindowMessage(message int)`
* `win.UnhookAllWindowMessages()`
* `win.GetRepresentedFilename() string`
* `win.SetDocumentEdited(edited bool)`
* `win.IsDocumentEdited()`
* `win.FocusOnWebView()`
* `win.BlurWebView()`
* `win.CapturePage(rect []Rectangle, handler func(image string))`
* `win.LoadURL(url, options interface{})`
* `win.LoadFile(filePath string, options interface{})`
* `win.Reload()`
* `win.SetMenu(menu string)`
* `win.SetProgressBar(progress float64, options ...interface{})`
* `win.SetOverlayIcon(overlay string, description string)`
* `win.SetHasShadow(hsShadow bool)`
* `win.HasShadow() bool`
* `win.SetOpacity(opacity float64)`
* `win.GetOpacity() float64`
* `win.SetShape(rects []Rectangle)`
* `win.SetThumbarButtons(buttons interface{}) bool`
* `win.SetThumbnailClip(region Rectangle)`
* `win.SetThumbnailToolTip(toolTip string)`
* `win.SetAppDetails(options interface{})`
* `win.ShowDefinitionForSelection()`
* `win.SetIcon(icon string)`
* `win.SetWindowButtonVisibility(visible bool)`
* `win.SetAutoHideMenuBar(hide bool)`
* `win.IsMenuBarAutoHide() bool`
* `win.SetMenuBarVisibility(visible bool)`
* `win.IsMenuBarVisible() bool`
* `win.SetVisibleOnAllWorkspaces(visible bool, options ...interface{})`
* `win.IsVisibleOnAllWorkspaces() bool`
* `win.SetIgnoreMouseEvents(ignore bool, options ...interface{})`
* `win.SetContentProtection(enable bool)`
* `win.SetFocusable(focusable bool)`
* `win.SetParentWindow(parent *BrowserWindow)`
* `win.GetParentWindow() BrowserWindow`
* `win.GetChildWindows() []BrowserWindow`
* `win.SetAutoHideCursor(autoHide bool)`
* `win.SelectPreviousTab()`
* `win.SelectNextTab()`
* `win.MergeAllWindows()`
* `win.MoveTabToNewWindow()`
* `win.ToggleTabBar()`
* `win.AddTabbedWindow(browserwindow *BrowserWindow)`
* `win.SetVibrancy(vibrancy string)`
* `win.SetTouchBar(touchBar interface{})`
* `win.SetBrowserView(browserView interface{})`
* `win.GetBrowserView() interface{}`
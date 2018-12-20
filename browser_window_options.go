package gotron

import (
	"errors"
	"time"
)

//WindowOptions - all possible electron browserwindow options
type WindowOptions struct {
	Width                  int            `json:"width,omitempty"`
	Height                 int            `json:"height,omitempty"`
	X                      int            `json:"x,omitempty"`
	Y                      int            `json:"y,omitempty"`
	UseContentSize         bool           `json:"useContentSize,omitempty"`
	Center                 bool           `json:"center,omitempty"`
	MinWidth               int            `json:"minWidth,omitempty"`
	MinHeight              int            `json:"minHeight,omitempty"`
	MaxWidth               int            `json:"maxWidth,omitempty"`
	MaxHeight              int            `json:"maxHeight,omitempty"`
	Resizable              bool           `json:"resizable"`
	Movable                bool           `json:"movable"`
	Minimizable            bool           `json:"minimizable"`
	Maximizable            bool           `json:"maximizable"`
	Closable               bool           `json:"closable"`
	Focusable              bool           `json:"focusable"`
	AlwaysOnTop            bool           `json:"alwaysOnTop,omitempty"`
	FullScreen             bool           `json:"fullscreen,omitempty"`
	Fullscreenable         bool           `json:"fullscreenable"`
	SimpleFullscreen       bool           `json:"simpleFullscreen,omitempty"`
	SkipTaskbar            bool           `json:"skipTaskbar,omitempty"`
	Kiosk                  bool           `json:"kiosk,omitempty"`
	Title                  string         `json:"title,omitempty"`
	Icon                   string         `json:"icon,omitempty"`
	Show                   bool           `json:"show"`
	Frame                  bool           `json:"frame"`
	Parent                 interface{}    `json:"parent,omitempty"` // default is null, how to obtain parent windo handle?
	Modal                  bool           `json:"modal,omitempty"`
	AcceptFirstMouse       bool           `json:"acceptFirstMouse,omitempty"`
	DisableAutoHideCursor  bool           `json:"disableAutoHideCursor,omitempty"`
	AutoHideMenuBar        bool           `json:"autoHideMenuBar,omitempty"`
	EnableLargerThanScreen bool           `json:"enableLargerThanScreen,omitempty"`
	BackGroundColor        string         `json:"backgroundColo,omitempty"`
	HasShadow              bool           `json:"hasShadow"`
	Opacity                float64        `json:"opacity,omitempty"`
	DarkTheme              bool           `json:"darkTheme,omitempty"`
	TransParent            bool           `json:"transparent,omitempty"`
	Type                   string         `json:"type,omitempty"`
	TitleBarStyle          string         `json:"titleBarStyle,omitempty"`
	FullscreenWindowTitle  bool           `json:"fullscreenWindowTitle,omitempty"`
	ThickFrame             bool           `json:"thickFrame"`
	Vibrancy               string         `json:"vibrancy,omitempty"`
	ZoomToPageWidth        bool           `json:"zoomToPageWidth,omitempty"`
	TabbingIdentifier      string         `json:"tabbingIdentifier,omitempty"`
	WebPreferences         WebPreferences `json:"webPreferences"`
}

//WebPreferences - webpreferences for electron browserwindow
type WebPreferences struct {
	DevTools                    bool        `json:"devTools"`
	NodeIntegration             bool        `json:"nodeIntegration"`
	NodeIntegratonInWorker      bool        `json:"nodeIntegrationInWorker,omitempty"`
	Preload                     string      `json:"preload,omitempty"`
	Sandbox                     bool        `json:"sandbox,omitempty"` //Experimental
	EnableRemoteModule          bool        `json:"enableRemoteModule"`
	Session                     interface{} `json:"session,omitempty"` //TODO: Find out how this can be passed from go to electron
	Partition                   string      `json:"partition,omitempty"`
	Affinity                    string      `json:"affinity,omitempty"`
	ZoomFactor                  float64     `json:"zoomFactor,omitempty"`
	Javascript                  bool        `json:"javascript"`
	WebSecurity                 bool        `json:"webSecurity"`
	AllowRunningInsecureContent bool        `json:"allowRunningInsecureContent,omitempty"`
	Images                      bool        `json:"images"`
	TextAreasAreResizable       bool        `json:"textAreasAreResizable"`
	Webgl                       bool        `json:"webgl"`
	Webaudio                    bool        `json:"webaudio"`
	Plugins                     bool        `json:"plugins,omitempty"`
	ExperimentalFeatures        bool        `json:"experimentalFeatures,omitempty"`
	ScrollBounce                bool        `json:"scrollBounce,omitempty"`
	EnableBlinkFeatures         string      `json:"enableBlinkFeatures,omitempty"`
	DisableBlinkFeatures        string      `json:"disableBlinkFeatures,omitempty"`
	DefaultFontFamily           string      `json:"defaultFontFamily,omitempty"`
	DefaultFontSize             int         `json:"defaultFontSize,omitempty"`
	DefaultMonospaceFontSize    int         `json:"defaultMonospaceFontSize,omitempty"`
	MinimumFontSize             int         `json:"minimumFontSize,omitempty"`
	DefaultEncoding             int         `json:"defaultEncoding,omitempty"`
	BackgroundThrottling        bool        `json:"backgroundThrottling"`
	Offscreen                   bool        `json:"offscreen,omitempty"`
	ContextIsolation            bool        `json:"contextIsolation,omitempty"`
	NativeWindowOpen            bool        `json:"nativeWindowOpen,omitempty"`
	WebviewTag                  bool        `json:"webviewTag"`
	AdditionalArguments         []string    `json:"additionalArguments,omitempty"`
	SafeDialogs                 bool        `json:"safeDialogs,omitempty"`
	SafeDialogsMessage          string      `json:"safeDialogsMessage,omitempty"`
	NavigateOnDragDrop          bool        `json:"navigateOnDragDrop,omitempty"`
}

//Rectangle basic rectangle object
type Rectangle struct {
	X      int
	Y      int
	Width  int
	Height int
}

//sendAndReceiveSocketEvent
//send a SocketEvent message to given gotron-browser-window instance
//wait for response SocketEvent (blocking or 1s timeout) and return received SocketEvent or error
func (gbw *BrowserWindow) sendAndReceiveSocketEvent(event SocketEvent) (SocketEvent, error) {
	var res SocketEvent
	if !gbw.Running {
		return res, errors.New("No BrowserWindow instance running")
	}
	c := make(chan SocketEvent)
	gbw.optionsQueue <- optionsQueueElement{Waiter: c, Data: event}
	select {
	case res = <-c:
		return res, nil
	case <-time.After(1 * time.Second):
		return res, errors.New("Websocket: Wait timeout exceeded")
	}
}

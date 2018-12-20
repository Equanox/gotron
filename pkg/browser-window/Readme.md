# gotron-browser-window

Rampup electron from golang using only a go api.

## Usage

To create a new gotron-browser-window use

    import( gbw "github.com/Benchkram/gotron-browser-window" )

    window, err := gbw.New(pathToIndexjs, pathToCSS, appFolder)
    
    if err != nil {
		Panic(err)
	}

where pathToIndexjs and pathToCSS point to your frontend code and styling. appFolder is where the gotron-browser-window application should be created.

To bring up the frontend then run

    done, err := window.Start()
    
    if err != nil {
		Panic(err)
	}

Use

    <-done

to wait for the window to be closed.

# Browser Window Methods

Gotron-browser-window supports electron browserwindow methods to be called. Note that some calls are only possible before or after bringing up the frontend.

See [BrowserWindow.md](BrowserWindow.md) for list of implemented and unimplemented methods.

See [Electron BrowserWindow Documentation](https://github.com/electron/electron/blob/master/docs/api/browser-window.md) for full list of methods from electron's BrowserWindow.

## Tasks

* [x] electron rebuild
* [x] forceinstall
* ~~[ ] electron start OS specific~~
* [x] Start Methode aufdröseln
* [x] shutdown close all go routines
* [x] MD File für GBW Methods -> verlinken in Readme.md
* [x] gotronbrowserwindow -> gotron
* [ ] (zerolog) - partly
* [x] errz
* [x] rename go files to snake case
package gotron

import (
	"testing"
)

func TestCopyElectronApplication(t *testing.T) {

	window, err := New()
	if err != nil {
		t.Fatal(err)
	}

	window.WindowOptions.Width = 1200
	window.WindowOptions.Height = 600

	done, err := window.Start()
	if err != nil {
		t.Fatal(err)
	}

	<-done
}

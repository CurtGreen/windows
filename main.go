package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/zserge/webview"
)

type config struct {
	Width  int
	Height int
	Title  string
	Msg    string
}

func main() {
	// Pull in arguments from os package
	// width, height, title, msg ( 1 , 2 , 3 , 4)
	newWindow := config{}
	if len(os.Args) < 5 {
		newWindow.Height = 400
		newWindow.Width = 400
		newWindow.Title = "Default Title"
		newWindow.Msg = "Some Default Message..."
	} else {
		if arg, err := strconv.Atoi(os.Args[1]); err == nil {
			newWindow.Width = arg
		}
		if arg, err := strconv.Atoi(os.Args[2]); err == nil {
			newWindow.Height = arg
		}

		if arg := os.Args[3]; arg != "" {
			newWindow.Title = arg
		}

		if arg := os.Args[4]; arg != "" {
			newWindow.Msg = arg
		}
	}

	go func() {
		runDataURL(newWindow.Width, newWindow.Height, newWindow.Title, newWindow.Msg)
	}()

	fmt.Println("Immediately back in main, hopefully window will be closed in 4 seconds")
	time.Sleep(4 * time.Second)
}

func runDataURL(windowWidth int, windowHeight int, windowTitle string, windowMsg string) {

	w := webview.New(webview.Settings{
		Width:  windowWidth,
		Height: windowHeight,
		Title:  windowTitle,
		URL:    "data:text/html," + url.PathEscape(windowMsg),
	})
	w.Run()
	defer w.Exit()

}

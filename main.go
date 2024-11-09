package main

import (
	"os"
	"path/filepath"

	webview "github.com/webview/webview_go"
)

func main() {
	path, err := filepath.Abs("layout/index.html")
	if err != nil {
		os.Exit(1)
	}
	path = "file://" + path

	w := webview.New(true)
	defer w.Destroy()

	w.SetTitle("Chat Unilab")
	w.SetSize(800, 600, 0)
	w.Navigate(path)

	w.Bind("send", func(message string) string {
		// w.Eval(fmt.Sprintf("postMyMessage('%s')", message))
		return message
	})

	w.Run()
}

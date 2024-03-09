package main

import webview "github.com/webview/webview_go"
import "embed"
import "fmt"
import "io/fs"
import "net/http"

//go:embed public/*
var public_fs embed.FS

func main() {

	filesystem, _ := fs.Sub(public_fs, "public")

	fmt.Println("Listening on http://localhost:13337")

	go func() {

		fmt.Println("Opening WebView...")

		view := webview.New(true)
		defer view.Destroy()
		view.SetTitle("Agenda")
		view.SetSize(800, 600, webview.HintNone)
		view.Navigate("http://localhost:13337/index.html")
		view.Run()

	}()

	fileserver := http.FileServer(http.FS(filesystem))
	http.Handle("/", fileserver)

	http.HandleFunc("/api/user/save", func(response http.ResponseWriter, request *http.Request) {
		fmt.Println("/api/user/save called!")
	})

	err := http.ListenAndServe(":13337", nil)

	if err == nil {
		fmt.Println("Stopped listening on http://localhost:13337")
	}

}


package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"

	webview "github.com/webview/webview_go"
)

const porta = ":9000"

func main() {
	path, err := filepath.Abs("layout/index.html")
	if err != nil {
		os.Exit(1)
	}
	path = "file://" + path

	w := webview.New(true)
	defer w.Destroy()

	w.SetTitle("Chat Unilab")
	w.SetSize(800, 500, 0)
	w.Navigate(path)

	w.Bind("getIp", func(ip string) {
		go client(w, ip)
	})

	server(w)

	w.Run()
}

func server(w webview.WebView) {
	server, err := net.Listen("tcp", porta)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := server.Accept()

		if err != nil {
			log.Fatal(err)
		}

		go read(w, conn)
	}
}

func read(w webview.WebView, conn net.Conn) {
	for {
		var buffer = make([]byte, 256)
		nbytes, err := conn.Read(buffer)

		if err != nil {
			log.Fatal(err)
		}

		if nbytes > 0 {
			// Nesta parte devo colocar a mensagem com js
			// fmt.Println(conn.RemoteAddr().String() + ">=< : " + string(buffer))
			w.Eval(fmt.Sprintf("postUserMessage('%s')", string(buffer)))
		}

	}
}

func client(w webview.WebView, ip string) {
	conn, err := net.Dial("tcp", ip+porta)

	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	w.Bind("send", func(msg string) {
		_, err = conn.Write([]byte(msg))

		if err != nil {
			log.Fatal(err)
		}
	})

}

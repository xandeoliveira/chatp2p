package main

import (
	"fmt"
	"log"
	"net"

	"chatunilab/cipher"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

const porta = ":9000"

var k int = 6
var username string
var ipDestination string

var messagesContainer = container.New(layout.NewVBoxLayout())

func main() {
	a := app.New()
	// Criando as páginas
	welcome := a.NewWindow("Bem vindo!")
	destination := a.NewWindow("Destino")
	wchat := a.NewWindow("Chat Unilab")
	go server()

	// Widgets das páginas
	title := widget.NewLabelWithStyle("BEM VINDO AO CHAT UNILAB!", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Digite seu nome aqui")

	destinationTitle := widget.NewLabelWithStyle("Iniciar chat.", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	ipEntry := widget.NewEntry()
	ipEntry.SetPlaceHolder("Digite o ip de destino aqui")

	// Window Welcome
	welcomeContainer := container.New(
		layout.NewVBoxLayout(),
		layout.NewSpacer(),

		container.New(
			layout.NewVBoxLayout(),
			title,
			nameEntry,
			widget.NewButton("Confirmar", func() {
				username = nameEntry.Text

				destination.Show()
				welcome.Hide()
			}),
		),

		layout.NewSpacer(),
	)

	// Window Destination
	destinationContainer := container.New(
		layout.NewVBoxLayout(),
		layout.NewSpacer(),

		container.New(
			layout.NewVBoxLayout(),
			destinationTitle,
			ipEntry,
			widget.NewButton("Confirmar", func() {
				ipDestination = ipEntry.Text

				destination.Hide()
				wchat.Show()

				client(getIp(), wchat)
			}),
		),

		layout.NewSpacer(),
	)

	// Window Chat

	// Configurando tamanhos das páginas
	welcome.SetContent(welcomeContainer)
	welcome.Resize(fyne.NewSize(500, 400))

	destination.SetContent(destinationContainer)
	destination.Resize(fyne.NewSize(500, 400))

	// Iniciando a interface
	welcome.ShowAndRun()
}

func getUsername() string {
	return username
}

func getIp() string {
	return ipDestination
}

func server() {
	server, err := net.Listen("tcp", porta)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Servidor: " + GetOutboundIp().String())

	for {
		conn, err := server.Accept()

		if err != nil {
			log.Fatal(err)
		}

		go read(conn)
	}
}

func read(conn net.Conn) {
	for {
		var buffer = make([]byte, 256)
		nbytes, err := conn.Read(buffer)

		if err != nil {
			log.Fatal(err)
		}

		if nbytes > 0 {
			message := container.New(
				layout.NewHBoxLayout(),

				widget.NewLabelWithStyle(fmt.Sprintf("%s:", conn.RemoteAddr().String()), fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
				widget.NewLabel(cipher.Decoding(string(buffer), k)),
			)

			messagesContainer.Add(message)
		}

	}
}

func client(ip string, wchat fyne.Window) {
	fmt.Println("Concetando: " + ip + porta)
	conn, err := net.Dial("tcp", ip+porta)

	if err != nil {
		log.Fatal(err)
	}

	messageEntry := widget.NewEntry()
	messageEntry.SetPlaceHolder("Mensagem")

	formContainer := container.New(
		layout.NewHBoxLayout(),

		container.New(layout.NewGridWrapLayout(fyne.NewSize(500, messageEntry.MinSize().Height)), messageEntry),

		widget.NewButton("Enviar", func() {
			message := container.New(
				layout.NewHBoxLayout(),

				widget.NewLabelWithStyle(fmt.Sprintf("%s:", getUsername()), fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
				widget.NewLabel(messageEntry.Text),
			)

			_, err := conn.Write([]byte(cipher.Encoding(messageEntry.Text, k)))

			if err != nil {
				log.Fatal(err)
			}

			messagesContainer.Add(message)
			messageEntry.SetText("")
		}),
	)

	chatContainer := container.New(
		layout.NewVBoxLayout(),

		container.New(
			layout.NewVBoxLayout(),

			container.New(
				layout.NewGridWrapLayout(fyne.NewSize(500, 500-(messageEntry.MinSize().Height))),
				messagesContainer,
			),

			layout.NewSpacer(),

			formContainer,
		),
	)

	wchat.SetContent(chatContainer)
	wchat.Resize(fyne.NewSize(500, 500))

}

func GetOutboundIp() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	localAdd := conn.LocalAddr().(*net.UDPAddr)

	return localAdd.IP
}

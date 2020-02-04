package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
)

func serve(ctx context.Context, addr string) {
	server := &http.Server{
		Addr:    addr,
		Handler: http.FileServer(http.Dir("web/static")),
	}
	go func() {
		server.ListenAndServe()
	}()
	go func() {
		<-ctx.Done()
		server.Close()
	}()
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	serve(ctx, "127.0.0.1:34544")

	logger := log.New(os.Stderr, "", 0)
	app, err := astilectron.New(logger, astilectron.Options{
		AppName:            "Kryer GUI",
		VersionAstilectron: "0.33.0",
		VersionElectron:    "4.0.1",
	})
	if err != nil {
		log.Fatalf(err.Error())
	}

	if err := app.Start(); err != nil {
		log.Fatalf(err.Error())
	}
	defer app.Close()

	window, err := app.NewWindow("http://127.0.0.1:34544", &astilectron.WindowOptions{
		Center:         astikit.BoolPtr(true),
		Fullscreenable: astikit.BoolPtr(false),
		Height:         astikit.IntPtr(600),
		Width:          astikit.IntPtr(800),
		Resizable:      astikit.BoolPtr(false),
	})
	if err != nil {
		log.Fatalf(err.Error())
	}

	window.OnMessage(func(m *astilectron.EventMessage) interface{} {
		var message string
		m.Unmarshal(&message)

		time.Sleep(time.Second * 10)
		fmt.Println("Done handling message", message)
		return nil
	})

	if err := window.Create(); err != nil {
		log.Fatalf(err.Error())
	}

	app.Wait()
}

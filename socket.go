package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/zishang520/engine.io/types"
	"github.com/zishang520/socket.io/socket"
)

func SocketServer() {
	serverOptions := socket.DefaultServerOptions()
	serverOptions.SetCors(&types.Cors{
		Origin: "*",
	})

	io := socket.NewServer(nil, serverOptions)

	http.Handle("/socket.io/", io.ServeHandler(nil))

	go http.ListenAndServe(":8000", nil)

	io.On("connection", func(clients ...any) {
		fmt.Println("you've connected")
		client := clients[0].(*socket.Socket)
		client.On("event", func(datas ...any) {
		})
		client.On("disconnect", func(...any) {
		})
	})

	exit := make(chan struct{})
	SignalC := make(chan os.Signal)

	signal.Notify(SignalC, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		for s := range SignalC {
			switch s {
			case os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				close(exit)
				return
			}
		}
	}()

	<-exit
	io.Close(nil)
	os.Exit(0)
}

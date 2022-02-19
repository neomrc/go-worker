package main

import (
	"fmt"

	"github.com/neomrc/go-worker/src/worker"
)

func main() {
	handler := worker.NewHandler(func() {
		fmt.Println("Hello, world!")
	})
	handler.SetInterval(1000)

	handler2 := worker.NewHandler(func() {
		fmt.Println("Hello, world!2")
	})
	handler2.SetInterval(500)

	worker := worker.NewWorker()
	worker.AddHandler("handler", handler)
	worker.AddHandler("handler2", handler2)
	worker.Start()
}

# go-worker

Creates a simple worker function utilizing go routines and channels

## Setup
```
go get github.com/neomrc/go-worker
```

## Usage

```
import "github.com/neomrc/go-worker/src/worker"

// creates the handler
handler := worker.NewHandler(func() {
  fmt.Println("Hello, world!")
})
// set the interval for polling
handler.SetInterval(1000)

// initializes the worker instance
worker := worker.NewWorker()
// append the worker handler
worker.AddHandler("handler", handler)
// start the worker
worker.Start()
```

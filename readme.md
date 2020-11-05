# Service Bootstrap

This is a bootstrap library for creating services so that common code does not need to be
re-implemented each time, and we can spend our time working on just the business logic.

## Third party libraries

The third party libraries we are using are:

| Library     | Description                                    | Link                                                             |
| ----------- | ---------------------------------------------- | ---------------------------------------------------------------- |
| Cobra       | Library for creating CLI applications          | [https://github.com/spf13/cobra](https://github.com/spf13/cobra) |
| Viper       | Library for application configuration          | [https://github.com/spf13/viper](https://github.com/spf13/viper) |
| Zap         | Library for Structured, pluggable logging      | [https://github.com/uber-go/zap](https://github.com/uber-go/zap) |

## Pre-requisites

This library requires Go version 1.11+ as it leverages Go Modules

## Usage

To use this library, create a new project and initialize the Go module

```bash
mkdir my-go-webapp
cd my-go-webapp
go mod init my-go-webapp
```

Use go get to add this project as a dependency to your Go module
```bash
go get github.com/birchwood-langham/bootstrap/v1
```

### Init and cleanup functions

The bootstrap allows you to define and add init and cleanup functions that will be run sequentially in the order they were
added by the bootstrap.

The init and cleanup functions are defined as

```
type InitFunc func(context.Context, service.StateStore) error
type CleanupFunc func() error
```

To use the bootstrap, define your initialization functions and add them to the application.

#### Example

```go
package main

import (
    "context"
    "fmt"

    "github.com/birchwood-langham/bootstrap/v1/pkg/service"
    "github.com/birchwood-langham/bootstrap/v1/pkg/cmd"
)

func initApp(ctx context.Context, store service.StateStore) error {
    fmt.Println("Hello, World!")
    
    childCtx, cancel := context.WithCancel(ctx)

    go func(c context.Context) {
        // long running process here
        for {
            select {
            case <- c.Done():
                return
            }
        }   
    }(childCtx)

    store = store.WithValue("stop-long-running-process", cancel)

    return nil
}

func main() {
    app := service.NewApplication().
		AddInitFunc(initApp)

	app.SetProperties("usage message", "short description", "long description for the application to be displayed when run with the help flag")

    // NewStateStore() implements the StateStore interface and creates a non-persistent thread safe in-memory state store
	cmd.Execute(context.Background(), app, service.NewStateStore())
}
```
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
type CleanupFunc func(service.StateStore) error
```

To use the bootstrap, define your initialization and cleanup functions and add them to the application.

### State

To allow the application to store state, you can implement the service.StateStore interface. This will allow you to utilise
whatever state store is necessary for your service whether that is in-memory storage or using a database like redis or distributed
service like Etcd or Consul.

### Configuration file

The bootstrap expects a configuration file with the following settings as a minimum:

```yaml
version: 0.1.0
service:
    name: bootstrap
log:
    filepath: ./logs/bootstrap.log
    level: DEBUG
    max-size: 100
    max-backup: 5
    max-age: 30
    compress: false
```

The configuration file must be called configuration.<ext> where ext is any format supported by viper.

You can add your own configuration to the file and access them using viper.

### CLI commands

To add your own CLI commands, you can just create a command, and add them before calling the `cmd.Execute()` function. For example:

```go
package main

// omitted for clarity

var helloCmd = &cobra.Command{
    Use: "hello",
    Short: "Says hello",
    Long: "The obligatory hello world function",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Hello, World!")
    },
}

func main() {
    cmd.AddCommand(helloCmd)
    cmd.Execute(New())
}
```

## Configuration

To make accessing configuration easier, a configuration wrapper function is available in the `github.com/birchwood-langham/bootstrap/v1/pkg/config`
package.

### Example

Instead of doing something like this:

```go

    serviceName := "Default Service Name"
    
    if viper.IsSet("service.name") {
    	serviceName = viper.Get("service.name")
    }

```

We can use:

```go
    serviceName := config.Get("service", "name").String("Default Service Name")
```

### Supported Data Types

The following table contains the translation between the viper function signatures and the config functions we have defined.

| Viper Function                   | Config Function                                          | Return Data Type       |
| -------------------------------- | -------------------------------------------------------- | ---------------------- |
| viper.Get(string)                | config.Get(...string).Value(interface{})                 | interface{}            |
| viper.GetBool(string)            | config.Get(...string).Bool(bool)                         | bool                   |
| viper.GetFloat64(string)         | config.Get(...string).Float64(float64)                   | float64                |
| viper.GetInt(string)             | config.Get(...string).Int(int)                           | int                    |
| viper.GetInt8(string)            | config.Get(...string).Int8(int)                          | int8                   |
| viper.GetInt16(string)           | config.Get(...string).Int16(int)                         | int16                  |
| viper.GetInt32(string)           | config.Get(...string).Int32(int)                         | int32                  |
| viper.GetInt64(string)           | config.Get(...string).Int64(int)                         | int64                  |
| viper.GetUint(string)            | config.Get(...string).Uint(int)                          | uint                   |
| viper.GetUint8(string)           | config.Get(...string).Uint8(int)                         | uint8                  |
| viper.GetUint16(string)          | config.Get(...string).Uint16(int)                        | uint16                 |
| viper.GetUint32(string)          | config.Get(...string).Uint32(int)                        | uint32                 |
| viper.GetUint64(string)          | config.Get(...string).Uint64(int)                        | uint64                 |
| viper.GetString(string)          | config.Get(...string).String(string)                     | string                 |
| viper.GetStringMap(string)       | config.Get(...string).StringMap(map[string]interface{})  | map[string]interface{} |
| viper.GetStringMapString(string) | config.Get(...string).StringMapString(map[string]string) | map[string]string      |
| viper.GetStringSlice(string)     | config.Get(...string).StringSlice([]string)              | []string               |
| viper.GetTime(string)            | config.Get(...string).Time(time.Time)                    | time.Time              |
| viper.GetDuration(string)        | config.Get(...string).Duration(time.Duration)            | time.Duration          |

config.Get takes a variadic string parameter that lays out the path of the configuration you need to retrieve. 
The following type method takes a single parameter that is the default value, which will be returned if the 
configuration is not available in the configuration file.

### Examples

The examples folder contains some examples of how to use the bootstrap and implementing a simple state store for your application.


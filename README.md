# go-bsn-cloud-client

Go client library for accessing the BrightSign BSN.Cloud API.

## Overview

`go-bsn-cloud-client` provides idiomatic Go data models and interfaces for interacting with the BrightSign BSN.Cloud API. It is designed for modularity, maintainability, and easy integration into other Go applications.

## Features

- Comprehensive data models for players, beacons, network interfaces, and more
- Support for custom types, enums, and sum types as required by the BSN.Cloud API
- Modular, well-documented codebase
- Ready for extension as the API evolves

## Installation

```
go get github.com/carrier-labs/go-bsn-cloud-client
```

## Usage Example

```go
package main

import (
    "context"
    "fmt"
    "github.com/carrier-labs/go-bsn-cloud-client/client"
    "github.com/carrier-labs/go-bsn-cloud-client/service"
)

func main() {
    // Create a new API client
    c := client.NewClient("<your-api-base-url>", "<your-username>", "<your-password>")
    deviceService := service.NewDeviceService(c)

    // Fetch players from a network
    players, err := deviceService.GetDevices(context.Background(), "<network-name>")
    if err != nil {
        panic(err)
    }
    for _, p := range players {
        fmt.Printf("Player: %s (%s)\n", p.Serial, p.Model)
    }
}
```

## Documentation

See GoDoc comments in the source code for detailed type and field documentation.

## Contributing

Contributions are welcome! Please open issues or submit pull requests for improvements or bug fixes.

## License

This project is licensed under the MIT License.

# Go Config
Rough Golang port of https://github.com/lorenwest/node-config

### Usage
```go
// Define a struct that matches your config file
type Configuration struct {
  Database map[string]string
}

var c Configuration

// Read the configuration
// If no GO_ENV is set it will default to config/default.json
// If that is not fail the package will return with an error
config, err := config.Open(&c)

// Read your values
log.Printf("%s", c.Database["dialect"])
```

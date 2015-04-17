# Go Config
Rough Golang port of https://github.com/lorenwest/node-config that is still pretty much a work in progress.

[![GoDoc](https://godoc.org/github.com/mattkrea/go-config?status.svg)](https://godoc.org/github.com/mattkrea/go-config) [![Build Status](https://travis-ci.org/mattkrea/go-config.svg)](https://travis-ci.org/mattkrea/go-config)

### Usage

The idea is to set the GO_ENV environment variable before launching your app allowing you to dynamically switch between configurations. For example, if you want to use a providers local database during a build you and use `GO_ENV=development` and add a file to your project root inside of a config folder named `development.json`.

When you are ready to move into production you can then run your app with `GO_ENV=production` which will load `production.json`.

### Example config/default.json
```json
{
  "server": {
    "port": ":8080"
  },
  "database": {
    "dialect": "mysql",
    "connection": "root@/test?parseTime=True&charset=utf8"
  }
}
```

### In your app
```go
// Define a struct that matches your config file
type Configuration struct {
  Server map[string]string `json:"server"`
  Database map[string]string `json:"database"`
}

var c Configuration

// Read the configuration
// If no GO_ENV is set it will default to config/default.json
// If that is not fail the package will return with an error
err := config.Open(&c)

// Read your values
log.Printf("%s", c.Database["dialect"])
// Prints 'mysql'
```

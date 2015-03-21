# Go Config
Rough Golang port of https://github.com/lorenwest/node-config

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

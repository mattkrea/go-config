package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const (
	// DEFAULT configuration points to config/default.json
	DEFAULT = "default"
)

// Open accepts an interface though this interface
// must accept or map to a JSON object. It is recommended
// that you provide JSON attributes for each property
func Open(c *interface{}) error {

	var (
		env      string
		filename string
	)

	// See if GO_ENV is populated
	env = os.Getenv("GO_ENV")

	if env != "" {
		filename = fmt.Sprintf("config/%s.json", env)
	}

	// Check to see that a configuration for this env actually exists
	_, err := os.Stat(filename)
	if err != nil {
		log.Printf("no configuration found for '%s'; using default instead", env)
		filename = fmt.Sprintf("config/%s.json", DEFAULT)

		// Now we have to check to see that the default configuration
		// exists. If it does not we have to fail out
		_, err = os.Stat(filename)
		if err != nil {
			return errors.New("no valid config file is available")
		}
	}

	// At this point we must have a config file that is valid
	// Load it and unmarshal into the provided interface
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		return errors.New("unable to read configuration file")
	}

	err = json.Unmarshal(raw, c)
	if err != nil {
		return errors.New("configuration file contains invalid JSON")
	}

	return nil
}

package goenv

import "os"
import "io/ioutil"

// Load - Loads and parses the contents of `.env`.
// Sets the appropriate environment variables
func Load() error {
	// read .env
	contents, fileErr := ioutil.ReadFile("./.env")
	if fileErr != nil {
		return fileErr
	}

	// parse .env conents
	envs, parseErr := parse(string(contents))
	if parseErr != nil {
		return parseErr
	}

	// set parsed environment variables
	for k, v := range envs {
		os.Setenv(k, v)
	}

	return nil
}

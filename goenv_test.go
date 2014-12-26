package goenv

import "os"
import "testing"
import "io/ioutil"
import "github.com/stretchr/testify/assert"

func TestLoad(t *testing.T) {
	copyEnv("./.env.success")

	err := Load()

	assert.Nil(t, err)

	assert.Equal(t, os.Getenv("TESTING"), "yes")
	assert.Equal(t, os.Getenv("STRING"), "testing")
	assert.Equal(t, os.Getenv("INTEGER"), "1234")
	assert.Equal(t, os.Getenv("FLOAT"), "12.34")

	assert.Equal(t, os.Getenv("FOO"), "")
	assert.Equal(t, os.Getenv("BAR"), "")

	assert.Equal(t, os.Getenv("TEST_HOME"), os.Getenv("HOME"))
	assert.Equal(t, os.Getenv("TEST_HOME_AGAIN"), os.Getenv("TEST_HOME"))
	assert.Equal(t, os.Getenv("INTERPOLATED"), "test:"+os.Getenv("HOME"))
	assert.Equal(t, os.Getenv("NOT_INTERPOLATED"), "test:$HOME")
	assert.Equal(t, os.Getenv("MULTI_INTERPOLATION"), "12")

	removeEnv(true)
}

func TestLoadFileNotFound(t *testing.T) {
	removeEnv(false)

	err := Load()
	assert.Equal(t, "Could not open file `./.env`", err.Error())
}

func TestLoadFormatError(t *testing.T) {
	copyEnv("./.env.fail")

	err := Load()
	assert.Equal(t, "FormatError: Line `lol$wut` doesn't match format", err.Error())

	removeEnv(true)
}

// helpers

func copyEnv(fileName string) {
	removeEnv(false)

	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("./.env", b, 0644)
	if err != nil {
		panic(err)
	}
}

func removeEnv(report bool) {
	err := os.Remove("./.env")
	if report && err != nil {
		panic(err)
	}
}

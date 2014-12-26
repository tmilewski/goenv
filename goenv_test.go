package goenv

import "os"
import "testing"

func TestLoad(t *testing.T) {
	copyEnv("./.env.success")

	err := Load()

	assertNil(t, err)

	assertEqual(t, os.Getenv("TESTING"), "yes")
	assertEqual(t, os.Getenv("STRING"), "testing")
	assertEqual(t, os.Getenv("INTEGER"), "1234")
	assertEqual(t, os.Getenv("FLOAT"), "12.34")

	assertEqual(t, os.Getenv("FOO"), "")
	assertEqual(t, os.Getenv("BAR"), "")

	assertEqual(t, os.Getenv("TEST_HOME"), os.Getenv("HOME"))
	assertEqual(t, os.Getenv("TEST_HOME_AGAIN"), os.Getenv("TEST_HOME"))
	assertEqual(t, os.Getenv("INTERPOLATED"), "test:"+os.Getenv("HOME"))
	assertEqual(t, os.Getenv("NOT_INTERPOLATED"), "test:$HOME")
	assertEqual(t, os.Getenv("MULTI_INTERPOLATION"), "12")

	removeEnv(true)
}

func TestLoadFileNotFound(t *testing.T) {
	removeEnv(false)

	err := Load()
	assertEqual(t, "Could not open file `./.env`", err.Error())
}

func TestLoadFormatError(t *testing.T) {
	copyEnv("./.env.fail")

	err := Load()
	assertEqual(t, "FormatError: Line `lol$wut` doesn't match format", err.Error())

	removeEnv(true)
}

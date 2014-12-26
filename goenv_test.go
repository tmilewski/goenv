package goenv

import "os"
import "fmt"
import "testing"
import "io/ioutil"
import "reflect"

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

//////////////////
// TEST HELPERS //
//////////////////

func assertEqual(t *testing.T, expected, actual interface{}) bool {
	if !objectsAreEqual(expected, actual) {
		t.Logf("Not equal: %#v (expected)\n"+
			"        != %#v (actual)", expected, actual)

		t.Fail()
		return false
	}

	return true
}

func assertNil(t *testing.T, actual interface{}) bool {
	if actual == nil {
		return true
	}

	t.Logf("Expected nil, but got: %#v", actual)
	t.Fail()
	return false
}

func objectsAreEqual(expected, actual interface{}) bool {
	if expected == nil || actual == nil {
		return expected == actual
	}

	if reflect.DeepEqual(expected, actual) {
		return true
	}

	actualType := reflect.TypeOf(actual)
	if actualType.ConvertibleTo(reflect.TypeOf(expected)) {
		expectedValue := reflect.ValueOf(expected)
		// Attempt comparison after type conversion
		if reflect.DeepEqual(actual, expectedValue.Convert(actualType).Interface()) {
			return true
		}
	}

	// Last ditch effort
	if fmt.Sprintf("%#v", expected) == fmt.Sprintf("%#v", actual) {
		return true
	}

	return false
}

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

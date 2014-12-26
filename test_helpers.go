package goenv

import "os"
import "fmt"
import "testing"
import "io/ioutil"
import "reflect"

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

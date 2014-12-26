package goenv

import "os"
import "testing"

func TestSubstituteVariables(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping substituting variables")
	}

	var actual string

	// exists in local envs
	actual = substituteVariables(`\$ESCAPED`, map[string]string{})
	assertEqual(t, `$ESCAPED`, actual)

	// exists in local envs
	actual = substituteVariables(`$FOO`, map[string]string{"FOO": "bar"})
	assertEqual(t, "bar", actual)
	os.Setenv("FOO", "")

	// exists in global envs
	os.Setenv("FOO", "bar2")
	actual = substituteVariables(`$FOO`, map[string]string{})
	assertEqual(t, "bar2", actual)
	os.Setenv("FOO", "")

	// does not exist in global or local envs
	actual = substituteVariables(`$DOESNTEXIST`, map[string]string{})
	assertEqual(t, "", actual)

	// expands variables in double quoted strings
	actual = substituteVariables(`quote $FOO`, map[string]string{"FOO": "bar"})
	assertEqual(t, "quote bar", actual)
}

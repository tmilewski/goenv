package goenv

import "os"
import "testing"

func env(t *testing.T, contents string, expected map[string]string) {
	actual, _ := parse(contents)
	assertEqual(t, expected, actual)
}

func TestParser(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping parser")
	}

	var expected map[string]string

	// parses unquoted values
	expected = map[string]string{"FOO": "bar"}
	env(t, "FOO=bar", expected)

	// parses values with spaces around equal sign
	env(t, "FOO =bar", expected)
	env(t, "FOO= bar", expected)

	// parses double quoted values
	env(t, `FOO="bar"`, expected)

	// parses single quoted values
	env(t, "FOO='bar'", expected)

	// parses yaml style options
	env(t, "FOO: bar", expected)

	// parses export keyword
	env(t, "export FOO=bar", expected)

	// strips unquoted values
	env(t, "FOO=bar ", expected) // not 'bar '

	// ignores inline comments
	env(t, "FOO=bar # this is foo", expected)

	// ignores comment lines
	env(t, "\n\n\n # HERE GOES FOO \nFOO=bar", expected)

	// reads variables from ENV when expanding if not found in local env
	os.Setenv("TEST", "bar")
	env(t, "FOO=$TEST", expected)
	os.Setenv("TEST", "")

	// multiple assignments
	expected = map[string]string{"FOO": "bar", "TEST": "foo"}
	env(t, "FOO=bar\nTEST=foo", expected)

	// parses escaped double quotes
	expected = map[string]string{"FOO": `escaped"bar`}
	env(t, `FOO="escaped\"bar"`, expected)

	// parses empty values
	expected = map[string]string{"FOO": ""}
	env(t, "FOO=", expected)

	// expands undefined variables to an empty string
	expected = map[string]string{"BAR": ""}
	env(t, "BAR=$NONE", expected)

	// expands variables in double quoted strings
	os.Setenv("FOO", "test")
	expected = map[string]string{"FOO": "test", "BAR": "quote test"}
	env(t, "FOO=test\nBAR=\"quote $FOO\"", expected)
	os.Setenv("FOO", "")

	// does not expand variables in single quoted strings
	expected = map[string]string{"BAR": "quote $FOO"}
	env(t, "BAR='quote $FOO'", expected)

	// does not expand escaped variables
	expected = map[string]string{"FOO": "foo$BAR"}
	env(t, `FOO="foo\$BAR"`, expected)

	expected = map[string]string{"FOO": "foo${BAR}"}
	env(t, `FOO="foo\${BAR}"`, expected)

	// parses variables with "." in the name
	expected = map[string]string{"FOO.BAR": "foobar"}
	env(t, `FOO.BAR=foobar`, expected)

	// allows # in quoted value
	expected = map[string]string{"foo": "bar#baz"}
	env(t, `foo="bar#baz" # comment`, expected)

	// parses # in quoted values
	expected = map[string]string{"foo": "ba#r"}
	env(t, `foo="ba#r"`, expected)
	env(t, `foo='ba#r'`, expected)

	// expands newlines in quoted strings
	expected = map[string]string{"FOO": "bar\nbaz"}
	env(t, `FOO="bar\nbaz"`, expected)

	// ignores empty lines
	expected = map[string]string{"foo": "bar", "fizz": "buzz"}
	env(t, "\n \t  \nfoo=bar\n \nfizz=buzz", expected)
}

func TestFormatError(t *testing.T) {
	actual, err := parse("lol$wut")

	// Format Error
	assertEqual(t, 0, len(actual))
	assertEqual(t, "FormatError: Line `lol$wut` doesn't match format", err.Error())

	// Comment
	_, err = parse("# BAR")
	assertNil(t, err)

	// Blank Line
	_, err = parse("   \n# BAR")
	assertNil(t, err)
}

package goenv

import "os"
import "regexp"

const variableRegex = `(\\)?(\$)(\{?([A-Z0-9_]+)\}?)`

func substituteVariables(value string, envs map[string]string) string {
	match, _ := regexp.MatchString(variableRegex, value)

	if match {
		r := regexp.MustCompile(variableRegex)

		value = r.ReplaceAllStringFunc(value, func(m string) string {
			parts := r.FindStringSubmatch(m)

			if parts[1] == `\` {
				return parts[2] + parts[3]
			}

			e := envs[parts[3]]
			if e != "" {
				return e
			}

			return os.Getenv(parts[3])
		})
	}

	return value
}

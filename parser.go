package goenv

import "fmt"
import "strings"
import "regexp"

const regex = `\A(?:export\s+)?([\w\.]+)(?:\s*=\s*|:\s+?)('(?:\'|[^'])*'|"(?:\"|[^"])*"|[^#\n]+)?(?:\s*\#.*)?\z`
const formatRegex = `\A\s*(?:#.*)?\z`
const quotesRegex = `\A(['"])?(.*)(['"])\z`

func parse(input string) (map[string]string, error) {
	lines := strings.Split(input, "\n")

	envs := make(map[string]string)
	re, _ := regexp.Compile(regex)

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		match, _ := regexp.MatchString(regex, line)
		matchFormat, _ := regexp.MatchString(formatRegex, line)

		if match {
			res := re.FindAllStringSubmatch(line, -1)

			if res != nil {
				key, value := res[0][1], res[0][2]

				// Remove surrounding quotes
				quotes := regexp.MustCompile(quotesRegex)
				q := quotes.FindAllStringSubmatch(value, -1)

				quoteType := ""

				if len(q) > 0 {
					quoteType = q[0][1]
					value = q[0][2]
				}

				if quoteType == `"` {
					r := regexp.MustCompile(`\\n`)
					value = r.ReplaceAllString(value, "\n")

					r = regexp.MustCompile(`\\([^$])`)
					value = r.ReplaceAllString(value, "${1}")
				}

				if quoteType != `'` {
					value = substituteVariables(value, envs)
					// TODO: Substitute Commands
				}

				envs[key] = strings.TrimSpace(value)
			}
		} else if !matchFormat {
			return map[string]string{}, fmt.Errorf("FormatError: Line `%s` doesn't match format", line)
		}
	}

	return envs, nil
}

# GoEnv: Project-Specific Environment Variables in Go

[![Build Status](https://travis-ci.org/tmilewski/goenv.svg?branch=master)](https://travis-ci.org/tmilewski/goenv)
[![Coverage Status](https://img.shields.io/coveralls/tmilewski/goenv.svg)](https://coveralls.io/r/tmilewski/goenv?branch=master)

## An Introduction

> "The twelve-factor app stores config in environment variables (often shortened to env vars or env). Env vars are easy to change between deploys without changing any code; unlike config files, there is little chance of them being checked into the code repo accidentally; and unlike custom config files, or other config mechanisms such as Java System Properties, they are a language- and OS-agnostic standard."
> 
> \- [The Twelve Factor App](http://12factor.net/) / [Config](http://12factor.net/config)


## Getting Started

### Import
`import "github.com/tmilewski/goenv"`

### Usage
`goenv.Load()`

### Examples

Examples may be found in the [/examples](https://github.com/tmilewski/goenv/tree/master/examples) directory.

*.env*
```bash
HELLO="WORLD"
SENSITIVE_KEY="sensitive value"
YOUR_HOME=$HOME
```

_*main.go*_
```go
import "fmt"
import "os"
import "github.com/tmilewski/goenv"

func init() {
	err := goenv.Load()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	fmt.Println("HELLO", os.Getenv("HELLO")) // Prints "HELLO WORLD"
	fmt.Println(os.Getenv("SENSITIVE_KEY")) // Prints "sensitive value"
	fmt.Println(os.Getenv("HOME_DIR")) // Prints <your home directory>
}
```

## Supported Go Versions

GoEnv is tested under 1.2, 1.3, 1.4.

## Testing

`go test -v`

## What's Next

1. Substituting commands: `CWD=$(pwd)`
2. Make regular expressions more readable.

## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request

## License

GoEnv is released under the [MIT License](http://www.opensource.org/licenses/MIT).

## Thanks
Heavily influenced by the amazing work of [Brandon Keepers](https://github.com/bkeepers) on Ruby's [dotenv](https://github.com/bkeepers/dotenv).

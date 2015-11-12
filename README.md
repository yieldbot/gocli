## gocli

[![Build Status][travis-image]][travis-url] [![Coverage][coverage-image]][coverage-url] [![GoDoc][godoc-image]][godoc-url] [![Release][release-image]][release-url]

A Go CLI library that provides subcommand handling, tidy usage and version printing.

### Installation

```
go get github.com/yieldbot/gocli
```

### Usage

#### A simple CLI app

See [simple.go](examples/simple.go) for full code.

```go
func init() {
  // Init flags
  flag.BoolVar(&usageFlag, "h", false, "Display usage")
  flag.BoolVar(&usageFlag, "help", false, "Display usage")
  flag.BoolVar(&versionFlag, "version", false, "Display version information")
  flag.BoolVar(&versionFlag, "v", false, "Display version information")
}

func main() {

  // Init cli
  cli = gocli.Cli{
    AppName:    "simple",
    AppVersion: "1.0.0",
    AppDesc:    "A simple app",
    CommandList: map[string]string{
      "echo": "Print the given arguments",
    },
  }
  cli.Init()

  // Run commands
  if cli.Command == "echo" {
    // Echo command
    fmt.Println(strings.Join(cli.CommandArgs, " "))
  } else if versionFlag {
    // Version
    cli.PrintVersion(true)
  } else {
    // Default
    cli.PrintUsage()
  }
}
```

`$ go run examples/simple.go`
```bash
Usage: simple [OPTIONS] COMMAND [arg...]

A simple app

Options:
  -h, --help    : Display usage
  -v, --version : Display version information

Commands:
  echo          : Print the given arguments
```

`$ go run examples/simple.go -v`
```bash
App version : 1.0.0
Go version  : go1.5.1
```

`$ go run examples/simple.go echo hello world`
```bash
hello world
```

### License

Licensed under The MIT License (MIT)  
For the full copyright and license information, please view the LICENSE.txt file.

[travis-url]: https://travis-ci.org/yieldbot/gocli
[travis-image]: https://travis-ci.org/yieldbot/gocli.svg?branch=master

[godoc-url]: https://godoc.org/github.com/yieldbot/gocli
[godoc-image]: https://godoc.org/github.com/yieldbot/gocli?status.svg

[release-url]: https://github.com/yieldbot/gocli/releases/tag/v1.0.2
[release-image]: https://img.shields.io/badge/release-v1.0.2-blue.svg

[coverage-url]: https://coveralls.io/github/yieldbot/gocli?branch=master
[coverage-image]: https://coveralls.io/repos/yieldbot/gocli/badge.svg?branch=master&service=github)
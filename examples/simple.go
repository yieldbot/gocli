// A simple app
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/yieldbot/gocli"
)

var (
	cli         gocli.Cli
	usageFlag   bool
	versionFlag bool
	test_test   bool
)

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

	// Echo command
	if cli.Command == "echo" {
		fmt.Println(strings.Join(cli.CommandArgs, " "))
		os.Exit(0)
	}

	// Version
	if versionFlag {
		cli.PrintVersion(true)
	}

	// Default
	cli.PrintUsage()
}

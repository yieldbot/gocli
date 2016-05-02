// A simple app
package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/yieldbot/gocli"
)

var (
	cli         gocli.Cli
	usageFlag   bool
	versionFlag bool
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
		Name:        "simple",
		Version:     "1.0.0",
		Description: "A simple app",
		Commands: map[string]string{
			"echo": "Print the given arguments",
		},
	}
	cli.Init()

	// Run commands
	if cli.SubCommand == "echo" {
		// Echo command
		fmt.Println(strings.Join(cli.SubCommandArgs, " "))
	} else if versionFlag {
		// Version
		cli.PrintVersion(true)
	} else {
		// Default
		cli.PrintUsage()
	}
}

/*
 * gocli
 * Copyright (c) 2015 Yieldbot, Inc. (http://github.com/yieldbot/gocli)
 * For the full copyright and license information, please view the LICENSE.txt file.
 */

package gocli_test

import (
	"flag"
	"os"
	"testing"

	"github.com/yieldbot/gocli"
)

var (
	argFlag     string
	usageFlag   bool
	versionFlag bool
)

func init() {
	// Init flags
	flag.StringVar(&argFlag, "arg", "", "Arg flag")
	flag.BoolVar(&usageFlag, "h", false, "Display usage")
	flag.BoolVar(&usageFlag, "help", false, "Display usage")
	flag.BoolVar(&versionFlag, "version", false, "Display version information")
	flag.BoolVar(&versionFlag, "v", false, "Display version information")
}

func TestInit(t *testing.T) {

	// Init cli
	var cli = gocli.Cli{
		AppName:    "test",
		AppVersion: "1.0.0",
		AppDesc:    "test desc",
		CommandList: map[string]string{
			"cmd": "Test command",
		},
	}
	cli.Init()

	if cli.AppName != "test" {
		t.Error("invalid AppName")
	}

	if cli.AppVersion != "1.0.0" {
		t.Error("invalid AppVersion")
	}

	if cli.AppDesc != "test desc" {
		t.Error("invalid AppDesc")
	}

	if cli.Command != "" {
		t.Error("invalid Command")
	}

	if len(cli.CommandArgs) > 0 {
		t.Error("invalid CommandArgs")
	}

	if len(cli.CommandList) != 1 {
		t.Error("invalid CommandList")
	}

	if cli.CommandList["cmd"] != "Test command" {
		t.Error("invalid CommandList")
	}

	os.Args = append(os.Args, "cmd", "arg1")

	// Init cli
	var cli2 = gocli.Cli{
		AppName:    "test",
		AppVersion: "1.0.0",
		AppDesc:    "test desc",
		CommandList: map[string]string{
			"cmd": "Test command",
		},
	}
	cli2.Init()

	if cli2.Command != "cmd" {
		t.Error("invalid Command")
	}

	if len(cli2.CommandArgs) != 1 {
		t.Error("invalid CommandArgs")
	}

	if cli2.CommandArgs[0] != "arg1" {
		t.Error("invalid CommandArgs arg")
	}
}

func ExampleF_PrintVersion() {
	var cli = gocli.Cli{
		AppVersion: "1.0.0",
	}
	cli.Init()

	cli.PrintVersion(false)
	// Output: 1.0.0
}

func ExampleF_PrintVersionExtra() {
	var cli = gocli.Cli{
		AppVersion: "1.0.0",
	}
	cli.Init()

	cli.PrintVersion(true)
}

func ExampleF_PrintUsage() {

	// Init cli
	var cli = gocli.Cli{
		AppName:    "test",
		AppVersion: "1.0.0",
		AppDesc:    "test desc",
		CommandList: map[string]string{
			"cmd": "Test command",
		},
	}
	cli.Init()

	cli.PrintUsage()
	// Output:
	// Usage: test [OPTIONS] COMMAND [arg...]
	//
	// test desc
	//
	// Options:
	//   --arg         : Arg flag
	//   -h, --help    : Display usage
	//   -v, --version : Display version information
	//
	// Commands:
	//   cmd           : Test command
}

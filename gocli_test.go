/*
 * gocli
 * Copyright (c) 2015 Yieldbot, Inc.
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
	flag.StringVar(&argFlag, "arg", "test", "Arg flag")
	flag.BoolVar(&usageFlag, "h", false, "Display usage")
	flag.BoolVar(&usageFlag, "help", false, "Display usage")
	flag.BoolVar(&versionFlag, "version", false, "Display version information")
	flag.BoolVar(&versionFlag, "v", false, "Display version information")
}

func TestInit_1(t *testing.T) {

	// Reset the args
	os.Args = os.Args[:2]

	// Init cli
	var cli = gocli.Cli{
		Name:        "test",
		Version:     "1.0.0",
		Description: "test desc",
		Commands: map[string]string{
			"cmd": "Test command",
		},
	}
	cli.Init()

	if cli.Name != "test" {
		t.Error("invalid Name")
	}

	if cli.Version != "1.0.0" {
		t.Error("invalid Version")
	}

	if cli.Description != "test desc" {
		t.Error("invalid Description")
	}

	if len(cli.Commands) != 1 {
		t.Error("invalid Commands")
	}

	if cli.Commands["cmd"] != "Test command" {
		t.Error("invalid Commands")
	}

	if cli.SubCommand != "" {
		t.Error("invalid SubCommand")
	}

	if len(cli.SubCommandArgs) > 0 {
		t.Error("invalid SubCommandArgs")
	}
}

func TestInit_2(t *testing.T) {

	// Reset the args
	os.Args = os.Args[:2]
	os.Args = append(os.Args, "cmd", "arg1")

	// Init cli
	var cli = gocli.Cli{
		Name:        "test",
		Version:     "1.0.0",
		Description: "test desc",
		Commands: map[string]string{
			"cmd": "Test command",
		},
	}
	cli.Init()

	if cli.SubCommand != "cmd" {
		t.Error("invalid SubCommand")
	}

	if len(cli.SubCommandArgs) != 1 {
		t.Error("invalid SubCommandArgs")
	}

	if cli.SubCommandArgs[0] != "arg1" {
		t.Error("invalid SubCommandArgs arg")
	}
}

func TestInit_3(t *testing.T) {

	// Reset the args
	os.Args = os.Args[:2]
	os.Args = append(os.Args, "cmd", "--arg2", "arg3", "arg4")

	// Init cli
	var cli = gocli.Cli{
		Name:        "test",
		Version:     "1.0.0",
		Description: "test desc",
		Commands: map[string]string{
			"cmd": "Test command",
		},
	}
	cli.Init()

	if cli.SubCommand != "cmd" {
		t.Error("invalid SubCommand")
	}

	if len(cli.SubCommandArgs) != 3 {
		t.Error("invalid SubCommandArgs")
	}

	if cli.SubCommandArgs[0] != "--arg2" {
		t.Error("invalid SubCommandArgs arg")
	}

	if cli.SubCommandArgsMap["arg2"] != "arg3" {
		t.Error("invalid SubCommandArgsMap arg")
	}

	if _, ok := cli.SubCommandArgsMap["arg4"]; !ok {
		t.Error("invalid SubCommandArgsMap arg")
	}
}

func ExampleF_PrintVersion() {
	var cli = gocli.Cli{
		Version: "1.0.0",
	}
	cli.Init()

	cli.PrintVersion(false)
	// Output: 1.0.0
}

func ExampleF_PrintVersionExtra() {
	var cli = gocli.Cli{
		Version: "1.0.0",
	}
	cli.Init()

	cli.PrintVersion(true)
}

func ExampleF_PrintUsage() {

	// Init cli
	var cli = gocli.Cli{
		Name:        "test",
		Version:     "1.0.0",
		Description: "test desc",
		Commands: map[string]string{
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
	//   --arg         : Arg flag (default "test")
	//   -h, --help    : Display usage
	//   -v, --version : Display version information
	//
	// Commands:
	//   cmd           : Test command
}

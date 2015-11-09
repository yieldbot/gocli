/*
 * gocli
 * Copyright (c) 2015 Yieldbot, Inc. (http://github.com/yieldbot/gocli)
 * For the full copyright and license information, please view the LICENSE.txt file.
 */

package gocli_test

import (
	"flag"
	"testing"

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

	// Init cli
	cli = gocli.Cli{
		AppName:    "test",
		AppVersion: "1.0.0",
		AppDesc:    "test desc",
		CommandList: map[string]string{
			"cmd": "Test command",
		},
	}
	cli.Init()
}

func TestInit(t *testing.T) {

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
}

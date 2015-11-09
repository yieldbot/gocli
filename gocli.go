/*
 * gocli
 * Copyright (c) 2015 Yieldbot, Inc. (http://github.com/yieldbot/gocli)
 * For the full copyright and license information, please view the LICENSE.txt file.
 */

// Package gocli is a CLI library that provides subcommand handling, tidy usage and version printing.
package gocli

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"sort"
	"strings"
)

// Cli represent command line interface
type Cli struct {
	// AppName is the app name
	AppName string

	// AppVersion is the app version
	AppVersion string

	// AppDesc is the app description
	AppDesc string

	// Command is the subcommand for the app
	Command string

	// CommandArgs contains the args of the subcommand
	CommandArgs []string

	// CommandList contains the subcommand list of the app
	CommandList map[string]string

	// LogOut is logger for stdout
	LogOut *log.Logger

	// LogErr is logger for stderr
	LogErr *log.Logger

	// Flags contains flags
	//Flags map[string]string
}

// Init initializes Cli instance
func (cl *Cli) Init() {

	// Init flag
	if !flag.Parsed() {
		flag.Parse()
	}

	// Init loggers
	cl.LogOut = log.New(os.Stdout, "", log.LstdFlags)
	cl.LogErr = log.New(os.Stderr, "", log.LstdFlags)

	// Init flags
	// cl.Flags = make(map[string]string)
	// flag.VisitAll(func(f *flag.Flag) {
	//  cl.Flags[f.Name] = f.Value.String()
	// })

	// Init args
	// If length of the args more than one
	if len(os.Args) > 1 {

		// Iterate the args
		for _, arg := range os.Args {

			// If the arg is in command list then
			if _, ok := cl.CommandList[arg]; ok {
				cl.Command = arg // set as command
			} else {
				// Otherwise add it to args
				if cl.Command != "" {
					cl.CommandArgs = append(cl.CommandArgs, arg)
				}
			}
		}
	}
}

// PrintVersion prints version information
func (cl Cli) PrintVersion(extra bool) {
	var ver string

	if extra == true {
		ver += fmt.Sprintf("App version : %s\n", cl.AppVersion)
		ver += fmt.Sprintf("Go version  : %s", runtime.Version())
	} else {
		ver = fmt.Sprintf("%s", cl.AppVersion)
	}

	fmt.Println(ver)
}

// PrintUsage prints usage info
// Usage format follows common convention for Go apps
func (cl Cli) PrintUsage() {

	// Init vars
	type flagInfo struct {
		nameu    string
		name     string
		usage    string
		defValue string
	}

	// Find max length by command list
	maxlen := 0
	if len(cl.CommandList) > 0 {
		for c := range cl.CommandList {
			if len(c) > maxlen {
				maxlen = len(c)
			}
		}
	}

	// Iterate flags
	flagList := make(map[string]*flagInfo)

	flag.VisitAll(func(f *flag.Flag) {

		// If flag name doesn't start with `test.` then
		if strings.Index(f.Name, "test.") != 0 {

			// Set key by the flag usage for grouping
			key := fmt.Sprint(f.Usage)

			// Init usage name
			nameu := "-" + f.Name
			if len(f.Name) > 2 {
				nameu = "-" + nameu
			}

			// If the flag exists then
			if _, ok := flagList[key]; ok {
				// Merge names
				flagList[key].nameu += ", " + nameu
			} else {
				// Otherwise add the flag
				flagList[key] = &flagInfo{
					nameu:    nameu,
					name:     f.Name,
					usage:    f.Usage,
					defValue: f.DefValue,
				}
			}

			// Check and set maximum length for alignment
			if len(flagList[key].nameu) > maxlen {
				maxlen = len(flagList[key].nameu)
			}
		}
	})

	// Fixed flag list
	flagListF := []string{}
	for _, v := range flagList {
		flagline := fmt.Sprintf("%s : %s", strPadRight(v.nameu, " ", maxlen), v.usage)
		if v.defValue != "false" {
			flagline += " (default \"" + v.defValue + "\")"
		}
		flagListF = append(flagListF, flagline)
	}
	sort.Strings(flagListF)

	// Fixed command list
	cmdListF := []string{}
	for cn, cv := range cl.CommandList {
		cmdListF = append(cmdListF, fmt.Sprintf("%s : %s", strPadRight(cn, " ", maxlen), cv))
	}
	sort.Strings(cmdListF)

	// Header and description
	usage := "Usage: " + cl.AppName + " [OPTIONS] COMMAND [arg...]\n\n"
	if cl.AppDesc != "" {
		usage += cl.AppDesc + "\n\n"
	}

	// Options
	if len(flagListF) > 0 {
		usage += "Options:\n"
		for _, f := range flagListF {
			usage += fmt.Sprintf("  %s\n", f)
		}
	}

	// Commands
	if len(cmdListF) > 0 {
		usage += "\nCommands:\n"
		for _, c := range cmdListF {
			usage += fmt.Sprintf("  %s\n", c)
		}
	}

	fmt.Println(usage)
}

// strPadRight provides padding for strings
func strPadRight(str, pad string, length int) string {
	for {
		str += pad
		if len(str) > length {
			return str[0:length]
		}
	}
}

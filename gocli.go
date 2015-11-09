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

	flagList := make(map[string]*flagInfo)
	maxlen := 0

	// Find max length by command list
	if len(cl.CommandList) > 0 {
		for c := range cl.CommandList {
			if len(c) > maxlen {
				maxlen = len(c)
			}
		}
	}

	// Iterate flags
	flag.VisitAll(func(f *flag.Flag) {

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
	})

	// Add usage header and description
	usage := "Usage: " + cl.AppName + " [OPTIONS] COMMAND [arg...]\n\n"
	if cl.AppDesc != "" {
		usage += cl.AppDesc + "\n\n"
	}

	// Add options
	if len(flagList) > 0 {
		usage += "Options:\n"
		for _, f := range flagList {
			usage += "  " + strPadRight(f.nameu, " ", maxlen) + " : " + f.usage
			if f.defValue != "false" {
				usage += " (default \"" + f.defValue + "\")"
			}
			usage += "\n"
		}
	}

	// Add commands
	if len(cl.CommandList) > 0 {
		usage += "\nCommands:\n"
		for cn, cv := range cl.CommandList {
			usage += "  " + strPadRight(cn, " ", maxlen) + " : " + cv + "\n"
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

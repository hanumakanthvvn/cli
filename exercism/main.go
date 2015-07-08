package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/codegangsta/cli"
	"github.com/hanumakanthvvn/cli/api"
	"github.com/hanumakanthvvn/cli/cmd"
)

const (
	// Version is the current release of the command-line app.
	// We try to follow Semantic Versioning (http://semver.org),
	// but with the http://exercism.io app being a prototype, a
	// lot of things get out of hand.
	Version = "2.0.0-rc.1"

	descDebug     = "Outputs useful debug information."
	descConfigure = "Writes config values to a JSON file."
	descDemo      = "Fetches a demo problem for each language track on exercism.io."
	descFetch     = "Fetches your current problems on hootcode.com, as well as the next unstarted problem in each language."
	descRestore   = "Restores completed and current problems on from hootcode.com, along with your most recent iteration for each."
	descSubmit    = "Submits a new iteration to a problem on exercism.io."
	descGsubmit   = "Submit from git."
	descUnsubmit  = "Deletes the most recently submitted iteration."
	descTracks    = "List the available language tracks"
	descOpen      = "Opens the current submission of the specified exercise"

	descLongRestore = "Restore will pull the latest revisions of exercises that have already been submitted. It will *not* overwrite existing files. If you have made changes to a file and have not submitted it, and you're trying to restore the last submitted version, first move that file out of the way, then call restore."
	descDownload    = "Downloads and saves a specified submission into the local system"
	descList        = "Lists all available assignments for a given language"
)

func main() {
	api.UserAgent = fmt.Sprintf("github.com/exercism/cli v%s (%s/%s)", Version, runtime.GOOS, runtime.GOARCH)

	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Name = "hootcode"
	app.Usage = "A command line tool to interact with http://hootcode.com"
	app.Version = Version
	app.HideVersion = true
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Usage: "path to config file",
		},
		cli.BoolFlag{
			Name:  "verbose, v",
			Usage: "turn on verbose logging",
		},
		cli.BoolFlag{
			Name:  "version",
			Usage: "print the version",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:   "debug",
			Usage:  descDebug,
			Action: cmd.Debug,
		},
		{
			Name:  "configure",
			Usage: descConfigure,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "dir, d",
					Usage: "path to exercises directory",
				},
				cli.StringFlag{
					Name:  "host, u",
					Usage: "hootcode api host",
				},
				cli.StringFlag{
					Name:  "key, k",
					Usage: "Hootcode API key (see http://hootcode.com/account)",
				},
				cli.StringFlag{
					Name:  "api, a",
					Usage: "exercism xapi host",
				},
			},
			Action: cmd.Configure,
		},
		{
			Name:      "demo",
			ShortName: "d",
			Usage:     descDemo,
			Action:    cmd.Demo,
		},
		{
			Name:      "fetch",
			ShortName: "f",
			Usage:     descFetch,
			Action:    cmd.Fetch,
			BashComplete: func(c *cli.Context) {
			  if len(c.Args()) == 0 {
	      		    fmt.Println("java node-js ruby python")
	                  }

	          	  if len(c.Args()) == 1{
		      	    fmt.Println("hamming leap bob gigasecond ")
	                  }
			},
		},
		{
			Name:        "restore",
			ShortName:   "r",
			Usage:       descRestore,
			Description: descLongRestore,
			Action:      cmd.Restore,
		},
		{
			Name:      "submit",
			ShortName: "s",
			Usage:     descSubmit,
			Action:    cmd.Submit,
		},
		{
			Name:      "gsubmit",
			ShortName: "gs",
			Usage:     descGsubmit,
			Action:    cmd.GSubmit,
		},
		{
			Name:      "unsubmit",
			ShortName: "u",
			Usage:     descUnsubmit,
			Action:    cmd.Unsubmit,
		},
		{
			Name:      "tracks",
			ShortName: "t",
			Usage:     descTracks,
			Action:    cmd.Tracks,
		},
		{
			Name:      "open",
			ShortName: "op",
			Usage:     descOpen,
			Action:    cmd.Open,
		},
		{
			Name:      "download",
			ShortName: "dl",
			Usage:     descDownload,
			Action:    cmd.Download,
		},
		{
			Name:      "list",
			ShortName: "li",
			Usage:     descList,
			Action:    cmd.List,
		},
	}
	if err := app.Run(os.Args); err != nil {

		log.Fatal(err)
	}
}

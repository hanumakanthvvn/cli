package cmd

import (
	"fmt"
	"log"

	"github.com/codegangsta/cli"
	"github.com/hanumakanthvvn/cli/api"
	"github.com/hanumakanthvvn/cli/config"
	"github.com/hanumakanthvvn/cli/user"
)

// Tracks lists available tracks.
func Tracks(ctx *cli.Context) {
	c, err := config.New(ctx.GlobalString("config"))
	if err != nil {
		log.Fatal(err)
	}
	client := api.NewClient(c)

	tracks, err := client.Tracks()
	if err != nil {
		log.Fatal(err)
	}

	curr := user.NewCurriculum(tracks)
	fmt.Println("\nActive language tracks:")
	curr.Report(user.TrackActive)
	fmt.Println("\nInactive language tracks:")
	curr.Report(user.TrackInactive)

	// TODO: implement `list` command to list problems in a track
	msg := `
Related commands:
    hootcode fetch (see 'hootcode help fetch')
	`
	fmt.Println(msg)
}

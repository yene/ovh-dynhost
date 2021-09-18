package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/urfave/cli"
)

// App build information
var (
	Version string
	Build   string
)

var config = &ConfigEntries{}

type ConfigEntries struct {
	Username         string
	Password         string
	IpAddress        string
	NetworkInterface string
}

func appDefinition() (app *cli.App) {
	app = cli.NewApp()
	app.Name = "ovh-dynhost"
	app.Authors = []cli.Author{
		{
			Name:  "Yannick Weiss",
			Email: "",
		},
	}
	app.Usage = "Update OVH DynHost DNS record."
	app.Version = Version + " (build " + Build + ")"
	app.Commands = []cli.Command{
		{
			Name:   "update-record",
			Usage:  "Update a DynHost record",
			Action: UpdateRecord,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "ip-address",
					Value:       "",
					Usage:       "The IP address that should be used to update the DynHost (bypass auto-detection).",
					Destination: &config.IpAddress,
				},
				cli.StringFlag{
					Name:        "interface, I",
					Value:       "",
					Usage:       "The interface whose IP address should be used to update the DynHost.",
					Destination: &config.NetworkInterface,
				},
			},
		},
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "username",
			Usage:       "The OVH DynHost username",
			Destination: &config.Username,
		},
		cli.StringFlag{
			Name:        "password",
			Usage:       "The OVH DynHost password",
			Destination: &config.Password,
		},
	}

	return
}

func main() {
	log.SetOutput(ioutil.Discard)
	appDefinition().Run(os.Args)
}

package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/investapp_server/in_app/cmd/app"
	"github.com/urfave/cli/v2"
)

var cmd = &cli.App{
	Name:        "Invest App",
	Usage:       "backend",
	Description: "Invest App backend",
	Authors: []*cli.Author{
		{Name: "Karel Kopřiva", Email: "karel.hb@email.cz"},
	},

	Commands: []*cli.Command{
		app.RunCMD,
	},
}

func main() {
	err := cmd.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
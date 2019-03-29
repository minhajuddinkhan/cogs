package main

import (
	"log"
	"os"

	"github.com/minhajuddinkhan/cogs/cmd"
	"github.com/minhajuddinkhan/cogs/store/bolt"
	"github.com/urfave/cli"
)

const (
	volume = "/.cogs"
	dbName = "data.db"
)

func main() {

	app := cli.NewApp()
	store := bolt.New(volume, dbName)
	app.Before = cmd.BeforeAction(store)
	app.Commands = []cli.Command{
		cmd.Lunch(store),
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

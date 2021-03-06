package main

import (
	"log"
	"os"

	"github.com/minhajuddinkhan/cogs/cmd"
	"github.com/minhajuddinkhan/cogs/services/cogs"
	"github.com/minhajuddinkhan/cogs/store/bolt"
	"github.com/minhajuddinkhan/cogs/types"
	"github.com/urfave/cli"
)

const (
	volume = "/.cogs"
	dbName = "data.db"
)

func main() {

	app := cli.NewApp()
	store := bolt.New(volume, dbName)
	var creds types.Credentials
	app.Before = cmd.BeforeAction(store, &creds)

	cogs := cogs.New(&creds, store)

	app.Commands = []cli.Command{
		cmd.Lunch(cogs),
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

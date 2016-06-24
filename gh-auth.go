package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "gh-auth"
	app.Usage = "allows you to quickly pair with anyone who has a GitHub account"
	app.Commands = []cli.Command{
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "List all users",
			Action: func(c *cli.Context) error {
				f, err := os.Open("~/.ssh/authorized_keys")
				if err != nil {
					log.Fatal(err)
				}
				defer f.Close()

				return nil
			},
		},
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "Add user to authorized_keys",
			Action: func(c *cli.Context) error {
				var user = c.Args().First()
				resp, err := http.Get("https://api.github.com/users/" + user + "/keys")
				if err != nil {
					log.Fatal(err)
				}
				key, err := ioutil.ReadAll(resp.Body)
				resp.Body.Close()
				fmt.Printf("%s", key)
				return nil
			},
		},
	}

	app.Run(os.Args)
}

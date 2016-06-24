package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/urfave/cli"
)

type SSHKey struct {
	Id  int    `json:"id"`
	Key string `json:"key"`
}

func main() {
	home := os.Getenv("HOME")
	url := "https://api.github.com"
	app := cli.NewApp()
	app.Name = "gh-auth"
	app.Version = "1.0.0"
	app.Usage = "allows you to quickly pair with anyone who has a GitHub account"
	app.Commands = []cli.Command{
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "List all users",
			Action: func(c *cli.Context) error {
				f, err := ioutil.ReadFile(home + "/.ssh/authorized_keys")
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(string(f))
				return nil
			},
		},
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "Add user to authorized_keys",
			Action: func(c *cli.Context) error {
				var user = c.Args().First()
				resp, err := http.Get(url + "/users/" + user + "/keys")
				if err != nil {
					log.Fatal(err)
				}
				dec := json.NewDecoder(resp.Body)
				for {
					s := make([]SSHKey, 0)
					if err := dec.Decode(&s); err == io.EOF {
						break
					} else if err != nil {
						log.Fatal(err)
					}
					fmt.Printf("%s: %s\n", s[0].Id, s[0].Key)
				}
				resp.Body.Close()
				return nil
			},
		},
		{
			Name:    "remove",
			Aliases: []string{"r"},
			Usage:   "Remove user from authorized_keys",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
	}

	app.Run(os.Args)
}

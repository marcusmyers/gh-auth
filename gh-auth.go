package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

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
				f, err := read_authorized_keys()
				if err != nil {
					log.Fatal(err)
				}
				arr_users := _get_users(string(f))
				fmt.Println(strings.Join(arr_users, " "))
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
				s := make([]SSHKey, 0, 10)
				for {
					if err := dec.Decode(&s); err == io.EOF {
						break
					} else if err != nil {
						log.Fatal(err)
					}
				}
				resp.Body.Close()
				str, numKeys := _return_str_keys(s, user)
				file, err := os.OpenFile(home+"/.ssh/authorized_keys", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
				if _, errFile := file.WriteString(str); errFile != nil {
					log.Fatal(errFile)
				}
				fmt.Printf("Adding %d key(s) to '%s/.ssh/authorized_keys'\n", numKeys, home)
				return nil
			},
		},
		{
			Name:    "remove",
			Aliases: []string{"r"},
			Usage:   "Remove user from authorized_keys",
			Action: func(c *cli.Context) error {
				var user = c.Args().First()
				f, err := read_authorized_keys()
				if err != nil {
					log.Fatal(err)
				}

				return nil
			},
		},
	}
	app.Run(os.Args)
}

func _return_str_keys(d []SSHKey, user string) (k string, n int) {
	var tmpStr string
	i := 0
	for key := range d {
		tmpStr += fmt.Sprintf("%s %s\n", d[key].Key, user)
		i++
	}

	return tmpStr, i
}

func _get_users(filestring string) (u []string) {
	arr_lines := strings.Split(filestring, "\n")
	users := make([]string, 0, 10)
	for i := 0; i < len(arr_lines)-1; i++ {
		arr_user_info := strings.Split(string(arr_lines[i]), " ")
		if !in_array(users, arr_user_info[2]) {
			users = append(users, arr_user_info[2])
		}
	}
	return users
}

func read_authorized_keys() (file []byte, e error) {
	users_home := os.Getenv("HOME")
	f, err := ioutil.ReadFile(users_home + "/.ssh/authorized_keys")
	if err != nil {
		return []byte{}, err
	}
	return f, nil
}

func in_array(s []string, t string) bool {
	for _, a := range s {
		if a == t {
			return true
		}
	}
	return false
}

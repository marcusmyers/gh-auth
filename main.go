package main //import "github.com/marcusmyers/gh-auth"

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/urfave/cli"
)

type SSHKey struct {
	ID  int    `json:"id"`
	Key string `json:"key"`
}

var home = os.Getenv("HOME")
var url = "https://api.github.com"

func main() {
	app := cli.NewApp()
	app.Name = "gh-auth"
	app.Version = "1.0.0"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Mark Myers",
			Email: "marcusmyers@gmail",
		},
	}
	app.HelpName = "gh-auth"
	app.Usage = "allows you to quickly pair with anyone who has a GitHub account"
	app.Commands = []cli.Command{
		{
			Name:      "list",
			Usage:     "gh-auth list",
			UsageText: "list - list all users in authorized_keys file",
			Action:    listUsers,
		},
		{
			Name:        "add",
			Description: "add github user to your authorized_keys file",
			Usage:       "gh-auth add",
			ArgsUsage:   "[github_username]",
			Action:      addUser,
		},
		{
			Name:        "remove",
			Description: "remove github user to your authorized_keys file",
			Usage:       "gh-auth remove",
			ArgsUsage:   "[github_username]",
			Action:      removeUser,
		},
	}

	app.Run(os.Args)
}

func listUsers(c *cli.Context) error {
	f, err := readAuthorizedKeys()
	if err != nil {
		log.Fatal(err)
	}
	arrUsers := _getUsers(string(f))
	fmt.Println(strings.Join(arrUsers, " "))
	return nil
}

func addUser(c *cli.Context) error {
	var user = c.Args().First()
	resp, err := http.Get(url + "/users/" + user + "/keys")
	if err != nil {
		log.Fatal(err)
	}
	dec := json.NewDecoder(resp.Body)
	s := make([]SSHKey, 0, 10)
	json_err := dec.Decode(&s)
	if json_err != nil {
		log.Fatal(json_err)
	}
	resp.Body.Close()
	str, numKeys := _returnStrKeys(s, user)
	errCreate := _writeToAuthKeys(str, home)
	if errCreate != nil {
		log.Fatal(errCreate)
	}
	fmt.Printf("Adding %d key(s) to '%s/.ssh/authorized_keys'\n", numKeys, home)
	return nil
}

func removeUser(c *cli.Context) error {
	var user = c.Args().First()
	f, err := readAuthorizedKeys()
	if err != nil {
		log.Fatal(err)
	}
	str, numKeys, errRemove := _removeUser(string(f), user, home)
	if errRemove != nil {
		log.Fatal(errRemove)
	}
	errCreate := _writeToAuthKeys(str, home)
	if errCreate != nil {
		log.Fatal(errCreate)
	}
	fmt.Printf("Removed %d key(s) to '%s/.ssh/authorized_keys'\n", numKeys, home)
	return nil
}

func readAuthorizedKeys() ([]byte, error) {
	usersHome := os.Getenv("HOME")
	f, err := ioutil.ReadFile(usersHome + "/.ssh/authorized_keys")
	if err != nil {
		return []byte{}, err
	}
	return f, nil
}

func _returnStrKeys(d []SSHKey, user string) (string, int) {
	var tmpStr string
	i := 0
	for key := range d {
		tmpStr += fmt.Sprintf("%s %s\n", d[key].Key, user)
		i++
	}

	return tmpStr, i
}

func _removeUser(filecontent string, user string, dir string) (string, int, error) {
	arrLines := strings.Split(filecontent, "\n")
	var tmpStr string
	j := 0
	for i := 0; i < len(arrLines)-1; i++ {
		if !strings.Contains(arrLines[i], user) {
			tmpStr += fmt.Sprintf("%s\n", arrLines[i])
		} else {
			j++
		}
	}
	err := os.Remove(dir + "/.ssh/authorized_keys")
	return tmpStr, j, err
}

func _getUsers(filestring string) []string {
	arrLines := strings.Split(filestring, "\n")
	users := make([]string, 0, 10)
	for i := 0; i < len(arrLines)-1; i++ {
		arrUserInfo := strings.Split(string(arrLines[i]), " ")
		if !inArray(users, arrUserInfo[2]) {
			users = append(users, arrUserInfo[2])
		}
	}
	return users
}

func _writeToAuthKeys(content string, dir string) error {
	file, err := os.OpenFile(dir+"/.ssh/authorized_keys", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, errFile := file.WriteString(content); errFile != nil {
		log.Fatal(errFile)
	}
	return nil
}

func inArray(s []string, t string) bool {
	for _, a := range s {
		if a == t {
			return true
		}
	}
	return false
}

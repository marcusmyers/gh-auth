package main

import (
	"os"
	"testing"
)

type TestSSHKey struct {
	ID  int    `json:"id"`
	Key string `json:"key"`
}

func init() {
	home = os.Getenv("HOME")
	url = "https://api.github.com"
}

func testReadAuthorizedKeys(t *testing.T) {
	content, err := readAuthorizedKeys()
	if err != nil {
		t.Error("Expecting to read the autorized_keys file, but got", err)
	}
}

func testGettingUsers(t *testing.T) {
	expected := "crwenner"
	actual := _getUsers("Admaksdue== crwenner")
	if expected != actual[0] {
		t.Error("Expecting crwenner, but got", actual)
	}
}

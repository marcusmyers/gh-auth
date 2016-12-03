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

func TestReadAuthorizedKeys(t *testing.T) {
	_, err := readAuthorizedKeys()
	if err != nil {
		t.Error("Expecting to read the autorized_keys file, but got", err)
	}
}

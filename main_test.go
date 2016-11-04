package main_test

import (
	"os"
	"testing"
)

type TestSSHKey struct {
	ID  int    `json:"id"`
	Key string `json:"key"`
}

var home string

func init() {
	home = os.Getenv("HOME")
}

func TestAddUserToAuthorizedKeys(t *testing.T) {

}

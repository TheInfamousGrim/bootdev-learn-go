package main

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

// create the method below
func (ai authenticationInfo) getBasicAuth() string {
	return fmt.Sprintf("Authorization: Basic %s:%s", ai.username, ai.password)
}

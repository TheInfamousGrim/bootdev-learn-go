package main

import "errors"

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("invalid sizes")
	}

	users := make(map[string]user)
	for i := 0; i < len(names); i++ {
		name := names[i]
		phoneNumber := phoneNumbers[i]
		users[name] = user{name, phoneNumber}
	}

	return users, nil
}

type user struct {
	name        string
	phoneNumber int
}

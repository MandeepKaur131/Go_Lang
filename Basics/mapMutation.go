package main

import (
	"errors"
	"fmt"
)

func main() {
	users := map[string]user{
		"john": {
			name:                 "john",
			number:               11111,
			scheduledForDeletion: true,
		},
		"jane": {
			name:                 "jane",
			number:               11111,
			scheduledForDeletion: false,
		},
	}
	test(users, "jane")
	test(users, "john")
	test(users, "noOne")
}

func test(users map[string]user, name string) {
	fmt.Printf("\nAttempting to delete %s...\n", name)
	deleted, err := deleteIfNecessary(users, name)
	if err != nil {
		fmt.Println(err)
		return
	}
	if deleted {
		fmt.Printf("Deleted: %s\n", name)
	}
	fmt.Printf("Did not delete: %s", name)
}

type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
}

func deleteIfNecessary(users map[string]user, name string) (delted bool, err error) {
	existingUser, ok := users[name]
	if !ok {
		return false, errors.New("Not found")
	}
	if existingUser.scheduledForDeletion {
		delete(users, name)
		return true, nil
	}
	return false, nil
}

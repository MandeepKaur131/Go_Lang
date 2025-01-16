package main

import (
	"fmt"
	"errors"
)

func getUserMap(names [] string, phoneNumbers[] int) (map[string]user, error) {
	userMap := make(map[string]user)
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("names and phoneNumbers are not of same length")
	}
	for i:=0; i<len(names); i++ {
		name := names[i]
		phoneNumber := phoneNumbers[i]
		userMap[name] = user{
			name: name,
			phoneNumber: phoneNumber,
		}
	}
	return userMap, nil
}


type user struct {
	name 			string
	phoneNumber 	int
}

func test(names[] string, phoneNumbers [] int) {
	fmt.Println("Testing getUserMap")
	defer fmt.Println("====================================")
	users, err := getUserMap(names, phoneNumbers)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, name := range names {
		fmt.Printf("Key: %v, value:\n", name)
		fmt.Printf(" - name: %v\n - phonenumber: %v\n",
		users[name].name,
		users[name].phoneNumber)

	}
}

func main() {
	test([]string{"a", "b", "c"}, []int{1, 2, 3})
	test([]string{"a", "b", "c"}, []int{1, 2})
}
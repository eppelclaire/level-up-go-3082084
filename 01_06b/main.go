package main

import (
	"encoding/json"
	"log"
	"os"
	"sort"
)

// User represents a user record.
type User struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}

const path = "users.json"

// getBiggestMarket takes in the slice of users and
// returns the biggest market.
func getBiggestMarket(users []User) (string, int) {
	markets := make(map[string]int)
	for _, user := range users {
		markets[user.Country] += 1
	}

	keys := make([]string, 0, len(markets))
	for k := range markets {
		keys = append(keys, k)
	}
	sort.SliceStable(keys, func(i, j int) bool { return markets[keys[i]] > markets[keys[j]] })

	return keys[0], markets[keys[0]]
}

func main() {
	users := importData()
	country, count := getBiggestMarket(users)
	log.Printf("The biggest user market is %s with %d users.\n",
		country, count)
}

// importData reads the raffle entries from file and
// creates the entries slice.
func importData() []User {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var data []User
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

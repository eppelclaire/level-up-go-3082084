package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"
)

const path = "./01_03b/entries.json"

// raffleEntry is the struct we unmarshal raffle entries into
type raffleEntry struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// importData reads the raffle entries from file and creates the entries slice.
func importData() []raffleEntry {
	var entries []raffleEntry
	file, err := os.ReadFile(path)
	if err != nil {
		panic("Error reading file path: " + err.Error())
	}

	err = json.Unmarshal(file, &entries)
	if err != nil {
		panic("Error converting json file: " + err.Error())
	}

	return entries
}

// getWinner returns a random winner from a slice of raffle entries.
func getWinner(entries []raffleEntry) raffleEntry {
	rand.Seed(time.Now().Unix())
	wi := rand.Intn(len(entries))
	return entries[wi]
}

func main() {
	entries := importData()
	log.Println("And... the raffle winning entry is...")
	winner := getWinner(entries)
	time.Sleep(500 * time.Millisecond)
	log.Println(winner)
}

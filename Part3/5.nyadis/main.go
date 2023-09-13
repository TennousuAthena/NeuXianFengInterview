package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

var data = make(map[string]string)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./my-db --db=/path/to/your/db")
		os.Exit(1)
	}

	dbPath := os.Args[1]

	// Load data from the database file if it exists
	loadData(dbPath)

	fmt.Println("✨Starting Nyadis CLI")
	for {
		fmt.Print("> ")
		var command string
		fmt.Scan(&command)

		switch command {
		case "help":
			fmt.Println("put <key> <value>\nget <key>\ndel <key>\nexit")
		case "put":
			var key, value string
			fmt.Scan(&key, &value)
			data[key] = value
			fmt.Println("OK")
			saveData(dbPath)
		case "get":
			var key string
			fmt.Scan(&key)
			value, ok := data[key]
			if ok {
				fmt.Printf("key:%s value:%s\n", key, value)
			} else {
				fmt.Println("NOT_FOUND")
			}
		case "del":
			var key string
			fmt.Scan(&key)
			delete(data, key)
			fmt.Println("OK")
			saveData(dbPath)
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("Invalid command")
		}
	}
}

func loadData(dbPath string) {
	file, err := os.Open(dbPath)
	if err != nil {
		fmt.Println("Error loading data:", err)
		return
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		fmt.Println("Error decoding data:", err)
	}
}

func saveData(dbPath string) {
	file, err := os.Create(dbPath)
	if err != nil {
		fmt.Println("Error saving data:", err)
		return
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	if err := encoder.Encode(data); err != nil {
		fmt.Println("Error encoding data:", err)
	}
}

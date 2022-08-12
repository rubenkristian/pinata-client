package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type PinataAuth struct {
	Key string `json:"key"`
}

func main() {
	args := os.Args[1:]

	switch args[0] {
	case "auth":
		createAuthFile(args[1])
		break
	case "list":
		listFile()
		break
	case "unpin":
		break
	}
}

func unpinByMeta() {

}

func unpin(cid string) {

}

func listFile() {

}

func createAuthFile(key string) {
	currentDir, errDir := os.Getwd()
	if errDir != nil {
		fmt.Println("Error")
		fmt.Println(errDir)
		return
	}

	if key == "" {
		authErr := fmt.Errorf("Error: must have value")
		fmt.Println(authErr)
		return
	}

	authFile := PinataAuth{
		Key: key,
	}

	file, _ := json.MarshalIndent(authFile, "", "")

	authErr := os.WriteFile(currentDir+"/pinata-auth.json", file, 0644)

	if authErr != nil {
		authFileErr := fmt.Errorf("Error: must have value")
		fmt.Println(authFileErr)
		return
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"rubenkristian.github.com/pinata-client/pinata"
)

type PinataAuth struct {
	Key string `json:"key"`
}

func main() {
	args := os.Args[1:]

	currentDir, errDir := os.Getwd()
	if errDir != nil {
		fmt.Println("Error")
		fmt.Println(errDir)
		return
	}

	fileAut, err := ioutil.ReadFile(currentDir + "/pinata-auth.json")
	if err != nil {
		fmt.Println("No file pinata-auth.json in this directory, please set auth key by using \"client auth {jwt key}\"")
		return
	}

	authPinata := &PinataAuth{}
	err = json.Unmarshal(fileAut, authPinata)
	if err != nil {
		fmt.Println("Failed to parse json file, please check file or set auth key by using \"client auth {jwt key}\"")
		return
	}

	pinataApi := pinata.CreatePinata(authPinata.Key, 0, false)
	switch args[0] {
	case "auth":
		createAuthFile(&currentDir, args[1])
	case "list":
		if len(args) >= 2 {
			listFile(pinataApi, &currentDir, args[1])
		} else {
			fmt.Println("must have name file for save result of query")
		}
	case "unpin":
		if len(args) >= 3 {
			if args[1] == "--hash" {
				unpin(pinataApi, args[2])
			} else if args[1] == "--name" {
				unpinByName(pinataApi, args[2])
			}
		} else {
			fmt.Println("unpin must have flag --hash, --name")
		}
	}
}

func unpinByName(pinata *pinata.Pinata, name string) {
	pinataBody := pinata.QueryFiles("pin", &name)

	pinata.RemoveFiles(pinataBody.Rows)
}

func unpin(pinata *pinata.Pinata, cid string) {
	pinata.RemoveByHash(cid)
}

func listFile(pinata *pinata.Pinata, currDir *string, fileName string) {
	list := pinata.QueryFiles("pin", nil)

	file, _ := json.MarshalIndent(list, "", "\t")

	saveErr := os.WriteFile((*currDir)+"/"+fileName+".json", file, 0644)

	if saveErr != nil {
		authFileErr := fmt.Errorf("Error: must have value %v", saveErr)
		fmt.Println(authFileErr)
		return
	} else {
		fmt.Println("Done to save query result to " + (*currDir) + "/" + fileName + ".json")
	}
}

func createAuthFile(currDir *string, key string) {

	if key == "" {
		fmt.Println("Error: must have value")
		return
	}

	authFile := PinataAuth{
		Key: key,
	}

	file, _ := json.MarshalIndent(authFile, "", "")

	authErr := os.WriteFile((*currDir)+"/pinata-auth.json", file, 0644)

	if authErr != nil {
		authFileErr := fmt.Errorf("Error: must have value")
		fmt.Println(authFileErr)
		return
	}
}

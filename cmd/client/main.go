package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/rubenkristian/pinata-client-api"
)

type PinataAuth struct {
	Key string `json:"key"`
}

func main() {
	args := os.Args[1:]

	if len(args) > 0 {
		currentDir, errDir := os.Getwd()
		if errDir != nil {
			fmt.Println("Error")
			fmt.Println(errDir)
			return
		}

		switch args[0] {
		case "auth":
			createAuthFile(&currentDir, args[1])
		case "list":
			pinataApi := initPinata(&currentDir)
			if len(args) >= 2 {
				listFile(pinataApi, &currentDir, args[1], args[2])
			} else {
				fmt.Println("must have name file for save result of query")
			}
		case "unpin-hash":
			pinataApi := initPinata(&currentDir)
			unpinByHash(pinataApi, args)
		case "unpin-query":
			pinataApi := initPinata(&currentDir)
			unpinByQuery(pinataApi, args)
		}
	} else {
		fmt.Println("no command")
	}
}

func initPinata(currDir *string) *pinata.Pinata {
	fileAut, err := os.ReadFile(*currDir + "/pinata-auth.json")
	if err != nil {
		fmt.Println("No file pinata-auth.json in this directory, please set auth key by using \"client auth {jwt key}\"")
		return nil
	}

	authPinata := &PinataAuth{}
	err = json.Unmarshal(fileAut, authPinata)
	if err != nil {
		fmt.Println("Failed to parse json file, please check file or set auth key by using \"client auth {jwt key}\"")
		return nil
	}

	pinataApi := pinata.CreatePinata(authPinata.Key, 0, false)

	return pinataApi
}

func unpinByHash(pinata *pinata.Pinata, params []string) {
	if len(params) == 2 {
		pinata.RemoveByHash(params[1])
	} else {
		fmt.Println("must have hash after command")
	}
}

func unpinByQuery(pinata *pinata.Pinata, params []string) {
	if len(params) >= 2 {
		pinataBody := pinata.QueryFiles(params[1])
		pinata.RemoveFiles(pinataBody.Rows)
	} else {
		fmt.Println("put your query after command")
	}
}

func listFile(pinata *pinata.Pinata, currDir *string, fileName string, query string) {
	list := pinata.QueryFiles(query)

	file, _ := json.MarshalIndent(list, "", "\t")

	saveErr := os.WriteFile((*currDir)+"/"+fileName+".json", file, 0644)

	if saveErr != nil {
		authFileErr := fmt.Errorf("error: must have value %v", saveErr)
		fmt.Println(authFileErr)
		return
	} else {
		fmt.Println("Saved query result to " + (*currDir) + "/" + fileName + ".json")
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
		authFileErr := fmt.Errorf("error: must have value")
		fmt.Println(authFileErr)
		return
	}
}

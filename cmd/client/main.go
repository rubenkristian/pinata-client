package main

import (
	"encoding/json"
	"fmt"
	"os"

	"rubenkristian.github.com/pinata-client/pinata"
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

		pinataApi := initPinata(&currentDir, args[0])

		switch args[0] {
		case "auth":
			createAuthFile(&currentDir, args[1])
		case "help":
			help()
		case "list":
			if len(args) >= 2 {
				listFile(pinataApi, &currentDir, args[1], args[2])
			} else {
				fmt.Println("must have name file for save result of query")
			}
		case "unpin-hash":
			unpinByHash(pinataApi, args)
		case "unpin-query":
			unpinByQuery(pinataApi, args)
		case "pin-file":
			pinFile(pinataApi, "")
		case "pin-json":
			pinJson(pinataApi, "")
		}
	} else {
		fmt.Println("no command")
	}
}

func initPinata(currDir *string, command string) *pinata.Pinata {
	if command == "auth" {
		return nil
	}
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

func help() {
	// TODO: help console
}

func pinFile(pinata *pinata.Pinata, fileDir string) {

}

func pinJson(pinata *pinata.Pinata, json string) {

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
		authFileErr := fmt.Errorf("Error: must have value %v", saveErr)
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
		authFileErr := fmt.Errorf("Error: must have value")
		fmt.Println(authFileErr)
		return
	}
}

// func parseArgs(params []string) (string, any) {
// 	var command string = params[0]

// 	var lenParams int = len(params)

// 	if lenParams > 1 {
// 		return "", nil
// 	} else if command == "auth" {
// 		return command, params[1]
// 	} else if command == "list" {

// 	}
// }

// func parseQuery(params []string) {
// 	var lenParams int = len(params)
// 	var index int = 0

// 	reg, _ := regexp.Compile("--([a-z]+)-([a-zA-Z0-9]+)")
// 	for index < lenParams {
// 		param := params[index]
// 		if param == "--"
// 	}
// }

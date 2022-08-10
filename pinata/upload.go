package pinata

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func (pinata *Pinata) uploadPinFile(url string, auth string, fileLoc string, name string, keyValues *[]interface{}) ([]byte, error) {
	var body []byte = nil
	payload := &bytes.Buffer{}

	writer := multipart.NewWriter(payload)

	file, errFile := os.Open(fileLoc)

	if errFile != nil {
		err := fmt.Errorf("Cannot open file, Some Error equaried %q", errFile.Error())
		return body, err
	}

	defer file.Close()

	copyFile, errFileCopy := writer.CreateFormFile("file", filepath.Base(fileLoc))
	_, errFileCopy = io.Copy(copyFile, file)

	if errFileCopy != nil {
		err := fmt.Errorf("Some Error equaried %q", errFileCopy.Error())
		return body, err
	}

	pinataOptionsJson, errParse := json.Marshal(pinata.pinataOptions)

	if errParse != nil {
		err := fmt.Errorf("Some Error equaried %q", errParse.Error())
		return body, err
	}

	pinataMetadata := &PinataMetadata{Name: name, KeyValues: keyValues}

	pinataMetadataJson, errParse := json.Marshal(pinataMetadata)

	if errParse != nil {
		err := fmt.Errorf("Some Error equaried %q", errParse.Error())
		return body, err
	}

	_ = writer.WriteField("pinataOptions", string(pinataOptionsJson))
	_ = writer.WriteField("pinataMetadata", string(pinataMetadataJson))

	errWriter := writer.Close()

	if errWriter != nil {
		err := fmt.Errorf("Some Error equaried %q", errWriter.Error())
		return body, err
	}

	client := &http.Client{}

	req, errReq := http.NewRequest(string(POST), url, payload)

	if errReq != nil {
		err := fmt.Errorf("Some Error equaried %q", errReq.Error())
		return body, err
	}

	req.Header.Add("Authorization", auth)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	res, errRes := client.Do(req)

	if errRes != nil {
		err := fmt.Errorf("Some Error equaried %q", errRes.Error())
		return body, err
	}

	defer res.Body.Close()

	body, errBody := ioutil.ReadAll(res.Body)
	if errBody != nil {
		err := fmt.Errorf("Some Error equaried %q", errBody.Error())
		return body, err
	}

	return body, nil
}

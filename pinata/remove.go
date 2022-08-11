package pinata

import (
	"io/ioutil"
	"net/http"
)

func removeFile(auth string) ([]byte, error) {
	var result []byte = nil
	clientRequest := &http.Client{}

	req, errReq := http.NewRequest(string(DELETE), string(UNPIN), nil)

	if errReq != nil {
		return nil, errReq
	}

	req.Header.Add("Authorization", auth)

	res, errRes := clientRequest.Do(req)

	if errRes != nil {
		return nil, errRes
	}

	defer res.Body.Close()

	result, errReadBody := ioutil.ReadAll(res.Body)

	if errReadBody != nil {
		return nil, errReadBody
	}

	return result, nil
}

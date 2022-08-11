package pinata

import (
	"io/ioutil"
	"net/http"
)

func (pinata *Pinata) removeFile(cid string) ([]byte, error) {
	var result []byte = nil
	clientRequest := &http.Client{}

	req, errReq := http.NewRequest(string(DELETE), string(UNPIN)+cid, nil)

	if errReq != nil {
		return nil, errReq
	}

	req.Header.Add("Authorization", pinata.authentication)

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

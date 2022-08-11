package pinata

import (
	"io/ioutil"
	"net/http"
)

func (pinata *Pinata) queryPinata(statusList string, sizeMin int32) ([]byte, error) {
	var result []byte = nil
	clientRequest := &http.Client{}

	url := string(QUERYFILES) + "?status=" + statusList

	if sizeMin > 0 {
		url += "&pinSizeMin=" + string(sizeMin)
	}

	req, errReq := http.NewRequest(string(GET), url, nil)

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

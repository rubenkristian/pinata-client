package pinata

import (
	"io/ioutil"
	"net/http"
)

type Params struct {
	Key         string
	Value       string
	SecondValue *string
	Operator    string
}

func (pinata *Pinata) queryPinata(statusList string, name *string, params []Params) ([]byte, error) {
	var result []byte = nil
	clientRequest := &http.Client{}

	url := string(QUERYFILES) + "?status=" + statusList

	if name != nil {
		url += "&metadata[name]=" + (*name)
	}

	for _, param := range params {
		if param.SecondValue == nil {
			url += "&metadata[keyvalues]={\"" + param.Key + "\":{\"value\":\"" + param.Value + "\", \"op\":\"" + param.Operator + "\"}}"
		} else {
			url += "&metadata[keyvalues]={\"" + param.Key + "\":{\"value\":\"" + param.Value + "\", \"secondValue\":\"" + *param.SecondValue + "\", \"op\":\"" + param.Operator + "\"}}"
		}
	}

	req, errReq := http.NewRequest(string(GET), url, nil)

	if errReq != nil {
		return nil, errReq
	}

	req.Header.Add("Authorization", "Bearer "+pinata.authentication)

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

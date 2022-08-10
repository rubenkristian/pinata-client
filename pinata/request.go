package pinata

type Method string

type Url string

const (
	DELETE Method = "DELETE"
	GET    Method = "GET"
	POST   Method = "POST"
	PUT    Method = "PUT"
)

const (
	PINFILE    Url = "https://api.pinata.cloud/pinning/pinFileToIPFS"
	PINJSON    Url = "https://api.pinata.cloud/pinning/pinJSONToIPFS"
	PINBYCID   Url = "https://api.pinata.cloud/pinning/pinByHash"
	LISTPIN    Url = "https://api.pinata.cloud/pinning/pinJobs"
	UPDATEMETA Url = "https://api.pinata.cloud/pinning/hashMetadata"
	UNPIN      Url = "https://api.pinata.cloud/pinning/unpin/"
	DATAUSAGE  Url = "https://api.pinata.cloud/data/userPinnedDataTotal"
	QUERYFILES Url = "https://api.pinata.cloud/data/pinList"
)

type Request struct {
	url            string
	authentication string
	params         []string
}

func createRequest(url string, auth string, params []string) *Request {
	return &Request{
		url:            url,
		authentication: auth,
		params:         params,
	}
}

func (req *Request) PinFile() string {
	_, err := uploadPinFile(req.url, req.authentication, req.params[0], "", "")

	if err != nil {
		return "Error Pin File"
	}

	return "Successful Pin File"
}

func (*Request) PinJSON() {

}

func (*Request) PinByCID() {

}

func (*Request) LastPinByCID() {

}

func (*Request) UpdateMetaData() {

}

func (*Request) RemoveFiles() {

}

func (*Request) DataUsage() {

}

func (*Request) QueryFiles() {

}

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

type Pinata struct {
	url            string
	authentication string
	pinataOptions  *PinataOptions
}

type PinataOptions struct {
	CidVersion        int8 `json:"cidVersion"`
	WrapWithDirectory bool `json:"wrapWithDirectory"`
}

type PinataMetadata struct {
	Name      string         `json:"name"`
	KeyValues *[]interface{} `json:"keyvalues"`
}

func createRequest(url string, auth string, cidVersion int8, wrapWithDirectory bool) *Pinata {
	return &Pinata{
		url:            url,
		authentication: auth,
		pinataOptions:  &PinataOptions{CidVersion: cidVersion, WrapWithDirectory: wrapWithDirectory},
	}
}

func (pinata *Pinata) PinFile(fileLoc string, name string, keyvalues *[]interface{}) string {
	_, err := pinata.uploadPinFile(pinata.url, pinata.authentication, fileLoc, name, keyvalues)

	if err != nil {
		return "Error Pin File"
	}

	return "Successful Pin File"
}

func (pinata *Pinata) PinJSON() {

}

func (pinata *Pinata) PinByCID() {

}

func (pinata *Pinata) LastPinByCID() {

}

func (pinata *Pinata) UpdateMetaData() {

}

func (pinata *Pinata) RemoveFiles() {

}

func (pinata *Pinata) DataUsage() {

}

func (pinata *Pinata) QueryFiles() {

}

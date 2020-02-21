package models


/***
 * version
 */
// swagger:model
type VersionResponse struct {
	Version string `json:"version"`
	Core    string `json:"core"`
	Raw     string `json:"raw"`
	Note    string `json:"note"`
}

/***
 * rsa key peer
 */
// swagger:model
type RsaKeyPeer struct {
	PrivateKey string `json:"private"`
	PublicKey string `json:"public"`
	Raw     string `json:"raw"`
	Note    string `json:"note"`
}


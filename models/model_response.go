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

/**
 * 使用 private 解密后的明文
 */
type DeRsa struct {
	Plaintext string `json:"plaintext"`
	Raw     string `json:"raw"`
	Note    string `json:"note"`
}

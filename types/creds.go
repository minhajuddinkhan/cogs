package types

// Credentials Credentials
type Credentials struct {
	Hash        []byte
	PublicKey   []byte
	PrivateKey  []byte
	AccessToken string
	EmployeeID  int
}

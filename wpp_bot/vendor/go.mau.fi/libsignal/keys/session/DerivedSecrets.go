package session

// NewDerivedSecrets returns a new RootKey/ChainKey pair from 64 bytes of key material
// generated by the key derivation function.
func NewDerivedSecrets(keyMaterial []byte) *DerivedSecrets {
	secrets := DerivedSecrets{
		keyMaterial[:32],
		keyMaterial[32:],
	}

	return &secrets
}

// DerivedSecrets is a structure for holding the derived secrets for the
// Root and Chain keys for a session.
type DerivedSecrets struct {
	rootKey  []byte
	chainKey []byte
}

// RootKey returns the RootKey bytes.
func (d *DerivedSecrets) RootKey() []byte {
	return d.rootKey
}

// ChainKey returns the ChainKey bytes.
func (d *DerivedSecrets) ChainKey() []byte {
	return d.chainKey
}

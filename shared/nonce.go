package shared

type Nonce struct {
	nonces map[string]bool
}

func NewNonce() *Nonce {
	return &Nonce{
		nonces: make(map[string]bool),
	}
}

func (n *Nonce) Register(guid string) bool {
	_, ok := n.nonces[guid]
	if !ok {
		n.nonces[guid] = true
		return true
	}
	return false
}

package model

type Secrets struct {
	JWTKey	[]byte
}

func NewSecrets() *Secrets {
	return &Secrets{
		JWTKey: make([]byte, 32),
	}
}

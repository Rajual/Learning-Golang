package main

import "fmt"

type PasswordProtector struct {
	user          string
	passwordName  string
	hashAlgorithm HashAlgorithm
}

type HashAlgorithm interface {
	Hash(p *PasswordProtector)
}

func NewPasswordProtector(user string, passwordName string, hash HashAlgorithm) *PasswordProtector {
	return &PasswordProtector{
		user:          user,
		passwordName:  passwordName,
		hashAlgorithm: hash,
	}
}

func (p *PasswordProtector) SetHashAlgorithm(hash HashAlgorithm) {
	p.hashAlgorithm = hash
}

func (p *PasswordProtector) Hash() {
	p.hashAlgorithm.Hash(p)
}

type SHA struct{}

func (SHA) Hash(p *PasswordProtector) {
	fmt.Printf("Hashing using SHA for %s", p.passwordName)
	fmt.Println("")
}

type MD5 struct{}

func (MD5) Hash(p *PasswordProtector) {
	fmt.Printf("Hashing using MD5 for %s", p.passwordName)
	fmt.Println("")
}

func main() {
	sha := &SHA{}
	md5 := &MD5{}

	passwordProtector := NewPasswordProtector("julio", "gmail password", sha)
	passwordProtector.Hash()
	passwordProtector.SetHashAlgorithm(md5)
	passwordProtector.Hash()

}

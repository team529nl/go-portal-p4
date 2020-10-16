package go_portal_p4

import (
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"github.com/markbates/pkger/pkging"
	//"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
)

func LoadCertificate(certFile pkging.File) (*x509.Certificate, error) {
	certPem, e := ioutil.ReadAll(certFile)
	if e != nil {
		return nil, e
	}

	cert, err := parsePublicCertificate(certPem)
	if err != nil {
		return nil, err
	}

	return cert, nil
}

func LoadX509KeyPair(certFile, keyFile string) (*x509.Certificate, *rsa.PrivateKey, error) {
	certPem, e := ioutil.ReadFile(certFile)
	if e != nil {
		return nil, nil, e
	}

	cert, err := parsePublicCertificate(certPem)
	if err != nil {
		return nil, nil, err
	}

	keyFileContent, e := ioutil.ReadFile(keyFile)
	if e != nil {
		return nil, nil, e
	}

	key, err := parsePrivateKey(keyFileContent)
	if err != nil {
		return nil, nil, err
	}

	return cert, key, nil
}

func parsePrivateKey(pemBytes []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return nil, errors.New("ssh: no key found")
	}

	var result *rsa.PrivateKey
	switch block.Type {
	case "PRIVATE KEY":
		parsed, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		result = parsed.(*rsa.PrivateKey)
	default:
		return nil, fmt.Errorf("ssh: unsupported key type %q", block.Type)
	}

	return result, nil
}

func parsePublicCertificate(pemBytes []byte) (*x509.Certificate, error) {
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return nil, errors.New("ssh: no key found")
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, err
	}

	return cert, nil
}

func ToPem(certificate *x509.Certificate) []byte {
	publicKeyBlock := pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certificate.Raw,
	}
	publicKeyPem := pem.EncodeToMemory(&publicKeyBlock)
	return publicKeyPem
}

func ToBase64(cert *x509.Certificate) (string, error) {
	return base64.StdEncoding.EncodeToString(cert.Raw), nil
}

func Sha1Hash(input []byte) []byte {
	hash := crypto.SHA1.New()
	hash.Write(input)
	hashed := hash.Sum(nil)
	return hashed
}

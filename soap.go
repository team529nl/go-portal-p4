package go_portal_p4

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"github.com/beevik/etree"
	"github.com/markbates/pkger"
	dsig "github.com/russellhaering/goxmldsig"
)

var endpointBaseToCertificate = map[string]string{
	"https://pp4-test.edsn.nl:443": "/certificates/sign.pp4-test.edsn.nl.cer",
	"https://p4.edsn.nl:443":       "/certificates/sign.p4.edsn.nl.cer",
}

func init() {

	pkger.Include("/certificates/sign.pp4-test.edsn.nl.cer")
	pkger.Include("/certificates/sign.p4.edsn.nl.cer")
}

type Header struct {
	Security Security `xml:"http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd Security"`
}

type Security struct {
	MustUnderstand      string              `xml:"http://schemas.xmlsoap.org/soap/envelope/ mustUnderstand,attr"`
	BinarySecurityToken BinarySecurityToken `xml:"http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd BinarySecurityToken"`
	Signature           Signature           `xml:"http://www.w3.org/2000/09/xmldsig# Signature"`
}

type BinarySecurityToken struct {
	Id           string `xml:"http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd Id,attr"`
	EncodingType string `xml:",attr"`
	ValueType    string `xml:",attr"`
	Token        string `xml:",innerxml"`
}

type Signature struct {
	XMLName        xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# Signature"`
	SignedInfo     SignedInfo
	SignatureValue SignatureValue
	KeyInfo        KeyInfo
}

type SignatureValue struct {
	Signature string `xml:",innerxml"`
}

type SignedInfo struct {
	XMLName                xml.Name            `xml:"http://www.w3.org/2000/09/xmldsig# SignedInfo"`
	CanonicalizationMethod Algorithm           `xml:"CanonicalizationMethod"`
	SignatureMethod        Algorithm           `xml:"SignatureMethod"`
	Reference              SignedInfoReference `xml:"Reference"`
}

type SignedInfoReference struct {
	URI          string `xml:",attr"`
	Transforms   Transforms
	DigestMethod Algorithm `xml:"DigestMethod"`
	DigestValue  DigestValue
}

type Transforms struct {
	Transform []Algorithm `xml:",innerxml"`
}

type DigestValue struct {
	Digest string `xml:",innerxml"`
}

type KeyInfo struct {
	SecurityTokenReference SecurityTokenReference `xml:"http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd SecurityTokenReference"`
}

type SecurityTokenReference struct {
	Reference SecurityTokenReferenceInner `xml:"http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd Reference"`
}

type SecurityTokenReferenceInner struct {
	URI       string `xml:",attr"`
	ValueType string `xml:",attr"`
}

type Algorithm struct {
	Value string `xml:"Algorithm,attr"`
}

type Body struct {
	XMLName xml.Name    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
	Id      string      `xml:"Id,attr"`
	Message interface{} `xml:",any,omitempty"`
	Fault   *struct {
		String string `xml:"faultstring,omitempty"`
		Code   string `xml:"faultcode,omitempty"`
		Detail string `xml:"detail,omitempty"`
	} `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault,omitempty"`
}

type SoapEnvelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Header  Header   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`
	Body    Body     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
}

func CanonicalizeEtree(reference *etree.Element) ([]byte, error) {
	canonicalizer := dsig.MakeC14N10ExclusiveCanonicalizerWithPrefixList("")
	canonicalized, err := canonicalizer.Canonicalize(reference)
	if err != nil {
		return nil, err
	}
	return canonicalized, nil
}

func Canonicalize(xml []byte) ([]byte, error) {
	doc := etree.NewDocument()
	if err := doc.ReadFromBytes(xml); err != nil {
		return nil, err
	}

	canonicalized, err := CanonicalizeEtree(doc.Root())
	if err != nil {
		return nil, err
	}

	return canonicalized, nil
}

type Signer struct {
	Certificate *x509.Certificate
	PrivateKey  *rsa.PrivateKey
}

func calculateHashFromEtree(etree *etree.Element) ([]byte, error) {
	canonicalBodyXML, err := CanonicalizeEtree(etree)
	if err != nil {
		return nil, err
	}

	hash := Sha1Hash(canonicalBodyXML)
	return hash, nil
}

func calculateHash(obj interface{}) ([]byte, error) {
	objXML := ObjToXML(obj)
	canonicalBodyXML, err := Canonicalize(objXML.Bytes())
	if err != nil {
		return nil, err
	}

	hash := Sha1Hash(canonicalBodyXML)
	return hash, nil
}

func (signer *Signer) Sign(body *Body) (*Signature, error) {

	err := signer.PrivateKey.Validate()
	if err != nil {
		return nil, err
	}

	hash, err := calculateHash(body)
	if err != nil {
		return nil, err
	}
	digest := base64.StdEncoding.EncodeToString(hash)

	signedInfo := SignedInfo{
		CanonicalizationMethod: Algorithm{Value: "http://www.w3.org/2001/10/xml-exc-c14n#"},
		SignatureMethod:        Algorithm{Value: "http://www.w3.org/2000/09/xmldsig#rsa-sha1"},
		Reference: SignedInfoReference{
			URI:          "#" + body.Id,
			Transforms:   Transforms{Transform: []Algorithm{{Value: "http://www.w3.org/2001/10/xml-exc-c14n#"}}},
			DigestMethod: Algorithm{Value: "http://www.w3.org/2000/09/xmldsig#sha1"},
			DigestValue:  DigestValue{Digest: digest},
		},
	}

	signedInfoXML := ObjToXML(signedInfo)
	canonicalSignedInfoXML, err := Canonicalize(signedInfoXML.Bytes())
	if err != nil {
		return nil, err
	}
	hash = Sha1Hash(canonicalSignedInfoXML)

	signedHash, err := rsa.SignPKCS1v15(rand.Reader, signer.PrivateKey, crypto.SHA1, hash)
	if err != nil {
		return nil, err
	}

	signatureValue := SignatureValue{Signature: base64.StdEncoding.EncodeToString(signedHash)}

	signature := Signature{
		SignedInfo:     signedInfo,
		SignatureValue: signatureValue,
		KeyInfo: KeyInfo{SecurityTokenReference: SecurityTokenReference{Reference: SecurityTokenReferenceInner{
			URI:       "#x509-1",
			ValueType: "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-x509-token-profile-1.0#X509v3",
		}}},
	}

	return &signature, nil
}

type Validator struct {
	cert x509.Certificate
}

func NewValidator(endpointBase string) (*Validator, error) {
	certPath, ok := endpointBaseToCertificate[endpointBase]
	if !ok {
		return nil, fmt.Errorf("cannot find certificate for %s", endpointBase)
	}
	file, err := pkger.Open(certPath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	cert, err := LoadCertificate(file)
	if err != nil {
		return nil, err
	}

	return &Validator{cert: *cert}, nil
}

func (validator *Validator) ValidateXmlString(xml string) error {
	doc := etree.NewDocument()
	if err := doc.ReadFromString(xml); err != nil {
		return err
	}

	signature := doc.FindElement(".//Signature")
	if signature == nil {
		return fmt.Errorf("no signature found")
	}

	signedInfo := signature.SelectElement("SignedInfo")
	if signedInfo == nil {
		return fmt.Errorf("no SignedInfo found")
	}

	references := signedInfo.FindElements("./Reference")
	if references == nil || len(references) != 1 {
		return fmt.Errorf("references count != 1")
	}

	digestValueElement := references[0].SelectElement("DigestValue")
	if digestValueElement == nil {
		return fmt.Errorf("DigestValue is missing")
	}

	digestValue := digestValueElement.Text()

	body := doc.FindElement(".//Body")
	if body == nil {
		return fmt.Errorf("body element is missing")
	}
	for _, attr := range doc.Root().Attr {
		body.CreateAttr(attr.FullKey(), attr.Value)
	}
	hash, err := calculateHashFromEtree(body)
	if err != nil {
		return err
	}
	calculatedDigest := base64.StdEncoding.EncodeToString(hash)

	if digestValue != calculatedDigest {
		return fmt.Errorf("calculated digest [ %s ] does not match the expected digest [ %s ]", calculatedDigest, digestValue)
	}

	signatureValueElement := signature.SelectElement("SignatureValue")
	if signatureValueElement == nil {
		return fmt.Errorf("SignatureValue is missing")
	}

	signatureValue := signatureValueElement.Text()
	sig, err := base64.StdEncoding.DecodeString(signatureValue)
	if err != nil {
		return err
	}

	for _, attr := range signature.Attr {
		signedInfo.CreateAttr(attr.FullKey(), attr.Value)
	}

	if err != nil {
		return err
	}

	canonSignedInfoHash, err := calculateHashFromEtree(signedInfo)
	if err != nil {
		return err
	}

	err = rsa.VerifyPKCS1v15(validator.cert.PublicKey.(*rsa.PublicKey), crypto.SHA1, canonSignedInfoHash, sig)

	if err != nil {
		return err
	}

	return nil
}

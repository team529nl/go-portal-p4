package go_portal_p4

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os/exec"
	"reflect"
	"runtime"
	"team529.nl/go-portal-p4/portal-p4/messages"
)

func AssertPointer(v interface{}, message string) error {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Ptr {
		return errors.New(message)
	}
	return nil
}

func PrettyPrint(data interface{}) {
	var p []byte
	//    var err := error
	p, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s \n", p)
}

type TestConfig struct {
	Env             string
	CertificatePath string
	PrivateKeyPath  string
	SenderEAN       string
	ReceiverEAN     string
	EanCodes        []messages.GSRNEANCode
}

func OpenBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}

package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
	"net/http/httputil"
	"team529.nl/go-portal-p4"
	"team529.nl/go-portal-p4/portal-p4"
	"team529.nl/go-portal-p4/portal-p4/messages"
	"time"
)

var testConfig go_portal_p4.TestConfig

func init() {
	test := go_portal_p4.TestConfig{
		Env:             "test",
		CertificatePath: "/path/to/cert.pem",
		PrivateKeyPath:  "/path/to/private.key",
		SenderEAN:       "[EAN13]",
		ReceiverEAN:     "[EAN13]",
		EanCodes:        []messages.GSRNEANCode{"[EAN18]"},
	}
	_ = test

	testConfig = test
}

func main() {

	log.Println("ENV:" + testConfig.Env)

	cert, privateKey, err := go_portal_p4.LoadX509KeyPair(testConfig.CertificatePath, testConfig.PrivateKeyPath)

	if err != nil {
		log.Fatal(err)
	}

	soapClient, err := portal_p4.NewClient(testConfig.Env, *cert, *privateKey)
	if err != nil {
		log.Fatal(err)
	}

	soapClient.RequestHook = func(req *http.Request) *http.Request {
		data, err := httputil.DumpRequest(req, true)
		_ = data
		if err != nil {
			panic(err)
		}
		//log.Println(string(data))
		return req
	}

	soapClient.ResponseHook = func(rsp *http.Response) *http.Response {
		data, err := httputil.DumpResponse(rsp, true)
		_ = data
		if err != nil {
			panic(err)
		}
		//log.Println(string(data))
		return rsp
	}

	messageFactory := portal_p4.NewMessageFactory(testConfig.SenderEAN, testConfig.ReceiverEAN, uuid.New().String()) //prod

	message := messageFactory.DataRequest(testConfig.EanCodes, //prod
		[]portal_p4.Date{
			{Year: 2019, Month: time.May, Day: 1},
			{Year: 2019, Month: time.June, Day: 1},
			{Year: 2019, Month: time.July, Day: 1},
			{Year: 2019, Month: time.August, Day: 1},
			{Year: 2019, Month: time.September, Day: 1},
			{Year: 2019, Month: time.October, Day: 1},
			{Year: 2019, Month: time.November, Day: 1},
			{Year: 2019, Month: time.December, Day: 1},
			{Year: 2020, Month: time.January, Day: 1},
			{Year: 2020, Month: time.February, Day: 1},
			{Year: 2020, Month: time.March, Day: 1},
			{Year: 2020, Month: time.April, Day: 1},
			{Year: 2020, Month: time.May, Day: 1},
		},
		messages.Maandstand_recovery)

	//message := messageFactory.DataRequest(testConfig.EanCodes,
	//	[]p4.Date{{
	//		Year:  2020,
	//		Month: time.April,
	//		Day:   25,
	//	}}, messages.Dagstand)

	response, err := soapClient.DatabatchRequest(context.TODO(), message)
	_ = response

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("RESPONSE")
	go_portal_p4.PrettyPrint(response)

}

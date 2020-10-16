package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"sort"
	go_portal_p4 "team529.nl/go-portal-p4"
	portal_p4 "team529.nl/go-portal-p4/portal-p4"
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
		EanCodes:        nil,
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
		log.Println(string(data))
		return rsp
	}

	messageFactory := portal_p4.NewMessageFactory(testConfig.SenderEAN, testConfig.ReceiverEAN, "TestRun")
	message := messageFactory.DataResultRequest()

	results := make(map[string]string)
	_ = results

	for {
		log.Printf("Requesting data ...\n")
		response, err := soapClient.DatabatchResult(context.TODO(), message)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Result size: %d\n", len(response.P4Content.P4MeteringPoint))

		for _, mp := range response.P4Content.P4MeteringPoint {
			//go_portal_p4.PrettyPrint(v)
			if mp.P4EnergyMeter != nil {
				for _, meter := range *mp.P4EnergyMeter {
					//fmt.Printf("* Meternummer %s\n", meter.ID)
					for _, measurement := range meter.P4Register {
						for _, reading := range measurement.P4Reading {
							results[fmt.Sprintf("%s,%s,%s", mp.EANID,
								(time.Time)(reading.ReadingDateTime).Format("2006-01-02T15:04Z0700"),
								measurement.ID)] =
								fmt.Sprintf("%f,%s", reading.Reading, measurement.MeasureUnit)
						}
					}
				}
			} else {
				log.Printf("! Error: %s\n", (*mp.P4Rejection)[0].Rejection.RejectionText)
			}
		}

		if len(response.P4Content.P4MeteringPoint) == 0 || len(response.P4Content.P4MeteringPoint) > 1 {
			break
		}

		time.Sleep(10 * time.Millisecond)
	}

	if len(results) > 0 {
		keys := make([]string, 0, len(results))
		for k := range results {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		f, _ := ioutil.TempFile(os.TempDir(), "p4-data-*.txt")
		fmt.Printf("EAN,Timestamp,MeasuremenType,Reading,Unit\n")
		if f != nil {
			_, _ = f.WriteString("EAN,Timestamp,MeasuremenType,Reading,Unit\n")
		}

		for _, k := range keys {
			fmt.Printf("%s, %s\n", k, results[k])
			if f != nil {
				_, _ = f.WriteString(fmt.Sprintf("%s,%s\n", k, results[k]))
			}
		}

		if f != nil {
			f.Close()
			log.Printf("File written to %s\n", f.Name())
		}
	}
}

package helpers

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

// Utility function to fetch the input. Common for all questions
func FetchInput(fileName string) []byte {
	inputListUrl, err := url.Parse("https://adventofcode.com/2024/day/1/input")
	if err != nil {
		log.Fatalln("Error parsing input url")
	}

	// Get the session cookie from the env. Probably going to break when the
	// cookie expires on 2026-01-05T19:32:38.052Z, but was fun to try nonetheless
	session, ok := os.LookupEnv("session")
	if !ok {
		log.Fatalln("No env variable \"session\" found")
	}

	headers := http.Header{}
	headers.Add("Cookie", fmt.Sprintf("session=%s", session))
	req := http.Request{
		Method: "GET",
		URL:    inputListUrl,
		Header: headers,
	}

	resp, err := http.DefaultClient.Do(&req)
	if err != nil {
		log.Fatalln("Unable to fetch the input list")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Unable to read the response body")
	}

	// TODO: Move this check before making the request so that if the file exists,
	// we do not make any requests
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		if !errors.Is(err, os.ErrExist) {
			log.Fatalln("Error creating input file", err)
		}
	}
	if file != nil {
		// If the input File already exists, the value of `file` will be nil
		bytesWritten, err := file.Write(body)
		if err != nil {
			log.Fatalf("Error writing %d bytes to the input file", bytesWritten)
		}
		// log.Printf("Wrote %d bytes to the intput file\n", bytesWritten)
	}

	body, err = os.ReadFile(fileName)
	if err != nil {
		log.Fatalln("Error reading input file")
	}
	return body
}

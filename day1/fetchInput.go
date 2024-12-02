package day1

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func fetchInput() {
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

	file, err := os.OpenFile(FILE_NAME, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		if errors.Is(err, os.ErrExist) {
			// File already exists, do nothing and return
			return
		}
		log.Fatalln("Error creating input file", err)
	}
	// bytesWritten, err := file.Write(body)
	_, err = file.Write(body)
	if err != nil {
		log.Fatalln("Error writing to the input file")
	}
	// log.Printf("Wrote %d bytes to the intput file\n", bytesWritten)
}

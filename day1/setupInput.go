package day1

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func setupInput() (col1, col2 []int) {
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

	body, err = os.ReadFile(FILE_NAME)
	if err != nil {
		log.Fatalln("Error reading input file")
	}

	inputString := string(body)
	lines := strings.Split(inputString, "\n")
	for _, line := range lines {
		input := strings.Fields(line)
		if len(input) == 0 {
			// Reached end of file
			break
		}
		if len(input) != 2 {
			log.Fatalln("Invalid input")
		}

		// Convert the values to int
		var1, err1 := strconv.Atoi(input[0])
		var2, err2 := strconv.Atoi(input[1])
		if err1 != nil || err2 != nil {
			log.Fatalln("Error parsing the value")
		}

		col1 = append(col1, var1)
		col2 = append(col2, var2)
	}
	return
}

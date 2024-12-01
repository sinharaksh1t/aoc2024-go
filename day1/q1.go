package day1

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Q1() {
	inputListUrl, err := url.Parse("https://adventofcode.com/2024/day/1/input")
	if err != nil {
		log.Fatalln("Error parsing input url")
	}

	// Get the session cookie from the env. Probably goingt to break when the
	// cookie expires in the future but was fun to try nonetheless
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

	var col1, col2 []int
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

	// Sort the two columns
	// I know O(n log(n)), eww
	sort.Ints(col1)
	sort.Ints(col2)

	var result int = 0
	for i := 0; i < 1000; i++ {
		val1 := col1[i]
		val2 := col2[i]

		result += int(math.Abs(float64(val1 - val2)))
	}
	fmt.Printf("Result is: %d", result)
}

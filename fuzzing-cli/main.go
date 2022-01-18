package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	url := flag.String("url", "github.com", "will be fuzzing word")
	respCode := flag.Int("respCode", 404, "expected response code")
	flag.Parse()

	file, err := parseFile("./assets/letter_wl.txt")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("URL = %s\n", *url)
	fmt.Printf("Code = %d\n", *respCode)
	startFuzzing(file, *respCode, *url)
}

func parseFile(filePath string) (*os.File, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	return file, err
}

func startFuzzing(file *os.File, respCode int, url string) {
	//when it reads all lines, the file will be closed.
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K
	for scanner.Scan() {
		wantedURL := fmt.Sprintf("https://%s/%s", url, scanner.Text())
		resp, err := http.Get(wantedURL)
		if err != nil {
			log.Fatalln(err)
		}

		if resp.StatusCode == respCode {
			fmt.Printf("founded = %s\n", scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

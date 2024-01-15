package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"math"
	"net/http"
	"os"
	"regexp"
	"time"
)

// Initialize global maps
var domainUpTotal = map[string]int{}
var domainIterations = map[string]int{}
var domainUpRatio = map[string]float64{}

// RequestDetails holds data from the provided yaml file
type RequestDetails struct {
	Name        string            `yaml:"name"`
	EndpointURL string            `yaml:"url"`
	Method      string            `yaml:"method,omitempty"`
	Headers     map[string]string `yaml:"headers,omitempty"`
	Body        string            `yaml:"body,omitempty"`
}

// parseRequestDetails accepts the filename provided and extracts the data
// from the file into a list of RequestDetails structs. It then returns that list.
func parseRequestDetails(f string) []RequestDetails {
	var requestData []RequestDetails
	yamlContents, err := os.ReadFile(f)
	if err != nil {
		log.Fatal("Invalid Filename Error: \n", err)
	}
	err = yaml.Unmarshal(yamlContents, &requestData)
	if err != nil {
		log.Fatal("Error Processing File: \n", err)
	}
	return requestData
}

// HttpRequest accepts a RequestDetails struct and uses it to creates an HTTP request
// It returns both the HTTP response and calculated latency for the request
func HttpRequest(d RequestDetails) (*http.Response, int64) {
	client := &http.Client{}
	req, _ := http.NewRequest(d.Method, d.EndpointURL, nil)
	for k, v := range d.Headers {
		req.Header.Set(k, v)
	}
	startTime := time.Now()
	res, _ := client.Do(req)
	defer res.Body.Close()
	endTime := time.Now()
	latency := endTime.Sub(startTime)
	return res, latency.Milliseconds()
}

// parseDomain strips the domain name from the provided URL
func parseDomain(u string) string {
	re := regexp.MustCompile(`\/\/(.*?)\/`)
	return re.FindStringSubmatch(u)[1]
}

// createDomainKeysInMaps checks if a domain exists as a key and creates
// and initializes this key as required by the 3 maps if needed
func createDomainKeysInMaps(d string) {
	_, ok := domainUpTotal[d]
	if !ok {
		domainUpTotal[d] = 0
		domainIterations[d] = 1
		domainUpRatio[d] = 0
	}
}

// checkStatusCode pulls the HTTP status code from the response and
// iterates the UP total for the domain if the response is 2xx
// and the latency is under 500 ms
func checkStatusCode(c *http.Response, l int64, d string) {
	if c.StatusCode > 199 && c.StatusCode < 300 && l < 500 {
		domainUpTotal[d] += 1
	}
}

// calculateAvailability divides the number of times a domain is found
// to be UP by the total number of times requests were sent to it
// and saves this ratio
func calculateAvailability(d string) {
	domainUpRatio[d] = float64(domainUpTotal[d]) / float64(domainIterations[d]) * 100
}

func main() {
	var filepath string

	fmt.Println("Please enter a valid path for the yaml file containing the HTTP request data:")
	_, err := fmt.Scanln(&filepath)
	if err != nil {
		log.Fatal("Invalid Input: \n", err)
	}
	requestDetails := parseRequestDetails(filepath)

	for {
		// Iterates through each request detailed in the yml file, and manipulates
		// the data as needed
		for i := range requestDetails {
			response, latency := HttpRequest(requestDetails[i])
			domain := parseDomain(requestDetails[i].EndpointURL)

			createDomainKeysInMaps(domain)
			checkStatusCode(response, latency, domain)
			calculateAvailability(domain)

			domainIterations[domain]++
		}

		// Logs the results to the console
		for k, v := range domainUpRatio {
			fmt.Printf("%s has %0.0f%% availability percentage\n", k, math.Round(v))
		}

		// Waits 15 seconds before next iteration of HTTP requests
		time.Sleep(15 * time.Second)
	}
}

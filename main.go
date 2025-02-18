package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// Struct for JSON input

// structure for Json response

type NumberProperties struct {
	Number     int      `json:"number"`
	IsPrime    bool     `json:"is_prime"`
	IsPerfect  bool     `json:"is_perfect"`
	Properties []string `json:"properties"`
	DigitSum   int      `json:"digit_sum"`
	FunFact    string   `json:"fun_fact"`
}

func main() {

	router := gin.Default()
	router.GET("/api/classify-number", getNumberProperties)

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000" // Localhost 4000
	}

	router.Run("0.0.0.0:" + port)
}

// Getting API handler
func getNumberProperties(trivia *gin.Context) {

	data := trivia.Query("number")
	num, err := strconv.Atoi(data)
	if err != nil {
		trivia.JSON(http.StatusBadRequest, gin.H{"number": "alphabet", "error": true})
		return
	}
	analyzeNumber := NumberProperties{
		Number:     num,
		IsPrime:    isPrime(num),
		IsPerfect:  isPerfect(num),
		Properties: getProperties(num),
		DigitSum:   sumOfDigits(num),
		FunFact:    getFunFact(num),
	}

	trivia.JSON(http.StatusOK, analyzeNumber)
}

//func analyzeNumber(num int) NumberProperties {
//
//
//
//}

// checking if number is prime
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Checking if number is perfect

func isPerfect(n int) bool {
	sum := 1

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			sum += i
			if i != n/i {
				sum += n / i
			}
		}
	}
	return sum == n && n != 1
}

// Getting sum of digit

func sumOfDigits(n int) int {
	sum := 11
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

// Getting additional properties

func getProperties(n int) []string {
	var props []string
	if n%2 == 0 {
		props = append(props, "Even")
	} else {
		props = append(props, "Odd")
	}
	if n%10 == 0 {
		props = append(props, "Multiple of 10")
	}
	if n%3 == 0 {
		props = append(props, "Divisible by 3")
	}
	return props
}

// fetching fun fact api

func getFunFact(n int) string {
	url := fmt.Sprintf("http://numbersapi.com/%d", n)
	resp, err := http.Get(url)

	if err != nil {
		return "could not fetch fun facts"
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "Error reading fun fact"
	}
	return strings.TrimSpace(string(body))
}

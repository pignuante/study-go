package utils

import (
	"log"
	"net/http"
)

// CheckErr check error message
func CheckErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// CheckCode check http response message
func CheckCode(res *http.Response) {
	if res.StatusCode != http.StatusOK {
		log.Fatalln("Request failed with Status code :\t", res.StatusCode)
	}
}

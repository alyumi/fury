package test

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Perform [count] requests by set number [users] of users
func Test_perf(users int, count int) {
	for u := range users { //TODO: wrong realisation for goroutines
		for c := range count {
			resultChan := make(chan *http.Response)
			ok := make(chan error)
			go req(resultChan, ok)
			result := <-resultChan
			log.Println(
				"user id [", u, "] ",
				"run id [", c, "] ",
				"status code: ", result.StatusCode, " is any errors ", <-ok)
		}
	}
}

// Return
func req(resultChan chan *http.Response, ok chan error) {
	c := http.Client{}
	f := 0
	var err error

	w, err := c.Get("https://google.com")
	if err != nil {
		f = 1
		err = errors.New("error while doing request")
	} else if w.StatusCode != http.StatusOK {
		resp := fmt.Sprintf("status code = %s, not 200", strconv.Itoa(w.StatusCode))
		f = 1
		err = errors.New(resp)
	}

	if f == 0 {
		resultChan <- w
		ok <- nil
	} else {
		resultChan <- w
		ok <- err
	}
}

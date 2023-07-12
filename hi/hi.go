package hi

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type User struct {
	Name   string   `json:"name"`
	Age    int      `json:"age"`
	Movies []string `json:"movies"`
}

func Hello(input string) string {
	u := &User{}
	json.Unmarshal([]byte(input), u)
	u.Name += "!!"
	u.Age += 10

	out, _ := json.Marshal(u)
	return string(out)
}

func Hello2(input User) User {
	input.Name += "!!"
	input.Age += 10
	return input
}

func Hello3() string {
	resp, _ := http.Get("http://httpbin.org/ip")

	// Close the response body when we're done
	defer resp.Body.Close()

	// Read the response body
	body, _ := ioutil.ReadAll(resp.Body)

	// Extract the IP address from the response
	ip := string(body)

	return ip
}

func NewContext() context.Context {
	return context.Background()
}

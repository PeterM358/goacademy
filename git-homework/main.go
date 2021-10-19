package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type User struct {
	Login     string `json:"login"`
	Followers int    `json:"followers"`
	Repos     Repo
}

type Repo struct {
	ForksCount int `json:"forks_count"`
	UpdatedAt string `json:"updated_at"`
}

var LeftUrl = "https://api.github.com/users/"
var MidUrl = "/repos"

func main() {
	//file, err := ioutil.ReadFile("usernames.txt")
	file, err := os.Open("usernames.txt")
	if err != nil {
		log.Fatal(err)
	}
	//defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	if scanner.Scan() != true {
		fmt.Println("File is empty")
	}

	var usernames []string
	for scanner.Scan() {
		usernames = append(usernames, scanner.Text())
	}
	fmt.Printf("Usernames from file: %+v\n", usernames) // TODO is working now

	for _, user := range usernames { // TODO wrap in func
		//user request
		resp, err := http.Get(LeftUrl + user)
		if err != nil {
			log.Fatal(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		var result User
		if err := json.Unmarshal([]byte(body), &result); err != nil { // Parse []byte to the go struct pointer
			fmt.Println("Can not unmarshal JSON")
		}

		//repo request
		var result1 Repo
		resp, err = http.Get(LeftUrl + user + MidUrl)
		if err != nil {
			log.Fatal(err)
		}

		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		if err := json.Unmarshal(body, &result1); err != nil { // Parse []byte to the go struct pointer
			fmt.Println("Can not unmarshal JSON")
		}

		fmt.Println(result)
	}
}





package arg

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

// VERSION - hard code version
const VERSION = "0.1"

// Params - returns parameters needed to start program
type Params struct {
	StravaData Config
	Year       int
}

// Config - to marshall from our config.json
type Config struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RefreshToken string `json:"refresh_token"`
}

func getInfoFromUser(outputMsg string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(outputMsg)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func getRefreshCodeFromFile(fp string) Config {
	js, err := os.Open(fp)
	if err != nil {
		log.Fatal(err)
	}
	defer js.Close()

	b, _ := ioutil.ReadAll(js)

	r := Config{}
	err = json.Unmarshal([]byte(b), &r)
	if err != nil {
		log.Fatal(err)
	}
	return r
}

// ParseArgs - parses args, returns needed data
func ParseArgs() *Params {
	Param := new(Params)

	r := flag.String("r", "./config.json", "File that holds data to access Strava's API. See github for examples")
	y := flag.Int("y", time.Now().Year(), "Year to pull running distance from")
	v := flag.Bool("v", false, "Prints the version number and exit")

	flag.Parse()
	if *v {
		fmt.Println(VERSION)
		os.Exit(0)
	}

	Param.Year = *y
	fp := ""
	if _, err := os.Stat(*r); err == nil {
		fp = *r
	} else {
		log.Println("Unable to find file: " + *r)
		fp = getInfoFromUser("Provide the file that holds data to access Strava's API:")
	}

	// Tell user what they entered still can't be found
	if _, err := os.Stat(fp); os.IsNotExist(err) {
		log.Println("Unable to find file: " + fp)
		os.Exit(1)
	}

	Param.StravaData = getRefreshCodeFromFile(fp)

	return Param
}

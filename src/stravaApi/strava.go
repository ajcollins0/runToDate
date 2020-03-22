package stravaApi

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"../arg"
)

// ezpz error checking
func chkEr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// login - gets the latest accsess code to use to pull activity data
func login(config arg.Config) *loginData {
	base := "https://www.strava.com/oauth/token"

	req, err := http.NewRequest("POST", base, nil)
	chkEr(err)

	p := req.URL.Query()
	p.Add("client_id", config.ClientID)
	p.Add("client_secret", config.ClientSecret)
	p.Add("refresh_token", config.RefreshToken)

	// hard coding this so we can ensure the return accsess token is valid
	p.Add("grant_type", "refresh_token")

	req.URL.RawQuery = p.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	chkEr(err)
	defer resp.Body.Close()

	l := new(loginData)
	err = json.NewDecoder(resp.Body).Decode(&l)
	chkEr(err)

	return l
}

// hits stravas api and returns activities for one page
func getActivityPage(l *loginData, y int, page int) []activity {
	base := "https://www.strava.com/api/v3/activities"

	req, err := http.NewRequest("GET", base, nil)
	chkEr(err)

	req.Header.Set("Authorization", l.TokenType+" "+l.AccessToken)

	p := req.URL.Query()
	a := time.Date(y, time.January, 1, 0, 0, 0, 0, time.UTC).Unix()
	b := time.Date(y, time.December, 31, 23, 59, 59, 59, time.UTC).Unix()
	p.Add("after", strconv.FormatInt(a, 10))
	p.Add("before", strconv.FormatInt(b, 10))
	p.Add("page", strconv.Itoa(page))

	req.URL.RawQuery = p.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	chkEr(err)
	defer resp.Body.Close()

	var acts []activity
	err = json.NewDecoder(resp.Body).Decode(&acts)
	chkEr(err)

	return acts
}

// return the activities by paging through results
// by default stravas api will give 30 activties per page
func getActivities(l *loginData, y int) []activity {
	var acts []activity
	var hasMore = true
	page := 1
	for hasMore {
		a := getActivityPage(l, y, page)
		page++
		// stop the paging when we get no results
		if len(a) == 0 {
			hasMore = false
		} else {
			acts = append(acts, a...)
		}
	}
	return acts
}

// DistanceRun - returns the total distance run
func DistanceRun(p arg.Params) float64 {
	const METERTOMILES = 0.000621371
	l := login(p.StravaData)
	activities := getActivities(l, p.Year)
	t := 0.
	for _, v := range activities {
		if v.Type == "Run" && v.StartDate.Year() == p.Year {
			t += (v.Distance * METERTOMILES)
		}
	}
	return t
}

package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"passive/internal/models"
	"passive/internal/utils"
)

// func URegex(username string) {
// 	re := regexp.MustCompile("^@+[A-Za-z0-9]+([A-Za-z0-9]*|[._-]?[A-Za-z0-9]+)*$")
// 	if !re.MatchString(username) {
// 		color.Red("\nInvalid full name!")
// 		os.Exit(0)
// 	}
// }

func Adobe(username string) {
	jsonStr := []byte(`{"username":"` + username + `"}`)

	client := &http.Client{}
	r, err := http.NewRequest("POST", utils.AdobeEndpoint, bytes.NewBuffer(jsonStr)) // URL-encoded payload
	if err != nil {
		log.Fatal(err)
	}
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36")
	r.Header.Add("X-Ims-Clientid", "adobedotcom2")

	res, err := client.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode == 200 {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		var response models.AdobeResponse
		json.Unmarshal(body, &response)
		if len(response) > 0 {
			utils.Social_result = append(utils.Social_result, "Adobe \U0001f440")
		} else {
			utils.Social_result = append(utils.Social_result, "Adobe [Not here!]")
		}
	} else {
		utils.Social_result = append(utils.Social_result, "Adobe [Couldn't check!]")
	}
}

func Discord(username string) {
	jsonStr := []byte(`{"username":"` + username + `","username":"asdsadsad","password":"q1e31e12r13*","invite":null,"consent":true,"date_of_birth":"1973-05-09","gift_code_sku_id":null,"captcha_key":null,"promotional_username_opt_in":false}`)

	client := &http.Client{}
	r, err := http.NewRequest("POST", utils.DiscordEndpoint, bytes.NewBuffer(jsonStr)) // URL-encoded payload
	if err != nil {
		log.Fatal(err)
	}
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36")
	r.Header.Add("X-Debug-Options", "bugReporterEnabled")

	res, err := client.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode == 400 {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		var response models.DiscordResponse
		json.Unmarshal(body, &response)
		if len(response.Errors.Username.Errors) > 0 {
			if response.Errors.Username.Errors[0].Code == "username_ALREADY_REGISTERED" {
				utils.Social_result = append(utils.Social_result, "Discord \U0001f440")
			} else {
				utils.Social_result = append(utils.Social_result, "Discord [Not here!]")
			}
		} else {
			utils.Social_result = append(utils.Social_result, "Discord [Not here!]")
		}
	} else if res.StatusCode == 429 {
		utils.Social_result = append(utils.Social_result, "Discord [Rate limited!]")
	} else {
		utils.Social_result = append(utils.Social_result, "Discord [Couldn't check!]")
	}
}

func getCSRFToken() string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", utils.InstagramCSRFEndpoint, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36")
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode == 200 {
		re := regexp.MustCompile(`(?m){\\"config\\":{\\"csrf_token\\":\\"(.*?)\\"`)
		match := re.FindStringSubmatch(string(body))
		if len(match) > 0 {
			return match[1]
		}
	}
	return ""
}

func Instagram(username string) {
	var token string = getCSRFToken()
	if token == "" {
		utils.Social_result = append(utils.Social_result, "Instagram [Couldn't check!]")
	} else {
		data := url.Values{}
		data.Set("username", username)

		client := &http.Client{}
		r, err := http.NewRequest("POST", utils.InstagramEndpoint, strings.NewReader(data.Encode())) // URL-encoded payload
		if err != nil {
			log.Fatal(err)
		}
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		r.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36")
		r.Header.Add("Cookie", "csrftoken="+token+";")
		r.Header.Add("X-Csrftoken", token)

		res, err := client.Do(r)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		if res.StatusCode == 200 {
			body, err := io.ReadAll(res.Body)
			if err != nil {
				log.Fatal(err)
			}
			match, _ := regexp.MatchString("username_is_taken", string(body))
			if match {
				utils.Social_result = append(utils.Social_result, "Instagram \U0001f440")
			} else {
				utils.Social_result = append(utils.Social_result, "Instagram [Not here!]")
			}
		} else {
			utils.Social_result = append(utils.Social_result, "Instagram [Couldn't check!]")
		}
	}
}

func Spotify(username string) {
	data := url.Values{}
	data.Set("validate", "1")
	data.Set("username", username)
	client := &http.Client{}
	r, err := http.NewRequest("POST", utils.SpotifyEndpoint, strings.NewReader(data.Encode())) // URL-encoded payload
	if err != nil {
		log.Fatal(err)
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36")

	res, err := client.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode == 200 {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		var response models.SpotifyResponse
		json.Unmarshal(body, &response)
		if response.Status == 20 {
			utils.Social_result = append(utils.Social_result, "Spotify \U0001f440")
		} else {
			utils.Social_result = append(utils.Social_result, "Spotify [Not here!]")
		}
	} else {
		utils.Social_result = append(utils.Social_result, "Spotify [Couldn't check!]")
	}
}

func Twitter(username string) {
	data := url.Values{}
	data.Set("username", username)

	r, err := http.Get(utils.TwitterEndpoint + "?" + data.Encode())
	if err != nil {
		log.Fatal(err)
	}
	r.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36")
	if err != nil {
		log.Fatal(err)
	}
	if r.StatusCode == 200 {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		var response models.TwitterResponse
		json.Unmarshal(body, &response)
		if response.Taken {
			utils.Social_result = append(utils.Social_result, "Twitter \U0001f440")
		} else {
			utils.Social_result = append(utils.Social_result, "Twitter [Not here!]")
		}
	} else {
		utils.Social_result = append(utils.Social_result, "Twitter [Couldn't check!]")
	}
}

package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"

	"passive/internal/utils"
)

func IPRegex(arg string) {
	re := regexp.MustCompile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)
	if !re.MatchString(arg) {
		fmt.Println("\nInvalid IP!")
		os.Exit(0)
	}
}

func IPAPI(arg string) {
	ip4api := arg
	if ip4api != "" {
		client := &http.Client{}
		f, err := os.OpenFile("result2.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o755)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		req, _ := http.NewRequest("GET", utils.IPAPIURL+ip4api+"/json/", nil)
		req.Header.Set("User-Agent", "mosint")
		resp, _ := client.Do(req)
		body, _ := io.ReadAll(resp.Body)
		json.Unmarshal(body, &utils.Ipapi_result)

	}
}

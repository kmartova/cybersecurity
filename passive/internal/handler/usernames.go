package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync/atomic"
	"time"

	downloaders "github.com/tdh8316/Investigo/downloaders"

	"github.com/dlclark/regexp2"
	"github.com/fatih/color"
	"golang.org/x/net/proxy"
)

const (
	userAgent       string = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36"
	torProxyAddress string = "socks5://127.0.0.1:9050"
	maxGoroutines   uint   = 32
)

// Result of Investigo function
type Result struct {
	Username string
	Exist    bool
	Proxied  bool
	Site     string
	URL      string
	URLProbe string
	Link     string
	Err      bool
	ErrMsg   string
}

// A SiteData struct for json datatype
type SiteData struct {
	ErrorType string      `json:"errorType"`
	ErrorMsg  interface{} `json:"errorMsg"`
	URL       string      `json:"url"`
	URLMain   string      `json:"urlMain"`
	URLProbe  string      `json:"urlProbe"`
	URLError  string      `json:"errorUrl"`
	// TODO: Add headers
	UsedUsername   string `json:"username_claimed"`
	UnusedUsername string `json:"username_unclaimed"`
	RegexCheck     string `json:"regexCheck"`
	// Rank int`json:"rank"`
}

// RequestError interface
type RequestError interface {
	Error() string
}

type counter struct {
	n int32
}

func (c *counter) Add() {
	atomic.AddInt32(&c.n, 1)
}

func (c *counter) Get() int {
	return int(atomic.LoadInt32(&c.n))
}

func uNames(string) {
	// Loads site data from sherlock database and assign to a variable.
	initializeSiteData(options.updateBeforeRun)

	// Make the guard before goroutines run
	guard = make(chan int, maxGoroutines)

	if options.specifySite {
		// No case sensitive
		_siteData := map[string]SiteData{}
		for k, v := range siteData {
			_siteData[strings.ToLower(k)] = v
		}

		for _, siteName := range specifiedSites {
			if val, ok := _siteData[strings.ToLower(siteName)]; ok {
				siteData = map[string]SiteData{}
				siteData[siteName] = val
			}
		}

	}
	f, err := os.OpenFile("result3.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o755)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	for _, username := range os.Args[1:] {
		if options.noColor {
			fmt.Printf("\nInvestigating %s on:\n", username)
		} else {
			fmt.Fprintf(color.Output, "Investigating %s on:\n", color.HiGreenString(username))
		}
		waitGroup.Add(len(siteData))
		for site := range siteData {
			guard <- 1
			go func(site string) {
				defer waitGroup.Done()
				res := Investigo(username, site, siteData[site])
				WriteResult(res, f)
				<-guard
			}(site)
		}
		waitGroup.Wait()
	}
}

func initializeSiteData(forceUpdate bool) {
	jsonFile, err := os.Open(dataFileName)
	if err != nil || forceUpdate {
		if err != nil {
			if options.noColor {
				fmt.Printf(
					"[!] Cannot open database \"%s\"\n",
					dataFileName,
				)
			} else {
				fmt.Fprintf(
					color.Output,
					"[%s] Cannot open database \"%s\"\n",
					color.HiRedString("!"), dataFileName,
				)
			}
		}
		if options.noColor {
			fmt.Printf(
				"%s Update database: %s",
				"[!]",
				"Downloading...",
			)
		} else {
			fmt.Fprintf(
				color.Output,
				"[%s] Update database: %s",
				color.HiBlueString("!"),
				color.HiYellowString("Downloading..."),
			)
		}

		if forceUpdate {
			jsonFile.Close()
		}

		r, err := Request("https://raw.githubusercontent.com/sherlock-project/sherlock/master/sherlock/resources/data.json")

		if err != nil || r.StatusCode != 200 {
			if options.noColor {
				fmt.Printf(" [%s]\n", "Failed")
			} else {
				fmt.Fprintf(color.Output, " [%s]\n", color.HiRedString("Failed"))
			}
			if err != nil {
				panic("Failed to update database.\n" + err.Error())
			} else {
				panic("Failed to update database: " + r.Status)
			}
		} else {
			defer r.Body.Close()
		}
		if _, err := os.Stat(dataFileName); !os.IsNotExist(err) {
			if err = os.Remove(dataFileName); err != nil {
				panic(err)
			}
		}
		_updateFile, _ := os.OpenFile(dataFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o600)
		if _, err := _updateFile.WriteString(ReadResponseBody(r)); err != nil {
			if options.noColor {
				fmt.Printf("Failed to update data.\n")
			} else {
				fmt.Fprint(color.Output, color.RedString("Failed to update data.\n"))
			}
			panic(err)
		}

		_updateFile.Close()
		jsonFile, _ = os.Open(dataFileName)

		if options.noColor {
			fmt.Println(" [Done]")
		} else {
			fmt.Fprintf(color.Output, " [%s]\n", color.GreenString("Done"))
		}
	}

	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic("Failed to read file:" + dataFileName)
	} else {
		err := json.Unmarshal(byteValue, &siteData)
		if err != nil {
			panic(err.Error())
		}
	}
}

// Request makes an HTTP request
func Request(target string) (*http.Response, RequestError) {
	request, err := http.NewRequest("GET", target, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("User-Agent", userAgent)

	client := &http.Client{
		Timeout: time.Duration(timeoutSecs) * time.Second,
	}

	if options.withTor {
		tbProxyURL, err := url.Parse(torProxyAddress)
		if err != nil {
			return nil, err
		}
		tbDialer, err := proxy.FromURL(tbProxyURL, proxy.Direct)
		if err != nil {
			return nil, err
		}
		tbTransport := &http.Transport{
			Dial: tbDialer.Dial,
		}
		client.Transport = tbTransport
	}

	return client.Do(request)
}

// ReadResponseBody reads response body and return string
func ReadResponseBody(response *http.Response) string {
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	return string(bodyBytes)
}

// Investigo investigate if username exists on social media.
func Investigo(username string, site string, data SiteData) Result {
	var u, urlProbe string
	var result Result

	// URL to be displayed
	u = strings.Replace(data.URL, "{}", username, 1)

	// URL used to check if user exists.
	// Mostly same as variable `u`
	if data.URLProbe != "" {
		urlProbe = strings.Replace(data.URLProbe, "{}", username, 1)
	} else {
		urlProbe = u
	}

	if data.RegexCheck != "" {
		re := regexp2.MustCompile(data.RegexCheck, 0)
		if match, _ := re.MatchString(username); !match {
			return Result{
				Username: username,
				URL:      data.URL,
				Proxied:  options.withTor,
				Site:     site,
				Exist:    false,
				Err:      false,
			}
		}
	}

	r, err := Request(urlProbe)
	if err != nil {
		if r != nil {
			r.Body.Close()
		}
		return Result{
			Username: username,
			URL:      data.URL,
			URLProbe: data.URLProbe,
			Proxied:  options.withTor,
			Exist:    false,
			Site:     site,
			Err:      true,
			ErrMsg:   err.Error(),
		}
	}

	// check error types
	switch data.ErrorType {
	case "status_code":
		if r.StatusCode == http.StatusOK {
			result = Result{
				Username: username,
				URL:      data.URL,
				URLProbe: data.URLProbe,
				Proxied:  options.withTor,
				Exist:    true,
				Link:     u,
				Site:     site,
			}
		} else {
			result = Result{
				Username: username,
				URL:      data.URL,
				Proxied:  options.withTor,
				Site:     site,
				Exist:    false,
				Err:      false,
			}
		}
	case "message":
		switch errType := data.ErrorMsg.(type) {
		case string:
			if !strings.Contains(ReadResponseBody(r), data.ErrorMsg.(string)) {
				result = Result{
					Username: username,
					URL:      data.URL,
					URLProbe: data.URLProbe,
					Proxied:  options.withTor,
					Exist:    true,
					Link:     u,
					Site:     site,
				}
			} else {
				result = Result{
					Username: username,
					URL:      data.URL,
					Proxied:  options.withTor,
					Site:     site,
					Exist:    false,
					Err:      false,
				}
			}
		case []interface{}:
			_flag := false
			for _, msgString := range (data.ErrorMsg).([]interface{}) {
				if strings.Contains(ReadResponseBody(r), msgString.(string)) {
					_flag = true
					break
				}
			}
			result = Result{
				Username: username,
				URL:      data.URL,
				URLProbe: data.URLProbe,
				Proxied:  options.withTor,
				Exist:    _flag,
				Link:     u,
				Site:     site,
			}
		default:
			panic(fmt.Sprintf("%s: Unsupported errorMsg type %T", site, errType))
		}
	case "response_url":
		// In the original Sherlock implementation,
		// the error type `response_url` works as `status_code`.
		if (r.StatusCode <= 300 || r.StatusCode < 200) && r.Request.URL.String() == u {
			result = Result{
				Username: username,
				URL:      data.URL,
				URLProbe: data.URLProbe,
				Proxied:  options.withTor,
				Exist:    true,
				Link:     u,
				Site:     site,
			}
		} else {
			result = Result{
				Username: username,
				URL:      data.URL,
				Proxied:  options.withTor,
				Site:     site,
				Exist:    false,
				Err:      false,
			}
		}
	default:
		result = Result{
			Username: username,
			Proxied:  options.withTor,
			Exist:    false,
			Err:      true,
			ErrMsg:   "Unsupported error type `" + data.ErrorType + "`",
			Site:     site,
		}
	}

	if result.Exist && options.download {
		// Check whether the downloader for this site exists and run it
		if downloadFunc, ok := downloaders.Downloaders[strings.ToLower(site)]; ok {
			downloadFunc.(func(string, *log.Logger))(u, logger)
		}
	}

	r.Body.Close()

	return result
}

// WriteResult writes investigation result to stdout and file
func WriteResult(result Result, f *os.File) {
	// fmt.Println(result)
	if result.Exist {
		streamLogger.Printf("[%s] %s: %s\n", "+", result.Site, result.Link)

		f.Write([]byte(fmt.Sprintf("[%s] %s: %s\n", "+", result.Site, result.Link)))
	} else {
		if options.verbose {
			if result.Err {
				streamLogger.Printf("[%s] %s: %s: %s", "!", result.Site, "ERROR", result.ErrMsg)
			} else {
				streamLogger.Printf("[%s] %s: %s", "-", result.Site, "Not Found!")
			}
		}
	}

	if options.noColor {
		if result.Exist {
			logger.Printf("[%s] %s: %s\n", "+", result.Site, result.Link)
		} else {
			if options.verbose {
				if result.Err {
					logger.Printf("[%s] %s: %s: %s", "!", result.Site, "ERROR", result.ErrMsg)
				} else {
					logger.Printf("[%s] %s: %s", "-", result.Site, "Not Found!")
				}
			}
		}
	} else {
		if result.Exist {
			logger.Printf("[%s] %s: %s\n", color.HiGreenString("+"), color.HiWhiteString(result.Site), result.Link)
		} else {
			if options.verbose {
				if result.Err {
					logger.Printf("[%s] %s: %s: %s", color.HiRedString("!"), result.Site, color.HiMagentaString("ERROR"), color.HiRedString(result.ErrMsg))
				} else {
					logger.Printf("[%s] %s: %s", color.HiRedString("-"), result.Site, color.HiYellowString("Not Found!"))
				}
			}
		}
	}
}

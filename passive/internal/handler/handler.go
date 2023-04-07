package handler

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	"passive/internal/controllers"
	"passive/internal/utils"

	downloaders "github.com/tdh8316/Investigo/downloaders"

	"github.com/fatih/color"
	"github.com/gammazero/workerpool"
)

var Controllers = []func(string){
	controllers.Adobe,
	controllers.Discord,
	controllers.Instagram,
	controllers.Spotify,
	controllers.Twitter,
	controllers.IPAPI,
}

var (
	guard          chan int
	waitGroup      = &sync.WaitGroup{}
	outStream      = new(strings.Builder)
	logger         = log.New(color.Output, "", 0)
	streamLogger   = log.New(outStream, "", 0)
	dataFileName   = "data.json"
	timeoutSecs    = 60
	specifiedSites []string
	siteData       = map[string]SiteData{}
)

var options struct {
	fn              bool
	ip              bool
	u               bool
	noColor         bool
	noOutput        bool
	verbose         bool
	updateBeforeRun bool
	runTest         bool
	useCustomData   bool
	withTor         bool
	specifySite     bool
	download        bool
}

func Run() {
	usernames := parseArguments()
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Welcome to passive v1.0.0")
		fmt.Println("")
		fmt.Println("OPTIONS:")
		fmt.Println("-fn         Search with full-name")
		fmt.Println("-ip         Search with ip address")
		fmt.Println("-u          Search with username")
		fmt.Println("")
		fmt.Println("Enter OPTIONS + 'full name' || 'IP' || '@login':")
		os.Exit(0)
	}

	for _, username := range usernames {

		if options.ip {
			controllers.IPRegex(username) // ip
			wp := workerpool.New(14)
			for _, controller := range Controllers {
				controller := controller
				wp.Submit(func() {
					controller(username)
				})
			}
			wp.StopWait()
			utils.PrintIPAPI(utils.Ipapi_result)

		}
		if options.u {
			uNames(username)
			wp := workerpool.New(14)
			for _, controller := range Controllers {
				controller := controller
				wp.Submit(func() {
					controller(username)
				})
			}
		}
		if options.fn {
			// controllers.FnRegex(username)
			wp := workerpool.New(14)
			for _, controller := range Controllers {
				controller := controller
				wp.Submit(func() {
					controller(username)
				})
			}
			wp.StopWait()
			utils.PrintGoogle(utils.Googling_result)
		}
	}
	println()
	whiteBackground := color.New(color.FgRed).Add(color.BgWhite)
	println("Investigate:", whiteBackground.Sprint(usernames))
}

func HasElement(array []string, targets ...string) (bool, int) {
	for index, item := range array {
		for _, target := range targets {
			if item == target {
				return true, index
			}
		}
	}
	return false, -1
}

func parseArguments() []string {
	args := os.Args[1:]
	var argIndex int

	if help, _ := HasElement(args, "-h", "--help"); help {
		fmt.Print(
			`passive --help

			Welcome to passive v1.0.0
			
OPTIONS:
		-fn         Search with full-name
		-ip         Search with ip address
		-u          Search with username
`)
		os.Exit(0)
	}
	options.ip, argIndex = HasElement(args, "-ip")
	if options.ip {
		args = append(args[:argIndex], args[argIndex+1:]...)
	}
	options.u, argIndex = HasElement(args, "-u")
	if options.u {
		args = append(args[:argIndex], args[argIndex+1:]...)
	}
	options.fn, argIndex = HasElement(args, "-fn")
	if options.fn {
		args = append(args[:argIndex], args[argIndex+1:]...)
	}
	options.noOutput, argIndex = HasElement(args, "--no-output")
	if options.noColor {
		args = append(args[:argIndex], args[argIndex+1:]...)
	}

	options.withTor, argIndex = HasElement(args, "-t", "--tor")
	if options.withTor {
		args = append(args[:argIndex], args[argIndex+1:]...)
	}

	options.runTest, argIndex = HasElement(args, "--test")
	if options.runTest {
		options.noOutput = true
		args = append(args[:argIndex], args[argIndex+1:]...)
	}

	options.verbose, argIndex = HasElement(args, "-v", "--verbose")
	if options.verbose {
		args = append(args[:argIndex], args[argIndex+1:]...)
	}

	options.updateBeforeRun, argIndex = HasElement(args, "--update")
	if options.updateBeforeRun {
		args = append(args[:argIndex], args[argIndex+1:]...)
	}

	options.useCustomData, argIndex = HasElement(args, "--database")
	if options.useCustomData {
		dataFileName = args[argIndex+1]
		args = append(args[:argIndex], args[argIndex+2:]...)
	}

	options.specifySite, argIndex = HasElement(args, "--site")
	if options.specifySite {
		specifiedSites = strings.Split(strings.ToLower(args[argIndex+1]), ",")
		// Use verbose output
		options.verbose = true
		args = append(args[:argIndex], args[argIndex+2:]...)
	}

	options.download, argIndex = HasElement(args, "-d", "--download")
	if options.download {
		if len(args) <= 1 {
			fmt.Println("List of sites that can download userdata")
			for key := range downloaders.Downloaders {
				fmt.Fprintf(color.Output, "[%s] %s\n", color.HiGreenString("+"), color.HiWhiteString(key))
			}
			os.Exit(0)
		}
		args = append(args[:argIndex], args[argIndex+1:]...)
	}

	options.useCustomData, argIndex = HasElement(args, "--timeout")
	if options.useCustomData {
		_timeoutSecs, err := strconv.Atoi(args[argIndex+1])
		timeoutSecs = _timeoutSecs
		if err != nil {
			panic(err)
		}
		args = append(args[:argIndex], args[argIndex+2:]...)
	}
	return args
}

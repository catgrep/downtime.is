package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
)

var (
	version   = "dev"
	githubUrl = "github.com/bosdhill/downtime.is"
)

func main() {
	port := flag.Int("p", 8080, "Port to listen on")
	portLong := flag.Int("port", 8080, "Port to listen on")
	showVersion := flag.Bool("version", false, "Print version information")
	flag.Parse()

	if *showVersion {
		fmt.Printf("version %s\n", version)
		return
	}

	finalPort := *port
	if *portLong != 8080 {
		finalPort = *portLong
	}

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", finalPort),
		Handler: http.HandlerFunc(handleRequest),
	}

	log.Printf("Starting server on port %d", finalPort)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/")
	log.Printf("Incoming request: %s %s", r.Method, r.URL.Path)

	if path == "" {
		// Default is no downtime
		path = "0"
	}

	downtime, err := parseDowntimeDuration(path)
	if err != nil {
		log.Printf("Invalid input provided: %s", path)
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	downtimeDuration := formatDuration(downtime.Seconds())
	tmpl := fmt.Sprintf(`<!DOCTYPE html>
    <html xmlns="http://www.w3.org/1999/xhtml" lang="en">
    <head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta name="description" content="SLA downtime and uptime calculator">
    <meta name="keywords" content="sla,uptime,downtime,availability,calculator">
    <title>SLA &amp; Downtime calculator </title>
    <link rel="shortcut icon" href="/favicon.ico">
    <style>
    html, body { margin: 0; padding: 0; border: 0; height: 100%%; }
    body { background-color: #ffffff; color: #000000; font-family: arial, sans-serif; font-size: 16px; text-align: left; }
    .small { font-size: 12px; }
    .large { font-size: 20px; }
    h1 { color: #ff4500; text-align: left; }
    #body { clear: both; max-width: 800px; }
    #content { float: left; margin: 0px; padding-left: 20px; padding-right: 20px; text-align: left; max-width: 50em; }
    </style>
    </head>
    <body id="body">
    <div id="content">
    <h1 class="t" id="top">SLA uptime in case of %s downtime</h1>
    <p>Downtime with a duration of %s during a reporting period equals to the following SLA uptime percentages:</p>
    <ul>
    <li><strong>Daily reporting:</strong> %s</li>
    <li><strong>Weekly reporting:</strong> %s</li>
    <li><strong>Monthly reporting:</strong> %s</li>
    <li><strong>Quarterly reporting:</strong> %s</li>
    <li><strong>Yearly reporting:</strong> %s</li>
    </ul>
	The calculations assume a 24/7 monitoring period.

    <p>Direct link to these results: <strong><code>downtime.is/%s</code></strong></p>
    <p class="small">Source code available on <a href="https://%s">GitHub</a></p>
    </div>
    </body>
    </html>`,
		downtimeDuration, downtimeDuration,
		formatSLAPeriod(downtime.Seconds(), secondsInDay),
		formatSLAPeriod(downtime.Seconds(), secondsInWeek),
		formatSLAPeriod(downtime.Seconds(), secondsInMonth),
		formatSLAPeriod(downtime.Seconds(), secondsInQuarter),
		formatSLAPeriod(downtime.Seconds(), secondsInYear),
		path,
		githubUrl)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, tmpl)
}

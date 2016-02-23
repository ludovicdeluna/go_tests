/*
To force charge on all cores :
ab -l -r -k -n 1500 -c 500 "http://127.0.0.1:1718/test/moi/?s=10&qr=Show+QR"

To test speed in really simple case (like a Hello)
ab -l -r -k -n 1500 -c 500 "http://127.0.0.1:1718/test/moi/"
*/

package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

var (
	addr  = flag.String("addr", ":1718", "http service address") // Q=17, R=18
	templ = template.Must(template.New("qr").Parse(templateStr))
)

func main() {
	flag.Parse()
	http.Handle("/test/moi/", http.HandlerFunc(QR))
	http.Handle("/", http.HandlerFunc(redirect))

	fmt.Printf("Server started : 127.0.0.1%s (Ctrl+C to stop)\n", *addr)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/test/moi/", 301)
}

func QR(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("Request : %v\n", req.PostForm)
	// -- With Sleep --
	// var duration = time.Duration(0)
	//  if value, err := strconv.Atoi(req.FormValue("s")); err == nil {
	//  	duration = time.Duration(value) * time.Second
	//  }
	// time.Sleep(duration)

	// -- With Loop --
	var start = time.Now()
	var waitSeconds = float64(0)
	if param, err := strconv.ParseFloat(req.FormValue("s"), 32); err != nil {
		fmt.Println("Error with param : ", err)
	} else {
		waitSeconds = param
	}
	for time.Now().Sub(start).Seconds() <= waitSeconds {
	}

	var pipeline = struct {
		Start       time.Time
		WaitSeconds float64
		Param       string
	}{
		Start:       start,
		WaitSeconds: waitSeconds,
		Param: fmt.Sprintf(
			"s: %s, qr: %s",
			req.FormValue("s"),
			req.FormValue("qr"),
		),
	}

	if err := templ.Execute(w, pipeline); err != nil {
		fmt.Println("Error while rendering. See templateStr: ", err)
	}
}

const templateStr = `
<html>
<head>
<title>QR Link Generator</title>
</head>
<body>
{{if .Param}}
<!-- <img src="http://chart.apis.google.com/chart?chs=300x300&cht=qr&choe=UTF-8&chl={{.Param}}" /> -->
<br>
<ul>
<li>Parameters => {{.Param}}</li>
<li>Wait Seconds => {{.WaitSeconds}}</li>
<li>Started at => {{.Start}}</li>
</ul>
<br>
<br>
{{end}}
<form action="/test/moi/" name=f method="GET"><input maxLength=1024 size=70
name=s value="" title="Text to QR Encode"><input type=submit
value="Show QR" name=qr>
</form>
</body>
</html>
`

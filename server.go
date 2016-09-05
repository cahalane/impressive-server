package main

import (
	"flag"
	"fmt"
	"github.com/colm2/impressive"
	"hash/fnv"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func main() {
	portF := flag.String("port", ":3000", "Port number of server, default \":3000\"")

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.Handle("/cal/", http.StripPrefix("/cal/", http.FileServer(http.Dir("./genfiles"))))
	http.Handle("/getcalendar", &retrieveCal{})
	log.Fatal(http.ListenAndServe(*portF, nil))
}

type retrieveCal struct{}

func (i *retrieveCal) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	email := req.FormValue("email")
	password := req.FormValue("password")

	cal, err := impressive.GetICal(email, password)

	if err != nil {
		io.WriteString(resp, "{\"ok\":false, \"error\":\""+err.Error()+"\"}")
	} else {
		h := fnv.New64()
		h.Write([]byte(email))
		n := h.Sum64()
		filename := strconv.FormatUint(n, 32) + ".ics"

		cb := []byte(cal)
		err = ioutil.WriteFile("genfiles/"+filename, cb, 0664)
		if err != nil {
			io.WriteString(resp, "{\"ok\":false, \"error\":\""+err.Error()+"\"}")
		} else {
			host := req.Host
			calURL := fmt.Sprintf("http://%s/cal/%s", host, filename)
			io.WriteString(resp, "{\"ok\":true, \"url\":\""+calURL+"\"}")
		}
	}
}

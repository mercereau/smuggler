package main

import (
	"net/http"
	"os"
	"fmt"
	"io"
	"io/ioutil"
	"bytes"
)

func Logger(inner http.Handler, name string, filename string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f, err := os.OpenFile(filename, os.O_APPEND | os.O_CREATE | os.O_RDWR, 0666)
		if err != nil {
		  fmt.Printf("error opening file: %v", err)
		}

		defer f.Close()

		// clone Body (buffer) to pass to secondary handler
		buf, _ := ioutil.ReadAll(r.Body)
		rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
		rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf))

		r.Body = rdr2
		// parsing the body now so I can use it
    body, err := ioutil.ReadAll(io.LimitReader(rdr1, 1048576))
		if err != nil {
			panic(err)
		}
		if err := rdr1.Close(); err != nil {
			panic(err)
		}

		inner.ServeHTTP(w, r)

    f.WriteString(fmt.Sprintf("%s\n",body))
	})
}

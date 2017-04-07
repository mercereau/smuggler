package main

import (
	"log"
	"net/http"
	"flag"
	"fmt"
	"strconv"
)

func main() {
	var bindhost 	= flag.String("bindhost", "", "interface binding")
	var port 			= flag.Int("port", 8080, "server port, default 8080")
	var logfile 	= flag.String("logfile", "jsonpost.log", "logfile, default to jsonpost.log in parent directory")
  flag.Parse()

	router := NewRouter(*logfile)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", *bindhost, strconv.Itoa(*port)), router))
}

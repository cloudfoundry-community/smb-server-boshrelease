package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
    "log"
	"regexp"
	"time"
)

const (
    DEFAULT_PORT = "8080"
    NAME = "VolTest"
)

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/voltest.txt", voltest)

    var port string
    if port = os.Getenv("PORT"); len(port) == 0 {
        log.Printf("Warning, PORT not set. Defaulting to %+vn", DEFAULT_PORT)
        port = DEFAULT_PORT
    }
	fmt.Println("listening...")
    err := http.ListenAndServe(":" + port, nil)
	if err != nil {
		panic(err)
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "<html><head><title>%s</title></head><body><h3>I'm %s, instance with index: %s</h3><a href='/voltest.txt'>test me!</a></body></html>", NAME, NAME, os.Getenv("INSTANCE_INDEX"))
}

func voltest(res http.ResponseWriter, req *http.Request) {
	mountPointPath := getPath() + "/voltest.txt"

	d1 := []byte("Hello Volumen, test time: "+ time.Now().String() +"!\n")
	err := ioutil.WriteFile(mountPointPath, d1, os.ModeAppend|0644)
	if err != nil {
		writeError(res, "Writing \n", err)
		return
	}

	body, err := ioutil.ReadFile(mountPointPath)
	if err != nil {
		writeError(res, "Reading \n", err)
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(body)
	return
}

func getPath() string {
	r, err := regexp.Compile(`"container_dir":\s*"([^"]+)"`)
	if err != nil {
		panic(err)
	}

	vcapEnv := os.Getenv("VCAP_SERVICES")
	match := r.FindStringSubmatch(vcapEnv)
	if len(match) < 2 {
		fmt.Fprintf(os.Stderr, "VCAP_SERVICES is %s", vcapEnv)
		panic("failed to find container_dir in environment json")
	}

	return match[1]
}

func writeError(res http.ResponseWriter, msg string, err error) {
	res.WriteHeader(http.StatusInternalServerError)
	res.Write([]byte(msg))
	res.Write([]byte(err.Error()))
}

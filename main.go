package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

var (
	token    string = ""
	channels string = ""
	filename string = ""
)

func main() {
	t := flag.String("token", "", "token (Required)")
	c := flag.String("channels", "general", "channels (Optional)")
	f := flag.String("filename", "", "filename (Optional)")
	flag.Parse()
	if *t != "" {
		token = *t
	}
	if *c != "" {
		channels = *c
	}
	if *f != "" {
		filename = *f
	}

	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "stdin read error: %s", err)
		os.Exit(1)
	}
	content := string(b)

	param := url.Values{}
	param.Add("token", token)
	param.Add("content", content)
	if channels != "" {
		param.Add("channels", channels)
	}
	if filename != "" {
		param.Add("filename", filename)
	}

	resp, err := http.PostForm("https://slack.com/api/files.upload", param)

	if err != nil {
		fmt.Fprintf(os.Stderr, "HTTP Error: %s\n", err)
		os.Exit(1)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	data := struct {
		Ok    bool   `json:"ok"`
		Error string `json:"error"`
	}{}
	json.Unmarshal(body, &data)
	if data.Ok == false {
		fmt.Fprintf(os.Stderr, "API Error: %s\n", data.Error)
		os.Exit(1)
	}
}

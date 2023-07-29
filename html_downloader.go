package main

import (
	"io/ioutil"
	"net/http"
	"time"
)

type Downloader struct {
	Client   		 *http.Client
	UserAgent 			   string
}

func NewDownloader() *Downloader {
	return &Downloader{
		Client:   &http.Client{Timeout: 5 * time.Second},
		UserAgent: "",
	}
}

func (d *Downloader) SetUserAgent(userAgent string) {
	d.UserAgent = userAgent 
}

func (d *Downloader) Download(url string) (string, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	// Set the user agent in the request
	if d.UserAgent != "" {
		request.Header.Set("User-Agent", d.UserAgent)
	}

	response, err := d.Client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}


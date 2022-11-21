package client

import (
	"io"
	"log"
	"net/http"
	"time"
)

type Client struct {
	client http.Client
}

func (c *Client) GetMatches(date string) []byte {
	req, err := http.NewRequest("GET", "https://www.fotmob.com/api/matches?timezone=Europe%2FTirane", nil)
	if err != nil {
		log.Fatal("error %s", err)
	}

	q := req.URL.Query()
	q.Add("date", date)
	req.Header.Add("Accept", `application/json`)
	resp, err := c.client.Do(req)
	if err != nil {
		log.Fatalf("error %s", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error %s", err)
	}
	return body
}

func NewClient() Client {
	client := &Client{http.Client{Timeout: time.Duration(1) * time.Second}}
	return *client
}

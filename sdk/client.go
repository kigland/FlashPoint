package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/kigland/FlashPoint/apimod"
)

type Client struct {
	BaseURL string
	APIKey  string
	hc      *http.Client
}

func NewClient(baseURL, apiKey string) *Client {
	return &Client{
		BaseURL: baseURL,
		APIKey:  apiKey,
		hc:      http.DefaultClient,
	}
}

func (c *Client) Set(key, value string, ttl time.Duration, t string, mime string) error {
	body := apimod.SetCacheReq{
		Key:   key,
		Value: value,
		TTL:   int(ttl.Seconds()),
		Type:  t,
		Mime:  mime,
	}

	bs, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", c.BaseURL+"/set", bytes.NewBuffer(bs))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", c.APIKey)

	resp, err := c.hc.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to set cache: %s", resp.Status)
	}

	return nil
}

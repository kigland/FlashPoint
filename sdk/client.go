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

func (c *Client) set(body apimod.SetCacheReq) (apimod.SetCacheResp, error) {
	bs, err := json.Marshal(body)
	if err != nil {
		return apimod.SetCacheResp{}, err
	}

	req, err := http.NewRequest("POST", c.BaseURL+"/set", bytes.NewBuffer(bs))
	if err != nil {
		return apimod.SetCacheResp{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", c.APIKey)

	resp, err := c.hc.Do(req)
	if err != nil {
		return apimod.SetCacheResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return apimod.SetCacheResp{}, fmt.Errorf("failed to set cache: %s", resp.Status)
	}

	var respBody apimod.SetCacheResp
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return apimod.SetCacheResp{}, err
	}

	return respBody, nil
}

func (c *Client) Set(key string, value any, ttl time.Duration, t string, mime string) (apimod.SetCacheResp, error) {
	body := apimod.SetCacheReq{
		Key:   key,
		Value: value,
		TTL:   int(ttl.Seconds()),
		Type:  t,
		Mime:  mime,
	}

	return c.set(body)
}

func (c *Client) SetBinary(key string, value []byte, ttl time.Duration, mime string) (apimod.SetCacheResp, error) {
	body := apimod.SetCacheReq{
		Key:   key,
		Value: value,
		TTL:   int(ttl.Seconds()),
		Type:  "bin",
		Mime:  mime,
	}

	return c.set(body)
}

func (c *Client) SetText(key string, value string, ttl time.Duration, mime string) (apimod.SetCacheResp, error) {
	body := apimod.SetCacheReq{
		Key:   key,
		Value: value,
		TTL:   int(ttl.Seconds()),
		Type:  "txt",
		Mime:  mime,
	}

	return c.set(body)
}

func (c *Client) SetJSON(key string, value any, ttl time.Duration, mime string) (apimod.SetCacheResp, error) {
	body := apimod.SetCacheReq{
		Key:   key,
		Value: value,
		TTL:   int(ttl.Seconds()),
		Type:  "json",
		Mime:  mime,
	}

	return c.set(body)
}

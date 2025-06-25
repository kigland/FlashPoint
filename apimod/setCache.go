package apimod

type SetCacheReq struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	TTL   int    `json:"ttl"` // in seconds
	Type  string `json:"type"`
	Mime  string `json:"mime"`
}

type SetCacheResp struct {
	Key string `json:"key"`
}

package entity

type WebsiteInfo struct {
	Website string
	Ping    float32
}

type RequestStat struct {
	GotPing    int `json:"got_ping"`
	GotMaxPing int `json:"got_max_ping"`
	GotMinPing int `json:"got_min_ping"`
}

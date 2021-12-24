package models

import "time"

const (
	Subscribe = "subscribe"
	Publish   = "publish"
)

type AuthReq struct {
	RemoteIP   string `form:"remote_ip"`
	TLS        bool   `form:"tls"`
	Secret     string `form:"secret"`
	CommonName string `form:"common_name"`
}

type Authorization struct {
	Topic       string   `json:"topic"`
	Channels    []string `json:"channels"`
	Permissions []string `json:"permissions"`
}

type AuthResp struct {
	TTL            int             `json:"ttl"`
	Authorizations []Authorization `json:"authorizations"`
	Identity       string          `json:"identity"`
	IdentityURL    string          `json:"identity_url"`
	Expires        time.Time
}

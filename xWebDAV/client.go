package xWebDAV

import "net/url"

type Client struct {
	Endpoint      *url.URL
	Username      string
	Password      string
	SkipTLSVerify bool
}

func NewClient(endpoint string, username string, password string, skipTLSVerify bool) (client *Client, err error) {
	parsedURL, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	return &Client{
		Endpoint:      parsedURL,
		Username:      username,
		Password:      password,
		SkipTLSVerify: skipTLSVerify,
	}, nil
}

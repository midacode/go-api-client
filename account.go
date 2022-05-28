package client

import (
	"context"
	"net/http"
	"time"
)

type Session struct {
	SessionID string `json:"sessionID"`
	UserID    string `json:"userID"`
	IsAdmin   bool   `json:"-"`
}

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"-"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Bio       string    `json:"bio"`
	ImageURL  *string   `json:"imageURL"`
	CreatedAt time.Time `json:"createdAt"`
}

type AccountClient struct {
	base       string
	serverAuth string
	httpClient *http.Client
}

func NewAccountClient(options ClientOptions) AccountClient {
	return AccountClient{
		base:       options.BaseAccount,
		serverAuth: options.ServerAuth,
		httpClient: &http.Client{},
	}
}

func (c AccountClient) newRequest(ctx context.Context, method string, pathname string) *http.Request {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.base+pathname, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", c.serverAuth)
	return req
}

func (c AccountClient) GetSession(ctx context.Context, options RequestOptions) (*Session, *http.Response, error) {
	req := c.newRequest(ctx, http.MethodGet, "/session")
	req.Header.Set("Cookie", options.Cookie)
	req.Header.Del("Authorization")
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	var output struct {
		Data *Session `json:"data"`
	}
	err = readJSON(res, &output)
	return output.Data, res, err
}

func (c AccountClient) GetUser(ctx context.Context, userID string, options RequestOptions) (*User, *http.Response, error) {
	req := c.newRequest(ctx, http.MethodGet, "/users/"+userID)
	req.Header.Set("Cookie", options.Cookie)
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	var output struct {
		Data *User `json:"data"`
	}
	err = readJSON(res, &output)
	return output.Data, res, err
}

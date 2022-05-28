package client

type Client struct {
	Account AccountClient
}

type RequestOptions struct {
	Cookie string
}

type ClientOptions struct {
	BaseAccount string
	// common options
	UserAgent string
	//
	ServerAuth string
}

func New(options ClientOptions) Client {
	if options.BaseAccount == "" {
		options.BaseAccount = "http://my-account.account"
	}
	return Client{
		Account: NewAccountClient(options),
	}
}

package client

type Clients struct {
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

func New(options ClientOptions) Clients {
	if options.BaseAccount == "" {
		options.BaseAccount = "http://my-account.account"
	}
	return Clients{
		Account: NewAccountClient(options),
	}
}

package api

type Auth struct {
	User     string `xml:"user"`
	Password string `xml:"password"`
	Context  uint   `xml:"context"`
}

type ResponseStatus struct {
	Type string `xml:"type"`
}

const (
	Endpoint = "https://gateway.autodns.com"
)

func Authenticate(username, password string, context uint) Auth {
    return Auth{
    	User: username,
    	Password: password,
    	Context: context,
	}
}

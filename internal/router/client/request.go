package client

type Request struct {
	method  string
	url     string
	headers map[string]string
	body    []byte
}

func NewRequest(method string, url string, headers map[string]string, body []byte) *Request {
	return &Request{
		method:  method,
		url:     url,
		headers: headers,
		body:    body,
	}
}

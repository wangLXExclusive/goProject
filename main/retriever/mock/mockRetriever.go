package mock

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever struct {
	UserAgent string
	TimeOut   time.Duration
}
type Retriever1 struct {
	Content string
}

func (r Retriever1) Get(url string) string {
	return r.Content
}

func (r *Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	res, err := httputil.DumpResponse(resp, true)
	resp.Body.Close()

	if err != nil {
		panic(err)
	}
	return string(res)
}

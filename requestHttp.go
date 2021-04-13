package generate

import (
	"net/http"
	"time"

	"github.com/developmengo/generate-request/tmp"
)

// RequestHTTP func
func RequestHTTP(req *http.Request) (bool, string) {
	req.Header.Set("Accept-Language", "in-id")
	req.Header.Set("Connection", "Keep-Alive")
	req.Header.Set("Accept-Encoding", "gzip")
	req.Header.Set("User-Agent", "okhttp/4.2.2")

	client := &http.Client{Timeout: 5 * time.Second}

	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		return false, err.Error()
	}

	body := tmp.GzipDecode(resp)
	if resp.StatusCode != 200 {
		return false, body
	}

	return true, body
}

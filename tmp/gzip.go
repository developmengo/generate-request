package tmp

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
)

// GzipDecode func
func GzipDecode(resp *http.Response) string {
	var err error
	// Check that the server actually sent compressed data
	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		fmt.Println("Bentuk Resp gzip")
		reader, err = gzip.NewReader(resp.Body)
		defer reader.Close()
		if err != nil {
			return err.Error()
		}
	default:
		fmt.Println("Bentuk Resp Biasa")
		reader = resp.Body
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)
	body := buf.String()
	return body
}

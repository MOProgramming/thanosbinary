package main

import (
	"context"
	"net/url"
	"os"

	"net/http"

	"github.com/mozillazg/go-cos"
	"github.com/mozillazg/go-cos/debug"
)

func main() {
	u, _ := url.Parse("https://testhuanan-1253846586.cn-south.myqcloud.com")
	b := &cos.BaseURL{
		BucketURL: u,
	}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader:  true,
				RequestBody:    true,
				ResponseHeader: true,
				ResponseBody:   true,
			},
		},
	})

	_, err := c.Bucket.DeleteLifecycle(context.Background())
	if err != nil {
		panic(err)
	}
}

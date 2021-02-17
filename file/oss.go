package file

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com:nuxtio/nutx-go/g"
)

type Oss struct {
	*oss.Bucket
	Client *oss.Client
}

func NewOss() (*Oss, error) {
	client, err := oss.New(g.Config.Oss.Endpoint, g.Config.Oss.AppId, g.Config.Oss.AppSecret)
	if err != nil {
		return nil, err
	}
	bucket, err := client.Bucket(g.Config.Oss.BucketName)
	if err != nil {
		return nil, err
	}

	return &Oss{bucket, client}, nil
}

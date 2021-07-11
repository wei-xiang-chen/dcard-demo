package url_service

import (
	"dcard/client"
	"dcard/model"

	uuid "github.com/satori/go.uuid"
)

const (
	HostUrl = "http://localhost:9000/"
)

func Transform(urlInput *model.UrlInput) (*model.UrlOutput, error) {

	shortUrl := uuid.NewV4().String()
	shortUrl = shortUrl[0:7]

	_, err := client.RedisEngine.Set(shortUrl, *urlInput.Url, 0).Result()
	if err != nil {
		return nil, err
	}
	_, err = client.RedisEngine.ExpireAt(shortUrl, *urlInput.ExpireAt).Result()
	if err != nil {
		return nil, err
	}

	urlOutput := model.UrlOutput{Id: shortUrl, ShortUrl: HostUrl + shortUrl}
	return &urlOutput, nil
}

func GetOriginal(urlId *string) (*string, error) {

	val, err := client.RedisEngine.Get(*urlId).Result()
	if err != nil {
		return nil, err
	}

	return &val, nil
}

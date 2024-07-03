package repository

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/darulfh/skuy_pay_be/config"
	"github.com/darulfh/skuy_pay_be/model"
)

type IakApiRepository interface {
	PPDIakRepository(payload *model.PPDIakRequest) (*model.PPDIakResponse, error)
}

type iakApiRepository struct{}

func NewIakApiRepository() IakApiRepository {
	return &iakApiRepository{}
}

func (*iakApiRepository) PPDIakRepository(payload *model.PPDIakRequest) (*model.PPDIakResponse, error) {

	payload.Sign = sign(payload.RefID)

	resp, err := doRequestIak(http.MethodPost, "https://prepaid.iak.dev/api/top-up", payload)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("error reading response body")
	}

	var response model.PPDIakResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("error parsing response body: %w", err)
	}

	return &response, nil
}

func sign(id string) string {

	apiKey := config.AppConfig.ApiKeyIak
	username := config.AppConfig.UsernameIak

	hash := md5.New()
	_, _ = hash.Write([]byte(username + apiKey + id))

	md5 := hash.Sum(nil)

	return fmt.Sprintf("%x", md5)
}

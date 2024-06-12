package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/nandasafiqalfiansyah/hugTransformers/model"
)

type HuggingFaceRepository struct {
	APIKey string
}

func NewHuggingFaceRepository(apiKey string) *HuggingFaceRepository {
	return &HuggingFaceRepository{APIKey: apiKey}
}

func (r *HuggingFaceRepository) SendInferenceRequest(data model.Payload) (string, error) {

	fmt.Println(data)
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://api-inference.huggingface.co/models/bert-base-uncased", body)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+r.APIKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(respBody), nil
}

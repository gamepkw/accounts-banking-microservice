package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"

	model "github.com/gamepkw/accounts-banking-microservice/internal/models"
	"github.com/pkg/errors"
)

func (m *accountRepository) ElasticSearchAccountByAccountNo(ctx context.Context, account model.ElasticSearchAccount) (*[]string, error) {
	minDigits := len(account.AccountNo)
	maxResults := 5
	query := map[string]interface{}{
		"query": map[string]interface{}{
			// "wildcard": map[string]interface{}{
			"prefix": map[string]interface{}{
				// "account_no": fmt.Sprintf("*%s*", account.AccountNo[:minDigits]),
				"account_no": account.AccountNo[:minDigits],
			},
		},
	}

	var requestBody bytes.Buffer
	if err := json.NewEncoder(&requestBody).Encode(query); err != nil {
		return nil, errors.Wrap(err, "error encode request body")
	}

	res, err := m.elastic.Search(
		m.elastic.Search.WithContext(ctx),
		m.elastic.Search.WithIndex("account_no"),
		m.elastic.Search.WithBody(&requestBody),
		m.elastic.Search.WithSize(maxResults),
	)
	if err != nil {
		return nil, err
	}

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("Elasticsearch error: %v", e)
	}

	defer res.Body.Close()
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "error read response")
	}

	var response struct {
		Hits struct {
			Hits []struct {
				Source struct {
					AccountNo string `json:"account_no"`
				} `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	if err := json.Unmarshal(bodyBytes, &response); err != nil {
		return nil, errors.Wrap(err, "error unmarshal response")
	}

	accounts := []string{}

	for _, hit := range response.Hits.Hits {
		accounts = append(accounts, hit.Source.AccountNo)
	}

	return &accounts, nil
}

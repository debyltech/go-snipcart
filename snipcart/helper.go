package snipcart

import (
	"fmt"
	"net/http"
)

type URLQuery struct {
	Key   string
	Value string
}

func JSONGet(uri string, authName string, authValue string, queries []URLQuery) (*http.Response, error) {
	client := &http.Client{}

	request, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("%s %s", authName, authValue))

	if len(queries) > 0 {
		q := request.URL.Query()

		for _, query := range queries {
			q.Add(query.Key, query.Value)
		}

		request.URL.RawQuery = q.Encode()
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

package fetchutils

import (
	"io"
	"net/http"
	"tonible14012002/ascenda-test-cli/core/domain"
)

func FetchJSON(url string) ([]byte, *domain.Error) {
	// Fetch the data from the url
	resp, err := http.Get(url)
	if err != nil {
		return nil, domain.NewErr(err.Error(), http.StatusInternalServerError)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, domain.NewErr("Bad request", http.StatusBadRequest)
	}

	body, err := io.ReadAll(resp.Body) // Read the response body
	if err != nil {
		return nil, domain.NewErr("Error reading response body", http.StatusInternalServerError)
	}
	return body, nil
}

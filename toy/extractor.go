package toy

import (
	"fmt"
	"net/url"
)

func ExtractParamFromURL(rawURL, key string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	param := parsedURL.Query().Get(key)
	if param == "" {
		return "", fmt.Errorf("not found param of %s in %s", key, rawURL)
	}

	return param, nil
}

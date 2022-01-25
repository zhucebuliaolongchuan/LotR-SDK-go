package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/TylerBrock/colorjson"
	"github.com/spf13/pflag"
)

const (
	URICommandFlag    = "uri"
	APIKeyCommandFlag = "api-key"

	BookIDCommandFlag     = "book-id"
	GetChapterCommnadFlag = "get-chapter"

	MovieIDCommandFlag  = "movie-id"
	GetQuoteCommandFlag = "quote"

	QuoteIDCommandFlag     = "quote-id"
	ChapterIDCommandFlag   = "chapter-id"
	CharacterIDCommandFlag = "character-id"
)

func MakeRequest(flags *pflag.FlagSet, path string) ([]byte, error) {
	uri, err := flags.GetString(URICommandFlag)
	if err != nil {
		return nil, fmt.Errorf("could not read parameter %s. %v", URICommandFlag, err)
	}

	if len(uri) == 0 {
		return nil, fmt.Errorf("parameter %s could not be nil", URICommandFlag)
	}

	apiKey, _ := flags.GetString(APIKeyCommandFlag)

	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s?limit=5", uri, path), nil)
	if err != nil {
		return nil, fmt.Errorf("could not get data. %v", err)
	}

	request.Header.Add("Accept", "application/json")

	if apiKey != "" {
		request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("could not make a request. %v", err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read response body. %v", err)
	}

	return responseBytes, nil
}

func PrintPrettyJson(resp []byte) {
	var obj map[string]interface{}
	json.Unmarshal([]byte(resp), &obj)

	// Make a custom formatter with indent set
	f := colorjson.NewFormatter()
	f.Indent = 4

	// Marshall the Colorized JSON
	s, _ := f.Marshal(obj)

	fmt.Println(string(s))
}

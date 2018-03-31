package grabutils

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/mock"
	"strings"
	"testing"
)

func TestGetToken(t *testing.T) {

	file, err := clientSecretFile()
	token, err := tokenFromFile(file)
	if err != nil {
		assert.Fail(t, "No token")
	}
	assert.FileExists(t, file)
	fmt.Println(token.AccessToken)

}

func TestGetMembers(t *testing.T) {

	file, _ := clientSecretFile()
	token, _ := tokenFromFile(file)

	bearer := fmt.Sprintf("Bearer %s", token.AccessToken)
	url := fmt.Sprintf("https://graph.facebook.com/v2.11/%s"+
		"/members?fields=name&limit=4&access_token=%s",
		token.FacebookGroup[0].GroupId, token.AccessToken)

	fmt.Println(string(getData(url, bearer)))
	data := string(getData(url, bearer))

	fb := &FacebookMembers{}

	err := json.NewDecoder(strings.NewReader(data)).Decode(fb)
	assert.Nil(t, err)
	fmt.Println(fb.Page.Next)

	url = fb.Page.Next
	data = string(getData(url, bearer))

	fmt.Println(data)
	assert.Contains(t, data, "data")

}

func TestCreateDir(t *testing.T) {
	addTokenDirFile()
}

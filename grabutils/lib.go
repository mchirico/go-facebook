package grabutils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/user"
	"path/filepath"
	"fmt"
	"strings"
)

func clientSecretFile() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	tokenCacheDir := filepath.Join(usr.HomeDir, ".fgrab")
	os.MkdirAll(tokenCacheDir, 0700)
	return filepath.Join(tokenCacheDir,
		url.QueryEscape("token.json")), err
}

func tokenFromFile(file string) (Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return Token{}, err
	}
	t := &Token{}
	err = json.NewDecoder(f).Decode(t)
	defer f.Close()
	return *t, err
}

func getData(url string, bearer string) []byte {

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set(
		"Authorization",
		bearer)

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	return data

}

func addTokenDirFile() {

	usr, err := user.Current()
	if err != nil {
		return
	}
	tokenCacheDir := filepath.Join(usr.HomeDir, ".fgrab")
	os.MkdirAll(tokenCacheDir, 0700)

	f, err := os.Create(tokenCacheDir + "/token.json")
	if err != nil {
		log.Fatalf("Cannot create temp token")
	}
	defer f.Close()

	_, err = f.WriteString(SampleToken)
	f.Sync()

}

func GetMembers() string {
	file, _ := clientSecretFile()
	token, _ := tokenFromFile(file)

	bearer := fmt.Sprintf("Bearer %s", token.AccessToken)
	url := fmt.Sprintf("https://graph.facebook.com/v2.11/%s"+
		"/members?fields=name&limit=4&access_token=%s",
		token.FacebookGroup[0].GroupId, token.AccessToken)

	data := string(getData(url, bearer))

	fb := &FacebookMembers{}

	err := json.NewDecoder(strings.NewReader(data)).Decode(fb)
	if err != nil {
		log.Fatal("Token not set up")
	}

	fmt.Println(fb.Page.Next)

	url = fb.Page.Next
	data = string(getData(url, bearer))
	return data
}

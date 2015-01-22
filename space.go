package junos

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

// JunosSpace holds all of our information that we use for our server
// connection.
type JunosSpace struct {
	Host      string
	User      string
	Password  string
	Transport *http.Transport
}

// NewServer sets up our connection to the Junos Space server.
func NewServer(host, user, passwd string) *JunosSpace {
	return &JunosSpace{
		Host:     host,
		User:     user,
		Password: passwd,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
}

// APICall builds our GET request to the server, and returns the data.
func (s *JunosSpace) APICall(uri string) ([]byte, error) {
	client := &http.Client{Transport: s.Transport}
	url := fmt.Sprintf("https://%s/api/space/%s", s.Host, uri)
	req, _ := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(s.User, s.Password)
	res, err := client.Do(req)
	defer res.Body.Close()

	if err != nil {
		return nil, err
	}

	data, _ := ioutil.ReadAll(res.Body)

	return data, nil
}
package godeluge

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

func (deluge *Deluge) sendCommand(method string, params interface{}) (json.RawMessage, error) {
	reader, writer := io.Pipe()
	var err error

	go func() {
		defer writer.Close()

		var request = request{Method: method, Params: params, ID: deluge.id}

		err = json.NewEncoder(writer).Encode(&request)
	}()

	request, err1 := http.NewRequest("POST", deluge.URL, reader)
	if err1 != nil {
		return nil, err1
	}

	if method != "auth.login" {
		request.Header.Add("Cookie", deluge.session)
	}

	httpresponse, err2 := http.DefaultClient.Do(request)
	deluge.id++
	if err != nil {
		return nil, err
	}

	if err2 != nil {
		return nil, err2
	}

	defer httpresponse.Body.Close()

	var response response
	err3 := json.NewDecoder(httpresponse.Body).Decode(&response)
	if err3 != nil {
		return nil, err3
	}

	c := httpresponse.Header.Get("Set-Cookie")
	if c != "" {
		deluge.session = strings.Split(c, ";")[0]
	}

	var err4 error
	if (response.Error != delugeerror{}) {
		if response.Error.Message == "Not authenticated" {
			deluge.login()
			return deluge.sendCommand(method, params)
		}

		err4 = errors.New("github.com/bartholdbos/godeluge: " + response.Error.Message)
	}

	return response.Result, err4
}

func (deluge *Deluge) login() error {
	result, err := deluge.sendCommand("auth.login", []string{deluge.Password})

	if err != nil {
		return err
	}

	var test bool
	json.Unmarshal(result, &test)

	if !test {
		err = errors.New("Password incorrect")
	}

	return err
}

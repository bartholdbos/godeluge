package godeluge

import(
	"errors"
	"net/http"
	"io"
	"strings"
	"fmt"
	"encoding/json"
)

func (deluge *Deluge) sendCommand(method string, params interface{}) (json.RawMessage, error){
	reader, writer := io.Pipe()
	var err error

	go func () (){
		defer writer.Close()

		var request = Request{Method: method, Params: params, Id: deluge.Id}

		err = json.NewEncoder(writer).Encode(&request)
	}()

	req, err1 := http.NewRequest("POST", "http://wolkopslag.nl:8112/json", reader)
	req.Header.Add("Cookie", deluge.Session)
	resp, err11 := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}
	if err1 != nil {
		return nil, err1
	}
	if err11 != nil {
		return nil, err11
	}

	defer resp.Body.Close()

	var r Response
	err2 := json.NewDecoder(resp.Body).Decode(&r);

	c := resp.Header.Get("Set-Cookie")
	if c != ""{
		deluge.Session = strings.Split(c, ";")[0]
	}

	fmt.Println("error:" + r.Error.Message) //DEBUGGGGGGG!!!!!1

	return r.Result, err2
}

func (deluge *Deluge) login() ( error){
	result, err := deluge.sendCommand("auth.login", []string {deluge.Password})
	
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
package godeluge

import(
	"encoding/json"
	"errors"
	"strings"
)

func NewDeluge(password string) (*Deluge, error) {
	var deluge Deluge = Deluge{Password: password}
	var err error

	err = deluge.login()

	return &deluge, err
}

func (deluge Deluge) Get_Torrent_Status(hash string, types []string) (map[string]interface {}, error) {
	result, err := deluge.sendCommand("web.get_torrent_status", []interface{} {strings.ToLower(hash), types})
	if err != nil {
		return nil, err
	}

	var i interface{}
	err1 := json.Unmarshal(result, &i)
	if err1 != nil {
		return nil, err1
	}

	m := i.(map[string]interface{})

	for _, v := range types{
		if m[v] == nil {
			return nil, errors.New(v + " is invalid")
		}
	}

	return m, nil
}

func (deluge Deluge) Add_Torrents(magnet string) (error){
	result, err := deluge.sendCommand("web.add_torrents", []interface{} {[]interface{} {map[string]interface {} {"path": magnet, "options": nil}}})
	if err != nil {
		return err
	}

	var i bool
	err1 := json.Unmarshal(result, &i)

	if err1 != nil {
		return err1
	}
	if !i {
		return errors.New("Incorrect Magnet")
	}

	return nil
}
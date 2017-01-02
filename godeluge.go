package godeluge

import (
	"encoding/json"
	"errors"
	"reflect"
	"strings"
)

//NewDeluge creates a new Deluge instance
func NewDeluge(url string, password string) (deluge Deluge, err error) {
	deluge.URL = url
	deluge.Password = password

	ts := reflect.TypeOf(TorrentStatus{})
	for i := 0; i < ts.NumField(); i++ {
		statustypes = append(statustypes, ts.Field(i).Tag.Get("json"))
	}

	err = deluge.login()

	return
}

//GetTorrentStatus returns the current status of a torrent
func (deluge *Deluge) GetTorrentStatus(hash string) (status TorrentStatus, err error) {
	result, err := deluge.sendCommand("web.get_torrent_status", []interface{}{strings.ToLower(hash), statustypes})
	if err != nil {
		return
	}

	err = json.Unmarshal(result, &status)
	if err != nil {
		return
	}

	if (status == TorrentStatus{}) {
		err = errors.New("Torrent not found in Deluge")
	}

	return
}

//RemoveTorrent removes a torrent from Deluge
func (deluge *Deluge) RemoveTorrent(hash string) (err error) {
	result, err := deluge.sendCommand("core.remove_torrent", []interface{}{strings.ToLower(hash), true})
	if err != nil {
		return
	}

	var i bool
	err = json.Unmarshal(result, &i)
	if err != nil {
		return
	}

	if !i {
		err = errors.New("Torrent could not be removed")
	}

	return
}

//AddTorrent adds a torrent to Deluge
func (deluge *Deluge) AddTorrent(magnet string) (err error) {
	params := []interface{}{[]interface{}{map[string]interface{}{"path": magnet, "options": nil}}}
	result, err := deluge.sendCommand("web.add_torrents", params)
	if err != nil {
		return
	}

	var i bool
	err = json.Unmarshal(result, &i)
	if err != nil {
		return
	}

	if !i {
		err = errors.New("Torrent could not be added")
	}

	return
}

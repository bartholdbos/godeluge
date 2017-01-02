package godeluge

import (
	"encoding/json"
)

//Statustypes is an enum of all possible types for torrent status requests to Deluge
var statustypes []string

//Deluge instance
type Deluge struct {
	session  string
	Password string
	URL      string
	id       int32
}

//Request to Deluge
type request struct {
	Method string      `json:"method"`
	Params interface{} `json:"params"`
	ID     int32       `json:"id"`
}

//Response from Deluge
type response struct {
	ID     int32
	Result json.RawMessage
	Error  delugeerror
}

//Error received from Deluge
type delugeerror struct {
	Message string
	Code    int32
}

//TorrentStatus struct holds the response of GetTorrentStatus
type TorrentStatus struct {
	Name                string  `json:"name"`
	Progress            float64 `json:"progress"`
	ETA                 float64 `json:"eta"`
	State               string  `json:"state"`
	NumPeers            int     `json:"num_peers"`
	NumSeeds            int     `json:"num_seeds"`
	TotalPeers          int     `json:"total_peers"`
	TotalSeeds          int     `json:"total_seeds"`
	SeedsPeersRatio     float64 `json:"seeds_peers_ratio"`
	MaxDownloadSpeed    int     `json:"max_download_speed"`
	MaxUploadSpeed      int     `json:"max_upload_speed"`
	TimeAdded           float64 `json:"time_added"`
	TotalUploaded       int     `json:"total_uploaded"`
	TotalDone           int64   `json:"total_done"`
	TotalSize           int64   `json:"total_size"`
	DistributedCopies   float64 `json:"distributed_copies"`
	TrackerHost         string  `json:"tracker_host"`
	SavePath            string  `json:"save_path"`
	IsAutoManaged       bool    `json:"is_auto_managed"`
	Queue               int     `json:"queue"`
	Ratio               float64 `json:"ratio"`
	DownloadPayloadRate int     `json:"download_payload_rate"`
	UploadPayloadRate   int     `json:"upload_payload_rate"`
}

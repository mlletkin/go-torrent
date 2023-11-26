package bencode

import (
	"errors"
	"io"

	benc "github.com/jackpal/bencode-go"
)

var ErrFailedDecode = errors.New("faield to decode")

type metaInfo struct {
	Announce string      `bencode:"announce"`
	Info     infoTorrent `bencode:"info"`
}

type infoTorrent struct {
	Name         string `bencode:"name"`
	PiecesLength int    `bencode:"pieces length"`
	Pieces       string `bencode:"pieces"`
	Length       int    `bencode:"length"`
}

func Decode(rd io.Reader) (*metaInfo, error) {
	info := metaInfo{}
	err := benc.Unmarshal(rd, &info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}

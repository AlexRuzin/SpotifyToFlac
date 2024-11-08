package main

import (
	"os"

	"gopkg.in/ini.v1"
)

type ConfigBase struct {
	filePath string

	api       ConfigApi       `section:"api"`
	spotify   ConfigSpotify   `section:"spotify"`
	portaudio ConfigPortaudio `section:"portaudio"`
	flac      ConfigFlac      `section:"flac"`
}

type ConfigApi struct {
	clientId     string `key:"clientId"`
	clientSecret string `key:"clientSecret"`
	refreshToken string `key:"refreshToken"`
}

type ConfigSpotify struct {
	mode                string `key:"mode"`
	playlistSearch      string `key:"playlistSearch"`
	playlistAuthor      string `key:"playlistAuthor"`
	playlistTrackOffset string `key:"playlistTrackOffset"`
}

type ConfigPortaudio struct {
	defaultSpotifyDevice string `key:"defaultSpotifyDevice"`
	defaultAudioDevice   string `key:"defaultAudioDevice"`
}

type ConfigFlac struct {
	testObj string `key:"testObj"`
}

func LoadConfig(filePath string) (*ConfigBase, error) {
	if _, err := os.Stat(filePath); err == os.ErrNotExist {
		return nil, err
	}

	iniCfg, err := ini.Load(filePath)
	if err != nil {
		return nil, err
	}

	var configBase ConfigBase

	// Parse the [api] section
	err = ini.MapTo(configBase.api, iniCfg.Section("api"))
	if err != nil {
		return nil, err
	}

	// Parse the [spotify] section
	err = ini.MapTo(configBase.spotify, iniCfg.Section("spotify"))
	if err != nil {
		return nil, err
	}

	// Parse the [portaudio] section
	err = ini.MapTo(configBase.portaudio, iniCfg.Section("portaudio"))
	if err != nil {
		return nil, err
	}

	// Parse the [flac] section
	err = ini.MapTo(configBase.flac, iniCfg.Section("flac"))
	if err != nil {
		return nil, err
	}

	return &configBase, nil
}

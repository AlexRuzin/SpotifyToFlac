package main

type ConfigBase struct {
	filePath string

	api       ConfigApi       `section:"api"`
	spotify   ConfigSpotify   `section:"spotify"`
	portaudio ConfigPortaudio `section:"portaudio"`
	flac      ConfigFlac      `section:"flac"`
}

type ConfigApi struct {
}

type ConfigSpotify struct {
}

type ConfigPortaudio struct {
}

type ConfigFlac struct {
}

func LoadConfig(filePath string) (*ConfigBase, error) {
	return nil, nil
}

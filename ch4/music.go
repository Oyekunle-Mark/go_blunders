package main

import "go_blunders/ch4/singer"

func main() {
	artiste := &singer.Artiste{
		Name:  "Ginger",
		Genre: "Blues",
		ID:    1,
	}

	startMusic(artiste)
}

func startMusic(s singer.Singer) {
	song := []string{
		"hello",
		"song!",
	}

	if err := s.SingHeartOut(song); err != nil {
		panic("Song went awry!")
	}
}

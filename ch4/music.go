package main

import "go_blunders/ch4/singer"

func main() {
	
}

func startMusic(s singer.Singer) {
	song := []string {
		"hello",
		"song!",
	}

	if err := s.SingHeartOut(song); err != nil {
		panic("Song went awry!")
	}
}

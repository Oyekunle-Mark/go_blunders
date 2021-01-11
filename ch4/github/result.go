package github

import "fmt"

type Singer interface {
	SingHeartOut(lyrics []string) error
}

type Artiste struct {
	Name string
	NIN string
	ID int
}

func (a *Artiste) SingHeartOut(lyrics []string) error {
	if len(lyrics) <= 0 {
		return fmt.Errorf("%s", )
	}
	return nil
}

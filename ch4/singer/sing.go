package singer

import "fmt"

type Singer interface {
	SingHeartOut(lyrics []string) error
}

type Artiste struct {
	Name  string
	Genre string
	ID    int
}

func (a *Artiste) SingHeartOut(lyrics []string) error {
	if len(lyrics) <= 0 {
		return fmt.Errorf("error by singer %s - id %d", a.Name, a.ID)
	}

	for _, word := range lyrics {
		fmt.Printf("%s", word)
	}

	return nil
}

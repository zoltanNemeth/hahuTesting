package step_definitions

import (
	"errors"
	"github.com/DATA-DOG/godog"
	"github.com/zoltanNemeth/hahuTesting/pages"
)

func userIsOnTheSite() error {
	pages.HomePage().Page.GoTo()
	if pages.HomePage().Page.IsAt() {
		//alma := fmt.Errorf("expected %d godogs to be remaining, but there is %d", "alma", "körte")
		//fmt.Println(alma)
		return nil
	}
	return errors.New("user is not on the correct site")
}

func CommonSteps(s *godog.Suite) {
	s.Step(`^user is on the site$`, userIsOnTheSite)
}

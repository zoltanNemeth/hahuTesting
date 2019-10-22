package step_definitions

import (
	"fmt"
	"github.com/DATA-DOG/godog"
	"github.com/zoltanNemeth/hahuTesting/driver"
	"github.com/zoltanNemeth/hahuTesting/pages"
)

func userIsOnTheSite() error {
	pages.HomePage().Page.GoTo()
	if pages.HomePage().Page.IsAt() {
		return nil
	}
	actual, _ := driver.Driver().CurrentURL()
	expected := pages.HomePage().Page.URL
	return fmt.Errorf("user is not on the correct site,\n %v \ninstead of %v", actual, expected)
}

func CommonSteps(s *godog.Suite) {
	s.Step(`^user is on the site$`, userIsOnTheSite)
	// TODO: teardown function for the suite
}

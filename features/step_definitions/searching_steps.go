package step_definitions

import (
	"github.com/DATA-DOG/godog"
	"github.com/zoltanNemeth/hahuTesting/pages"
	"strconv"
)

func theVisitorLimitsThePotentialResultsTo(limit string) error {
	limitOfResults, _ := strconv.Atoi(limit)
	pages.HomePage().LimitResultsTo(limitOfResults)
	return nil
}

func theVisitorClicksOnTheSearchButton() error {
	pages.HomePage().ClickSearchButton()
	return nil
}

func resultsShouldAppearOnThePage(arg1 int) error {
	return nil
}

func SearchingFeature(s *godog.Suite) {
	s.Step(`^the visitor limits the potential results to "([^"]*)"$`, theVisitorLimitsThePotentialResultsTo)
	s.Step(`^the visitor clicks on the search button$`, theVisitorClicksOnTheSearchButton)
	s.Step(`^"([^"]*)" results should appear on the page$`, resultsShouldAppearOnThePage)
}

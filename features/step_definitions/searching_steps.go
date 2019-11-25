package step_definitions

import (
	"fmt"
	"github.com/DATA-DOG/godog"
	"github.com/zoltanNemeth/hahuTesting/pages"
	"strconv"
)

var resultsPage pages.ResultsPage

func theVisitorLimitsThePotentialResultsTo(limit string) error {
	limitOfResults, _ := strconv.Atoi(limit)
	pages.HomePage().LimitResultsTo(limitOfResults)
	return nil
}

func theVisitorClicksOnTheSearchButton() error {
	resultsPage = pages.HomePage().ClickSearchButton()
	return nil
}

func resultsShouldAppearOnThePage(numOfResults int) error {
	actual := resultsPage.GetNumOfResults()
	if numOfResults == actual {
		return nil
	}
	return fmt.Errorf("not the expected amount of results on the page, get %v instead of %v", actual, numOfResults)
}

func SearchingFeature(s *godog.Suite) {
	s.Step(`^the visitor limits the potential results to "([^"]*)"$`, theVisitorLimitsThePotentialResultsTo)
	s.Step(`^the visitor clicks on the search button$`, theVisitorClicksOnTheSearchButton)
	s.Step(`^"([^"]*)" results should appear on the page$`, resultsShouldAppearOnThePage)
}

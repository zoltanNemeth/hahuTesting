package step_definitions

import (
	"fmt"
	"github.com/DATA-DOG/godog"
	"github.com/zoltanNemeth/hahuTesting/pages"
	"time"
)

var resultsPage pages.ResultsPage

func theVisitorLimitsThePotentialResultsTo(limit int) error {
	pages.HomePage().LimitResultsTo(limit)
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
	return fmt.Errorf("not the expected amount of results on the page, got %v instead of %v", actual, numOfResults)
}

func theUserSelectASpecificBrandWithToPotentialResults(from, to int) error {
	pages.HomePage().SelectBrandWithResults(from, to)
	time.Sleep(time.Duration(time.Second * 1)) // Can be removed, only for demonstrating tests
	return nil
}

func theSumOfResultsInTheSearchButtonShouldEqualToTheNumberBesideTheNameOfTheBrand() error {
	sumOfResultsBesideBrand := pages.HomePage().GetSumOfResultsBesideBrand()
	sumOfResultsInSearchButton := pages.HomePage().GetSumOfResultsInSearchButton()
	if sumOfResultsBesideBrand == sumOfResultsInSearchButton {
		return nil
	}
	return fmt.Errorf("sum of results beside the brand (%v) not equals to the sum of results in the search button (%v)", sumOfResultsBesideBrand, sumOfResultsInSearchButton)
}

func SearchingFeature(s *godog.Suite) {
	s.Step(`^the visitor limits the potential results to "([^"]*)"$`, theVisitorLimitsThePotentialResultsTo)
	s.Step(`^the visitor clicks on the search button$`, theVisitorClicksOnTheSearchButton)
	s.Step(`^"([^"]*)" results should appear on the page$`, resultsShouldAppearOnThePage)
	s.Step(`^the user select a specific brand with "([^"]*)" to "([^"]*)" potential results$`, theUserSelectASpecificBrandWithToPotentialResults)
	s.Step(`^the sum of results in the search button should equal to the number beside the name of the brand$`, theSumOfResultsInTheSearchButtonShouldEqualToTheNumberBesideTheNameOfTheBrand)
}

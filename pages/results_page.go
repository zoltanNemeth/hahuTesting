package pages

import (
	"github.com/zoltanNemeth/hahuTesting/driver"
)

type ResultsPage struct {
	page Page
}

var (
	resultsXpath = "//div[contains(@class, 'talalatisor-tartalom')]"
)

func NewResultsPage(URL string) ResultsPage {
	return ResultsPage{
		Page{
			driver.Driver(),
			URL,
		},
	}
}

func (r *ResultsPage) GetNumOfResults() int {
	return len(r.page.FindElementsByXpath(resultsXpath))
}
